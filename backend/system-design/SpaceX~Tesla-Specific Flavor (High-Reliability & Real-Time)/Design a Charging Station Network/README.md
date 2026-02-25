# EV Charging Platform System Design (2026–2027)

## Classic System Design Path

1. **Requirements Gathering** (Functional & Non-Functional)
2. **Capacity Planning** (Back-of-the-envelope estimation)
3. **High-level Architecture**
4. **API & Key Entities**
5. **Component Deep-dive** (backend, IoT, mobile app)
6. **Databases, Sharding, Cache**
7. **Critical Trade-offs & Bottlenecks**
8. **Reliability, Security, Observability**
9. **Extensions** (V2G, dynamic pricing, fleet charging, etc.)

---

## Step 1. Requirements Gathering

### Functional Requirements

- Driver can find the nearest available charging stations on a map
- See charger types (Level 2 / DC Fast / Ultra-fast 350kW+), price, power, availability
- Reserve a slot (15–30 min)
- Start/stop charging (via app, RFID, Plug & Charge — ISO 15118)
- Payment (card, subscription, pay-per-kWh, pay-per-minute)
- Operator (CPO) sees all station statuses in real time
- Operator can remotely enable/disable charging, update firmware
- Notifications: charging completion, errors, car low charge
- Roaming support (Hubject / Gireve / Electromaps, etc.)

### Non-Functional Requirements

- Scale: 500,000 – 5 million charging points
- Peak utilization: 10–30% of stations charging simultaneously
- Latency: status updated in <5–10 sec
- Availability: 99.95–99.99%
- Real-time: critical
- Geo-distribution: multi-region/country, regulatory variations
- IoT devices: tens/hundreds of thousands, heartbeat every 10–60sec
- Throughput: millions of messages per minute peak

---

## Step 2. Capacity Planning (Real 2025–2026 numbers)

**USA (2025–2026):**
- Public fast-charging sessions: ~141M in 2025 (+30% YoY)
- New DCFC ports: +18k–19.5k/year → 80–90k ports by end of 2026
- Avg. utilization per port: 10–20% (up to 50–70% on highways)
- Sessions per day/port: 6–12 (city), 15–25 (highways)
- Avg. session: 25–45 kWh, 30–50 min

**Large Global Network (1–5M ports):**
- Simultaneous active sessions: 8–18% (80k–900k)
- Msgs/min: heartbeat (30–60 sec) + meter values (10–60 sec) + events → 5–15M at peak
- Kafka/Pulsar: 10–50 partitions by region, geo-replication

**2026–2027 Outlook:**
- Ultra-fast (350+ kW): shorter sessions (15–25 min), higher throughput, more telemetry per port
- V2G: bidirectional MeterValues → 1.5–2× telemetry load

---

## Step 3. High-level Architecture (2025–2026)

```
App/Web/Car App (Plug&Charge)
      ↕ HTTPS/gRPC/WebSocket
  API Gateway + Auth (OAuth2/OpenID)
           ↓
    Backend services (microservices)
 ┌─────────────┬─────────────┬─────────────┐
 │ Charging    │ Location    │ Payment/    │
 │ Session     │ & Search    │ Billing     │
 │ Service     │ Service     │ Service     │
 │             │             │             │
 └─────────┬───┴─────┬───────┴────────┬───┘
           │         │               │
     (Kafka/Pulsar/NATS)
           ↓
   IoT Backend/Telemetry Layer
           ↕ MQTT / OCPP 2.0.1/2.1 (WS)
          Charging Stations
```

**Key Points:**
- OCPP 2.0.1/2.1 between station and backend (JSON over WebSocket)
- MQTT for telemetry (very cheap pub/sub for 100k+ devices)
- Event-driven, async architecture

---

## Step 4. Key Services (Microservices)

**Charging Station Management (CSMS/IoT Gateway Layer):**
- Accepts connections (OCPP 2.0.1/2.1 over WebSocket)
- Supports: heartbeat, BootNotification, StatusNotification, MeterValues, Start/StopTransaction, Reservation, etc.
- OCPP → MQTT bridge for scaling (see AWS OCPP Gateway)

**Charging Session Service:**
- Manages session lifecycle (authorization, start/stop, billing triggers)
- Active sessions stored in high-speed DB (Redis/DynamoDB)
- Event-driven integration with Payment/Billing

**Location & Availability Service:**
- Geospatial search, nearest station lookup, status (free/occupied/faulty/reserved)
- Redis geo-hash, PostgreSQL+PostGIS

**User/Account/Reservation Service:**
- Profiles, subscriptions, reservations, Plug&Charge, OAuth2/OpenID
- Optimistic locking for slots, TTL & penalties for no-show

**Payment & Billing Service:**
- Stripe/Adyen/local payment integration, per-kWh/per-minute/idle fees/subs
- Event-sourced billing: each MeterValue → delta → final invoice

**Roaming & OCPI Service:**
- OCPI 2.2+, Hubject/EMSP integration

**Monitoring/Observability/Alerting:**
- Prometheus/Grafana/Loki, alerts (offline >5min, high error rate), anomaly detection ML

---

## Step 5. Databases & Sharding

| Component               | Primary DB                     | Additionally       | Why?                                               |
|-------------------------|-------------------------------|--------------------|----------------------------------------------------|
| Charging stations       | PostgreSQL / CockroachDB      | Redis (cache)      | Consistency on config/location/firmware             |
| Active sessions         | Redis / DynamoDB              | Kafka (history)    | Ultra-low latency, TTL                              |
| Session history/Meter   | Cassandra/TimescaleDB         | S3 (raw archive)   | Heavy write, time-series                            |
| User accounts           | PostgreSQL                    | Redis              | Transactions, tokens                                |
| Locations/Search        | PostgreSQL+PostGIS            | OpenSearch         | Geo + text search                                   |
| Events/Audit            | Kafka → ClickHouse/BigQuery   | —                  | Analytics, reconciliation, ML                       |

**Sharding:**  
- By charge point ID (station registry, sessions)
- By region/country (for location/search, etc.)
- By time (for meter value time-series)
- Multi-region: active-active or active-passive (replication lag <2sec)

**Cache:** Redis Cluster for status, CDN for maps/photos.

---

## Step 6. Trade-offs & Bottlenecks

- **Consistency vs. Availability:**  
  Station status → AP (eventual consistency ok), Billing/Auth → CP (strict consistency)
- **MQTT vs. OCPP-WS:**  
  MQTT is cheaper/faster at scale, OCPP-WS is simpler for security (TLS+certs)  
  → Hybrid (OCPP→MQTT proxy) — the 2025–2026 best practice
- **Real-time updates to apps:** WebSocket/SSE or push/poll fallback
- **MeterValues flood:** Use sampling, edge aggregation, downsampling in DB

---

## Deep Dive: Modern Trends & Challenges

### 1. OCPP 2.1: Key Features & Architecture Support

**Core improvements in OCPP 2.1 (vs 2.0.1):**

- Full ISO 15118-20 support (bidirectional power/V2G, Plug & Charge)
- New Bidirectional Charging functional block — EV can return energy (V2G, V2H, V2B)
- Enhanced Smart Charging — flexible profiles, local tariff calculation (even offline)
- Advanced authorization: prepaid, ad-hoc, dynamic QR codes
- Improved Device Model: granular monitoring, custom threshold alerts (e.g., AcDcConverter temperature)
- Enhanced security: stricter profiles, SecurityEventLog

**How to support in architecture:**

- OCPP→MQTT proxy/gateway (like AWS IoT Core) — main scale pattern now
- OCPP 2.1 over WebSocket (WSS + TLS 1.3 + mutual auth) for critical commands (RemoteStart, Transaction, Reservation, SignedMeterValues)
- MQTT for high-volume telemetry (MeterValues, Heartbeat, StatusNotification, component variables), with retained/last-will for offline detection
- Plug & Charge (ISO 15118-20): EV → charger sends contract certificate (PKI); charger → CSMS: GetCertificateStatusRequest; CSMS checks via V2G Root CA (or eMSP backend). Backend must support full PKI chain; cache certificate status in Redis (with TTL for revocation lists)
- V2G/Bidirectional: Add DischargeRequest, DischargeParameter, V2G MeterValues (negative power). Separate state machine for discharge mode in Charging Session Service. Integrate with Energy Management/Grid APIs (OSCP, OpenADR) for scheduling; settlement/credits for grid discharge
- Edge computing: On-station smart charging fallback/local cost calculator if offline

### 2. Capacity Planning (2025–2026, real-world):

- US: 141M fast-charging sessions (2025)
- 80–90k DCFC ports, adding ~20k/year
- 6–12 sessions/day/port (city), 15–25 (highways)
- 25–45 kWh avg. session, 30–50min
- At 1–5M ports: 80k–900k simultaneous active, 5–15M msgs/min peak (Kafka/Pulsar/geo-replication)
- Ultra-fast/350+ kW: shorter sessions, more msgs/sec, V2G = up to 2x telemetry

### 3. Security (2026 Best Practices)

**Current threats:**
- MeterValues spoofing → billing fraud
- Replay attacks on ISO 15118 (relay emulation)
- Rogue chargers / hardware cloning
- Firmware downgrade attacks / unsigned updates
- Offline manipulation (local cost calculator abuse)

**Protection/Best Practices:**
- Transport: WSS + TLS 1.3 mandatory, mutual TLS (client certs)
- PKI for Plug & Charge (ISO 15118-20: EV cert → SECC → MO → V2G Root CA), CRL/OCSP stapling, revocation caching
- OCPP Security Profile 3: mutual authentication, encrypted payloads
- Signed firmware + secure boot on charger
- SecurityEventLog → CSMS+SIEM

**Against replay/MITM:**
- Nonce + timestamps in ISO 15118, signed messages
- Rate limiting by charger ID
- ML anomaly detection (unusual power curves, dupes, bad session stats)
- Zero-trust: every charger is a separate identity in the IoT registry

### 4. Edge Computing (local smart charging, fallback)

- Local smart charging controller (OCPP 2.1) — station enforces profile on its own, continues charging during connectivity loss
- Offline authorization (pre-approved tokens, dynamic QR, Plug & Charge)
- Local cost calculator — charging continues even during outage
- Fallback: last-known tariff via MQTT retained topic; local transaction storage — resync on reconnect
- Pros: 99.99% uptime even with network failures; Cons: more complex billing reconciliation, eventual consistency + audit logs

### 5. ML for Predictive Maintenance / Dynamic Pricing

**Predictive Maintenance:**
- Features: temp trends, error codes, power ripple, session success rate
- Model: time-series (LSTM/Prophet) on ClickHouse/Timescale
- Alert: "AcDcConverter temp > threshold in 48h" → dispatch maintenance

**Dynamic Pricing:**
- Input: grid load (OSCP), renewables forecast, station occupancy, time of use
- Output: real-time kWh price → push to app/charger display
- ML: reinforcement learning for maximizing revenue + grid stability

---

## Additional Vectors & Extensions

- Multi-region reliability & SLA
- Crypto/hardware tokens for Plug & Charge on chargers
- Extensions: V2G (Vehicle-to-Grid), dynamic pricing, fleet charging, etc.

---

**Topics for further exploration:**  
- API/endpoint design, mobile/IoT app internals, data flow deep-dive?
- Ready for interview/meetup defense?