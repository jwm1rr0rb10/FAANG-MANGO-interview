# EV Charging Platform System Design (2026–2027)

## Классический путь проектирования

1. **Уточнение требований** (Functional & Non-Functional)
2. **Оценка масштаба** (Capacity planning / Back-of-the-envelope)
3. **High-level architecture** (Общая архитектура)
4. **API & ключевые сущности**
5. **Детализация компонентов** (backend, IoT, мобильное приложение)
6. **База данных, шардинг, кэш**
7. **Критические trade-offs и bottlenecks**
8. **Надёжность, безопасность, observability**
9. **Расширение** (V2G, dynamic pricing, fleet charging и т.д.)

---

## Шаг 1. Уточнение требований

### Функциональные требования

- Водитель может найти ближайшие свободные зарядки на карте
- Просмотр типов зарядок (Level 2 / DC Fast / Ultra-fast 350 кВт+), цены, мощности, доступности
- Бронирование слота (15–30 минут)
- Старт/стоп зарядки (через app, RFID, Plug & Charge — ISO 15118)
- Оплата (карта, подписка, по кВт·ч, по времени)
- Оператор (CPO) видит статус всех станций в реальном времени
- Оператор может удалённо включать/выключать зарядку, обновлять ПО
- Уведомления: завершение зарядки, ошибка, низкий заряд у автомобиля
- Поддержка роуминга (Hubject / Gireve / Electromaps и др.)

### Нефункциональные требования

- Масштаб: 500 000 – 5 млн зарядных точек
- Пиковая нагрузка: 10–30% станций одновременно заряжают
- Latency: статус обновляется < 5–10 сек
- Доступность: 99.95–99.99%
- Реал-тайм: критично
- Геораспределённость: несколько регионов/стран, различные регуляции
- IoT-устройства: десятки-сотни тысяч; heartbeat каждые 10–60 сек
- Пропускная способность: миллионы сообщений/мин в пике

---

## Шаг 2. Capacity planning (планирование ёмкости) – реальные цифры 2025–2026

**США (2025–2026):**

- Public fast-charging сессий: ~141 млн в 2025 (+30% YoY)
- Новых DCFC портов: +18–19.5k/год, итог — 80��90k портов к концу 2026
- Средняя утилизация порта: 10–20% (на трассах до 50–70%)
- Сессий в день на порт: 6–12 (город), 15–25 (магистрали)
- Средняя сессия: 25–45 кВт·ч, 30–50 мин

**Глобальная крупная сеть (1–5 млн портов):**
- Активных сессий одновременно: 8–18% (80k–900k)
- Сообщений/мин: heartbeat (30–60 сек) + meter values (10–60 сек) + события → 5–15 млн в пике
- Kafka/Pulsar: 10–50 partition'ов по регионам, geo-replication

**Прогноз 2026–2027:**
- Ultra-fast (350+ кВт): сессии короче (15–25 мин), выше throughput, больше телеметрии на порт.
- V2G: bidirectional MeterValues → ×1.5–2 нагрузка на телеметрию.

---

## Шаг 3. High-level Architecture (2025–2026)

```
App/Web/Car App (Plug&Charge)
      ↕ HTTPS/gRPC/WebSocket
  API Gateway + Auth (OAuth2/OpenID)
           ↓
    Backend services (микросервисы)
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

**Ключевые моменты:**
- OCPP 2.0.1/2.1 между зарядкой и бэкендом (JSON over WebSocket)
- MQTT для телеметрии (cheap pub/sub для сотен тысяч устройств)
- Event-driven архитектура

---

## Шаг 4. Ключевые сервисы (микросервисы)

**Charging Station Management (CSMS/IoT Gateway Layer):**
- Принимает соединения (OCPP 2.0.1/2.1 over WebSocket)
- Поддержка heartbeat, BootNotification, StatusNotification, MeterValues, Start/StopTransaction, Reservation и др.
- OCPP → MQTT bridge для масштабирования (пример AWS OCPP Gateway)

**Charging Session Service:**
- Управление жизненным циклом сессии (authorization, start/stop, billing triggers)
- Активные сессии — хранение в high-speed store (Redis/DynamoDB)
- Событийная интеграция с Payment/Billing

**Location & Availability Service:**
- Геопоиск, поиск ближайших станций, status (free/occupied/faulty/reserved)
- Redis с geo-hash, PostgreSQL+PostGIS

**User/Account/Reservation Service:**
- Профили, ��одписки, резервации, Plug&Charge, OAuth2/OpenID
- Optimistic locking на slot, TTL & no-show penalty

**Payment & Billing Service:**
- Stripe/Adyen/local интеграции, биллинг per kWh и т.д.
- Event-sourced billing: каждое MeterValue → delta → финальный invoice

**Roaming & OCPI Service:**
- OCPI 2.2+, Hubject/EMSP интеграция

**Monitoring/Observability/Alerting:**
- Prometheus/Grafana/Loki, алерты (offline > 5 мин, ошибки), anomaly detection

---

## Шаг 5. База данных и шардинг

| Компонент               | Primary DB                     | Дополнительно      | Why?                                               |
|-------------------------|-------------------------------|--------------------|----------------------------------------------------|
| Charging stations       | PostgreSQL / CockroachDB      | Redis (cache)      | Консистентность конфигурации, location, firmware   |
| Active sessions         | Redis / DynamoDB              | Kafka (history)    | Ultra-low latency, TTL                             |
| Session history/Meter   | Cassandra/TimescaleDB         | S3 (raw archive)   | Много точек данных, time-series                    |
| User accounts           | PostgreSQL                    | Redis              | Транзакционность, токены                           |
| Locations/Search        | PostgreSQL+PostGIS            | OpenSearch         | Geo + текстовый поиск                              |
| Events/Audit            | Kafka → ClickHouse/BigQuery   | —                  | Аналитика, reconciliation, ML                      |

**Шардинг:**  
- По charge point ID (для stations и sessions)  
- По региону/стране (для поиска)  
- По времени (time-series meter values)  
- Multi-region: active-active или active-passive (replication lag < 2 сек)

**Кэш:** Redis Cluster для статуса, CDN для карт/фото.

---

## Шаг 6. Критические trade-offs и bottlenecks

- **Consistency vs. Availability:**  
  Status → AP (eventual consistency норм), Billing/Auth → CP (full consistency)
- **MQTT vs OCPP-WS:**  
  MQTT дешевле по CPU/сети, легче масштаб; OCPP-WS проще security (TLS+certs)
  → Гибрид (OCPP→MQTT proxy) — стандарт 2025–2026
- **Real-time обновления в апп:** WebSocket/SSE или push/poll fallback
- **MeterValues flood:** Use sampling, edge aggregation, downsampling в БД

---

## Глубокие детали и современные тренды

### 1. OCPP 2.1: Key Features & Архитектурная поддержка

**Главные фичи OCPP 2.1 (vs 2.0.1):**
- Полная поддержка ISO 15118-20 (bidirectional power/V2G, Plug & Charge)
- Новый функциональный блок Bidirectional Charging — EV может отдавать энергию (V2G, V2H, V2B)
- Улучшенный Smart Charging — гибкие профили, локальный расчёт стоимости
- Advanced authorization: prepaid, ad-hoc, dynamic QR
- Device Model: granular monitoring, alerts по кастомным порогам (например, температура)
- Security: строгие профили, SecurityEventLog

**Как поддерживать:**
- OCPP→MQTT proxy/gateway (пример AWS IoT Core) — масштаб и надёжность
- OCPP 2.1 over WebSocket (WSS + TLS 1.3 + mutual auth) — для критичных команд (Start, Stop, SignedMeterValues)
- MQTT — для high-volume телеметрии (MeterValues, Heartbeat), retained + last-will (offline detection)
- Plug & Charge (ISO 15118-20): EV → charger — контрактный сертификат (PKI); charger → CSMS — GetCertificateStatusRequest, PKI chain (SECC, MO, OEM, Root CA), хранение status cert-ов (cache/Redis)
- V2G/Bidirectional: DischargeRequest, V2G MeterValues (neg. power), state machine для разряда, grid API (OSCP, OpenADR) для графика, отдельный billing на V2G revenue
- Edge: локальный smart charging fallback — локальный cost calculator, offline продолжение зарядки

### 2. Capacity planning (2025–2026, реальность):

- 141 млн быстрозарядных сессий в США (2025)
- 80–90k DCFC портов — рост ~20k/год
- 6–12 сессий/день/порт (город), 15–25 (магистрали)
- 25–45 кВт·ч — средняя сессия, 30–50 мин
- Для 1–5млн портов: 80k–900k активных, 5–15 млн сообщений/мин (Kafka/Pulsar/geo-replication)
- Ultra-fast/350+ кВт: короче сессии, выше messages/sec, V2G — до 2x нагрузка телеметрии

### 3. Security (2026 best practices)

**Угрозы:**
- Spoofing MeterValues → мошенничество с биллингом
- Replay атак на ISO 15118 (relay attack)
- Rogue chargers / Cloning устройств
- Firmware downgrade / незаподписанные апдейты
- Offline manipulation of cost calculator

**Защита:**
- WSS + TLS 1.3, mutual TLS (client certs)
- PKI для Plug & Charge (ISO 15118-20: EV cert → SECC → MO → V2G Root CA), CRL/OCSP stapling, кеширование revoke-листов
- OCPP Security Profile 3: mutual auth, encryption
- Подписи firmware, secure boot на станции
- SecurityEventLog → CSMS/SIEM, monitoring

**Replay/MITM:**
- Nonce + timestamp в ISO 15118, подписи сообщений, rate limiting, ML-анализ аномалий (power curve, дубли, неожиданные charging patterns)
- Zero-trust: каждый charger — отдельная identity в IoT registry

### 4. Edge computing (fallback, smart charging локально)

- Локальный smart charging controller (ограничения по расписанию/тарифу на станции)
- Offline authorization (pre-approved tokens, dynamic QR, Plug & Charge)
- Local cost calculation
- Fallback: last-known тариф через MQTT-retained; хранение локальных транзакций → синхронизация при переподключении
- Преимущества: 99.99% uptime даже при сбоях сети, недостаток — reconciliation slo-mo

### 5. ML для predictive maintenance / динамического ценообразования

**Predictive maintenance:**
- Фичи: температура, error codes, power ripple, успешность сессий
- Модель: time-series (LSTM/Prophet), ClickHouse/Timescale
- Alert: "AcDcConverter temp > X через 48ч" → вызов инженера

**Dynamic pricing:**
- Вход: grid load (OSCP), прогноз генерации, занятость, time-of-use
- Выход: real-time price/kWh в app/charger display
- ML: reinforcement learning для максимизации выручки+стабильности сети

---

## Дополнительные шаги и векторы развития

- Мульти-региональная надёжность и SLA
- Крипто- и hardware tokens для Plug & Charge на зарядках
- Расширения: V2G (Vehicle-to-Grid), динамические тарифы, флит-зарядка

---

**Вопросы для обсуждения и проработки:**  
- На чём сфокусировать детализацию: архитектуру API, мобильное/IоT-приложение, потоки данных?
- Готовы ли к defense на собеседовании/митапе?