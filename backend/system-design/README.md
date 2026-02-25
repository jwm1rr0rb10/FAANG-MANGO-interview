# Advanced System Design Topics for FAANG/MANGA, SpaceX & Tesla Interviews (2025 Edition)

|Name Section                                                                                                                   | Size | Descriptio                                 | 
|:---------------------------------------------------------                                                                     |:-----|:-------------------------------------------|
|[**Collaboration, Real-Time Sync & Editing**](#a-collaboration-real-time-sync--editing)                                         | 6    | Design of real-time collaborative systems like Google Docs, Notion, or Figma. Core challenges: low-latency synchronization (<200 ms), concurrent editing without conflicts, operational transformation (OT) vs CRDTs for merge resolution, cursor/presence tracking, WebSocket fan-out, state reconciliation on reconnect. Key patterns: OT (Google Docs style), CRDTs (Yjs, Automerge), pub/sub (Redis Pub/Sub or Kafka for durability), in-memory canonical state (Redis), version history (snapshots to S3), conflict-free replicated data types for offline-first. Scale to thousands of concurrent editors per document with eventual convergence and minimal latency.                                            |
|[**Consumer-Scale Systems & Real-Time Features**](#b-consumer-scale-systems--real-time-features)                               | 13   | Building high-scale consumer apps with real-time updates (e.g., TikTok live, Instagram stories, Twitter/X feeds, emoji reactions in millions-viewer events). Focus: push-based delivery (WebSocket, Server-Sent Events, long-polling fallback), pub/sub at scale (Kafka/Pulsar + fan-out), low-latency streaming (<2 s), presence/heartbeat detection, emoji/reaction broadcasting, stream processing (Flink/Kafka Streams), backpressure handling, eventual consistency for non-critical data. Handles spikes (e.g., 10M+ concurrent users), hybrid push/pull, edge caching, and observability for tail latency.                                           |
|[**Distributed Storage, Consistency & Databases**](#c-distributed-storage-consistency--databases)                              | 12   | Deep dive into distributed data stores and CAP theorem trade-offs. Topics: strong vs eventual vs causal consistency, linearizability vs sequential, synchronous vs async replication, Raft/Paxos/ZAB consensus, multi-leader vs leader-follower, sharding strategies, conflict resolution (last-write-wins, CRDTs, version vectors), read-your-writes/monotonic reads, databases like CockroachDB/Spanner (strong CP), Cassandra/DynamoDB (AP eventual), TiDB (Raft + strong consistency), hybrid (NewSQL), geo-replication lag, anti-entropy, hinted handoff, quorums (W+R > N), Jepsen-style verification. Critical for mission-critical vs high-throughput workloads.                                           |
|[**E-commerce, Fintech & Mission-Critical Systems**](#d-e-commerce-fintech--mission-critical-systems)                           | 8    |  Architecture of high-stakes transactional systems (Amazon-style e-commerce, Stripe/PayPal payments, banking cores). Key: double-entry ledgers for financial integrity, idempotency + exactly-once semantics, saga pattern / distributed transactions (TCC, 2PC alternatives), PCI-DSS compliance, fraud detection (real-time rules + ML), payment orchestration (PSP integration), inventory race conditions (optimistic/pessimistic locking), order fulfillment workflows, high availability (99.999%), circuit breakers, rate limiting, eventual consistency for non-money paths, sharding by merchant/user, event sourcing for audit. Mission-critical: zero data loss, strong consistency on money movement, disaster recovery.                                          |
|[**High-Performance Distributed Compute & ML Infra**](#e-high-performance-distributed-compute--ml-infra)                        | 9    | Global-scale reliability: multi-region active-active/passive, traffic management (global load balancing, Anycast, DNS failover), latency-based routing, cross-region replication (CRDB/Spanner sync, eventual async), network partitions handling, failure detection (heartbeats, gossip), chaos engineering, SLO/SLA budgeting, blue-green/canary deployments, rollback strategies, service mesh (Istio/Linkerd), zero-trust networking, edge computing (CDN + compute@edge), QUIC/HTTP3 adoption, observability in partitioned networks. Goal: 99.99%+ global uptime with minimal cross-region latency.                                           |
|[**Networking, Reliability & Multi-Region Ops**](#f-networking-reliability--multi-region-ops)                                  | 10   | Global-scale reliability: multi-region active-active/passive, traffic management (global load balancing, Anycast, DNS failover), latency-based routing, cross-region replication (CRDB/Spanner sync, eventual async), network partitions handling, failure detection (heartbeats, gossip), chaos engineering, SLO/SLA budgeting, blue-green/canary deployments, rollback strategies, service mesh (Istio/Linkerd), zero-trust networking, edge computing (CDN + compute@edge), QUIC/HTTP3 adoption, observability in partitioned networks. Goal: 99.99%+ global uptime with minimal cross-region latency.                                           |
|[**Observability & Monitoring**](#g-observability--monitoring)                                                                  | 3    | Full-stack observability: metrics (Prometheus), distributed tracing (Jaeger/OpenTelemetry), structured logs (Loki/ELK), alerting (Alertmanager/PagerDuty), SLO/SLI/error budgets, anomaly detection (ML-based), root-cause analysis, cardinality explosion mitigation, dashboards (Grafana), chaos + canary monitoring, cost attribution per service. Goal: unknown-unknowns visibility, fast MTTR in distributed systems.                                           |

|[**Search, Recommendations & Personalization (AI-Heavy in 2025)**](#h-search-recommendations--personalization-ai-heavy-in-2025) | 9    | Modern search + recsys at scale (Google/TikTok/Netflix/Amazon style). Core: inverted indexes (Elasticsearch/OpenSearch), vector/hybrid search (dense embeddings + sparse), approximate nearest neighbors (HNSW, IVF, FAISS/PGVector), ranking (learning-to-rank, two-tower models), personalization (contextual bandits, reinforcement learning), real-time feedback loops (impressions → clicks → conversions), feature pipelines (streaming + batch), cold-start mitigation, A/B infra, diversity/serendipity, debiasing, multi-objective optimization (relevance + engagement + revenue). Heavy AI: transformers, embeddings from multimodal models, real-time updates via Kafka.                                           |
|[**Security & Privacy**](#i-security--privacy)                                                                                  | 3    | Core security architecture: zero-trust, least privilege, mTLS, PKI (Plug&Charge certs), secrets management (Vault), encryption at rest/transit, tokenization, WAF + bot detection, rate limiting + idempotency, audit logging, compliance (GDPR/CCPA/PCI/SOC2), DLP, secure boot/firmware, anomaly-based threat detection, key rotation, secure enclaves (SGX/TEE), privacy-preserving ML (federated, differential privacy). Short but critical — breaches kill trust instantly.  
|[**SpaceX/Tesla-Specific Flavor (High-Reliability & Real-Time)**](#j-spacextesla-specific-flavor-high-reliability--real-time)   | 8    |  Extreme reliability + real-time for vehicles, rockets, autonomy (Tesla FSD, Starlink, Starship telemetry). Topics: deterministic systems, fault-tolerant compute (redundancy, voting), real-time OS patterns, edge autonomy (local decisions with cloud fallback), OTA updates (atomic + rollback), telemetry streaming at scale (millions of vehicles → Kafka/Pulsar), anomaly detection, predictive maintenance (time-series ML), V2X communication, safety-critical software (ISO 26262/ASIL), zero-trust device identity, simulation-in-loop testing, high-bandwidth low-latency links (Starlink), over-the-air preconditioning/charging coordination. Focus: never-lose-control, millisecond decisions, extreme uptime.                                          |                                         |

A comprehensive list for collaboration, real-time systems, consumer-scale, distributed storage, e-commerce, high-performance compute, networking, AI, and security.  

---

## A. Collaboration, Real-Time Sync & Editing

| Topic | Complexity | Key Aspects | Resources/Notes |
|-------|------------|------------|----------------|
| [**Design a Collaborative Online Spreadsheet System**]() | Medium | Cell locking, formula sync | FAANG-like (Google Sheets) |
| [**Design a Google Docs-like Real-time Collaboration System (CRDT/OT)**]() | Hard/Advanced | Real-time sync, multi-user editing, offline support, merge semantics | Medium article; add AI for auto-complete |
| [**Design a Live Comment System**]() | Medium | Threaded replies, notifications | Meta-style (Facebook comments) |
| [**Design a Multiplayer Game Backend**]() | Medium | State sync, tick model, lag compensation, low-latency, anti-cheat, matchmaking | From Medium; Tesla vehicle sync |
| [**Design a Real-time Presence and Status Service**]() | Easy | User online/offline, typing indicators | Common for chat apps; integrate with WebSockets |
| [**Design a Realtime Document Versioning System with merge semantics**]() | Advanced | Git-like merges, history rollback | Add CRDT for conflict-free merges |

---  

## B. Consumer-Scale Systems & Real-Time Features

| Topic | Complexity | Key Aspects | Resources/Notes |
|-------|------------|------------|----------------|
| [**Design a ChatGPT-like LLM Serving Architecture**]() | Hard | Model inference, token limits, cost optimization | 2026 trend: AI-heavy; add agentic tools |
| [**Design a Global Social Graph Storage**]() | Hard | Graph DBs, sharding, privacy | Meta-specific |
| [**Design a Global User Profile Service**]() | Medium | Multi-region sync, GDPR compliance | — |
| [**Design a Large-scale Feature Store**]() | Advanced | Offline/online features, vector embeddings | ML infra |
| [**Design a Low-Latency Multi-region Notification System**]() | Medium | Push, fan-out, reliability | Apple/Google push |
| [**Design a Multi-region Active-Active Newsfeed System**]() | Hard | Ranking, A/B tests | Facebook Newsfeed |
| [**Design a Real-time Inference Service**]() (10M RPS, <10 ms latency) | Advanced | GPU scaling, quantization | Tesla AI (autopilot inference) |
| [**Design a Real-time Personalized Recommendation Engine**]() | Hard | Two-tower models, feedback loops | Netflix/Amazon |
| [**Design Instagram**]() | Medium | Feeds, stories, media upload | FAANG staple |
| [**Design Privacy-Aware User Data Storage**]() | Advanced | Encryption, access logs | 2026 focus: privacy |
| [**Design TikTok-Style Video Feed**]() | Hard | Short-form content, ML ranking | ByteDance-style |
| [**Design Time-series Storage for Billions of Events/sec**]() | Advanced | InfluxDB-like, compression | Monitoring tools |
| [**Design YouTube and Netflix**]() | Medium | Streaming, recommendations | Content delivery |

---

## C. Distributed Storage, Consistency & Databases

| Topic | Complexity | Key Aspects | Resources/Notes |
|-------|------------|------------|----------------|
| [**Design a Columnar Storage System for Analytics**]() | Advanced | Query optimization, partitioning | Parquet/BigQuery-like |
| [**Design a Distributed Key-Value Store**]() | Hard | CAP trade-offs, replication | Dynamo/Bigtable-style |
| [**Design a Distributed Metadata Service**]() | Advanced | High availability, consensus | — |
| [**Design a Global Locking/Coordination Service**]() | Hard | Leader election, watches | Zookeeper/Chubby-like |
| [**Design a Large-scale Blob Storage**]() | Hard | Durability, geo-replication | S3-like |
| [**Design a Multi-zone Conflict-free Replication System**]() | Advanced | Eventual consistency | CRDT-based |
| [**Design a Sharded SQL Database with Online Resharding**]() | Hard | Zero-downtime migrations | Vitess-style |
| [**Design a Vector Database for Semantic Search/ANN**]() | Advanced | HNSW indexes, hybrid search | Pinecone/Weaviate-scale |
| [**Design Log-as-Database Architecture**]() | Hard | Stream processing, immutability | Kafka/Pravega/Lakehouse |
| [**Design a Wide Column Database**]() | Hard | Sparse data | Cassandra-like |
| [**Design a Distributed File System**]() | Hard | Fault tolerance, replication | HDFS |
| [**Design a Modern Data Lakehouse**]() | Advanced | ACID transactions, time travel, schema evolution | Iceberg/Delta Lake |

---

## D. E-commerce, Fintech & Mission-Critical Systems

| Topic | Complexity | Key Aspects | Resources/Notes |
|-------|------------|------------|----------------|
| [**Design a Fraud Detection System**]() | Hard | Event ingestion + real-time ML scoring, anomaly detection | Stripe/ML |
| [**Design a Global Shopping Cart Service**]() | Medium | Session replication, conflict resolution | Amazon Cart |
| [**Design a High-throughput Search Auction System**]() | Hard | Real-time bidding, latency <50ms | Google Ads |
| [**Design an Idempotent Payment Processing Pipeline**]() | Hard | Transactions, retries | Stripe-like |
| [**Design an Inventory Consistency System**]() | Hard | Distributed locking/sagas | E-commerce |
| [**Design an E-commerce Service**]() | Medium | Orders, inventory, recommendations | Shopify |
| [**Design an Online Payment Service**]() | Medium | PCI compliance, gateways | PayPal |
| [**Design an Auction System**]() | Hard | Bidding, winner determination | eBay |

---

## E. High-Performance Distributed Compute & ML Infra

| Topic | Complexity | Key Aspects | Resources/Notes |
|-------|------------|------------|----------------|
| [**Design a Batch Compute Engine**]() | Advanced | Job DAGs, fault recovery | MapReduce/Spark-like |
| [**Design a GPU/TPU Orchestration Cluster**]() | Advanced | Scheduling, multi-tenancy | Kubernetes+GPUs |
| [**Design a Priority-based Job Scheduler**]() | Advanced | Fairness, SLAs, backfilling & preemption | Borg/K8s |
| [**Design a Stream Processing Pipeline**]() | Advanced | Windowing, state management | Flink/Spark Streaming |
| [**Design a Workflow Orchestration System**]() | Medium | DAGs, retries | Airflow/Dagster-like |
| [**Design an Autoscaling System for Real-time Compute**]() | Advanced | Metrics-based, predictive | AWS ASG |
| [**Design an LLM Fine-Tuning & Serving Platform**]() | Hard | Distributed training, quantization-aware training | 2026 AI |
| [**Design a Big Data Processing Pipeline**]() | Advanced | ETL at scale | — |
| [**Design a High-Performance Computing Cluster**]() | Advanced | MPI, low-latency networks | SpaceX simulations |

---

## F. Networking, Reliability & Multi-Region Ops

| Topic | Complexity | Key Aspects | Resources/Notes |
|-------|------------|------------|----------------|
| [**Design a Disaster Recovery System**]() | Advanced | Backups, failover, RPO/RTO | AWS DR |
| [**Design a Global Edge Compute Platform**]() | Advanced | Serverless, anycast | Cloudflare Workers/AWS Lambda@Edge |
| [**Design a Multi-region Load Balancing System**]() | Hard | Geo-routing, health checks | Akamai |
| [**Design a Multi-tenant SaaS System**]() | Advanced | Namespaces, quotas | Salesforce |
| [**Design a Real-time CDN with edge compute**]() | Hard | Purging, prefetching | Fastly |
| [**Design a Zero-downtime Deployment Framework**]() | Medium | Canary, blue/green, feature flags | CI/CD |
| [**Design an API Gateway**]() | Medium | AuthN/authZ, rate-limits, quotas & observability | Envoy/Istio |
| [**Design a Hybrid Cloud Infrastructure**]() | Advanced | On-prem + cloud sync | — |
| [**Design a Virtualization System**]() | Advanced | Hypervisors, VMs | VMware |
| [**Design a Feature Flag & A/B Testing Platform**]() | Medium | User segmentation, statistical analysis | LaunchDarkly-like |

---


## G. Observability & Monitoring

| Topic | Complexity | Key Aspects | Resources/Notes |
|-------|------------|------------|----------------|
| [**Design a Distributed Tracing System**]() | Medium | Trace context propagation, sampling, storage | Jaeger/Zipkin-like, OpenTelemetry |
| [**Design a Metrics & Alerting System**]() | Medium | Time-series DB, aggregation, alert routing | Prometheus + Alertmanager at scale |
| [**Design a Log Aggregation Pipeline**]() | Medium | Log shippers, indexing, storage tiers, query engine | ELK/Loki at petabyte scale |


## H. Search, Recommendations & Personalization (AI-Heavy in 2025)

| Topic | Complexity | Key Aspects | Resources/Notes |
|:-------|:------------|:------------|:----------------|
| [**Design a Large-scale Embedding Store**]() | Advanced | FAISS, vector indexes | ML serving |
| [**Design a Real-time Clickstream Analytics System**]() | Medium | Event processing, dashboards | Kafka+ELK |
| [**Design a Recommendation Ranking Pipeline**]() | Hard | Retrieval + scoring + rerank + two-tower | Amazon Recs |
| [**Design a Vector Search Engine**]() | Advanced | Semantic search, ANN hybrid retrieval | Pinecone-like, RAG |
| [**Design an Agentic AI System**]() | Hard | Tool calling, multi-agent orchestration, memory | LangChain-style |
| [**Design Google Search**]() | Hard | Crawling, indexing, ranking | Full search engine |
| [**Design Google Search Autocomplete**]() | Medium | Trie + ML | Google |
| [**Design Typeahead Suggestion**]() | Medium | Predictive search | — |
| [**Design a RAG Pipeline for LLM Applications**]() | Hard | Indexing, retrieval, reranking, context window management | Retrieval-Augmented Generation |

---

## I. Security & Privacy

| Topic | Complexity | Key Aspects | Resources/Notes |
|-------|------------|------------|----------------|
| [**Design a Secure Authentication/Authorization System**]() | Hard | OAuth2, OIDC, RBAC/ABAC, identity federation, token management | Google/Auth0 |
| [**Design a Secrets Management System**]() | Medium/Hard | Encryption at rest and in transit, audit logging, dynamic secrets | Vault-like |
| [**Design a Data Anonymization / Differential Privacy Pipeline**]() | Advanced | k-anonymity, l-diversity, differential privacy budgets, GDPR | Privacy-preserving data sharing |

---

## J. SpaceX/Tesla-Specific Flavor (High-Reliability & Real-Time)

| Topic | Complexity | Key Aspects | Resources/Notes |
|-------|------------|------------|----------------|
| [**Design a Charging Station Network**]() | Medium | Load balancing, availability | Tesla Superchargers |
| [**Design a Fleet Management System for Autonomous Vehicles**]() | Hard | Tracking, routing, charging | Tesla-specific |
| [**Design a Low-Latency OTA Update System for Fleet**]() | Advanced | Delta updates, rollback safety | Tesla vehicles |
| [**Design a Model Update & Versioning System for Edge Devices**]() | Advanced | Delta updates, rollback, version tracking | Tesla / IoT |
| [**Design a Safety-Critical Real-Time Control System**]() | Hard | Redundancy, voting, fault tolerance | SpaceX engines |
| [**Design a Telemetry Pipeline for Millions of Sensors**]() | Advanced | Low-latency ingest, anomaly detection | SpaceX: stream processing |
| [**Design Edge AI Inference on Vehicles/Satellites**]() | Advanced | On-device ML, cloud fallback | Tesla Autopilot |
| [**Design Satellite Constellation Management**]() | Advanced | Inter-satellite links, dynamic routing | Starlink-like |

---