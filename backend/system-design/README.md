# Advanced System Design Topics for FAANG/MANGA, SpaceX & Tesla Interviews (2025 Edition)

A comprehensive list for collaboration, real-time systems, consumer-scale, distributed storage, e-commerce, high-performance compute, networking, AI, and security.  

---

## A. Collaboration, Real-Time Sync & Editing – 6

| Topic | Complexity | Key Aspects | Resources/Notes |
|-------|------------|------------|----------------|
| []()Design a Collaborative Online Spreadsheet System | Medium | Cell locking, formula sync | FAANG-like (Google Sheets) |
| Design a Google Docs-like Real-time Collaboration System (CRDT/OT) | Hard/Advanced | Real-time sync, multi-user editing, offline support, merge semantics | Medium article; add AI for auto-complete |
| Design a Live Comment System | Medium | Threaded replies, notifications | Meta-style (Facebook comments) |
| Design a Multiplayer Game Backend | Medium | State sync, tick model, lag compensation, low-latency, anti-cheat, matchmaking | From Medium; Tesla vehicle sync |
| Design a Real-time Presence and Status Service | Easy | User online/offline, typing indicators | Common for chat apps; integrate with WebSockets |
| Design a Realtime Document Versioning System with merge semantics | Advanced | Git-like merges, history rollback | Add CRDT for conflict-free merges |

---

## B. Consumer-Scale Systems & Real-Time Features – 13

| Topic | Complexity | Key Aspects | Resources/Notes |
|-------|------------|------------|----------------|
| Design a ChatGPT-like LLM Serving Architecture | Hard | Model inference, token limits, cost optimization | 2026 trend: AI-heavy; add agentic tools |
| Design a Global Social Graph Storage | Hard | Graph DBs, sharding, privacy | Meta-specific |
| Design a Global User Profile Service | Medium | Multi-region sync, GDPR compliance | — |
| Design a Large-scale Feature Store | Advanced | Offline/online features, vector embeddings | ML infra |
| Design a Low-Latency Multi-region Notification System | Medium | Push, fan-out, reliability | Apple/Google push |
| Design a Multi-region Active-Active Newsfeed System | Hard | Ranking, A/B tests | Facebook Newsfeed |
| Design a Real-time Inference Service (10M RPS, <10 ms latency) | Advanced | GPU scaling, quantization | Tesla AI (autopilot inference) |
| Design a Real-time Personalized Recommendation Engine | Hard | Two-tower models, feedback loops | Netflix/Amazon |
| Design Instagram | Medium | Feeds, stories, media upload | FAANG staple |
| Design Privacy-Aware User Data Storage | Advanced | Encryption, access logs | 2026 focus: privacy |
| Design TikTok-Style Video Feed | Hard | Short-form content, ML ranking | ByteDance-style |
| Design Time-series Storage for Billions of Events/sec | Advanced | InfluxDB-like, compression | Monitoring tools |
| Design YouTube and Netflix | Medium | Streaming, recommendations | Content delivery |

---

## C. Distributed Storage, Consistency & Databases – 12

| Topic | Complexity | Key Aspects | Resources/Notes |
|-------|------------|------------|----------------|
| Design a Columnar Storage System for Analytics | Advanced | Query optimization, partitioning | Parquet/BigQuery-like |
| Design a Distributed Key-Value Store | Hard | CAP trade-offs, replication | Dynamo/Bigtable-style |
| Design a Distributed Metadata Service | Advanced | High availability, consensus | — |
| Design a Global Locking/Coordination Service | Hard | Leader election, watches | Zookeeper/Chubby-like |
| Design a Large-scale Blob Storage | Hard | Durability, geo-replication | S3-like |
| Design a Multi-zone Conflict-free Replication System | Advanced | Eventual consistency | CRDT-based |
| Design a Sharded SQL Database with Online Resharding | Hard | Zero-downtime migrations | Vitess-style |
| Design a Vector Database for Semantic Search/ANN | Advanced | HNSW indexes, hybrid search | Pinecone/Weaviate-scale |
| Design Log-as-Database Architecture | Hard | Stream processing, immutability | Kafka/Pravega/Lakehouse |
| Design a Wide Column Database | Hard | Sparse data | Cassandra-like |
| Design a Distributed File System | Hard | Fault tolerance, replication | HDFS |
| Design a Modern Data Lakehouse | Advanced | ACID transactions, time travel, schema evolution | Iceberg/Delta Lake |

---

## D. E-commerce, Fintech & Mission-Critical Systems – 8

| Topic | Complexity | Key Aspects | Resources/Notes |
|-------|------------|------------|----------------|
| Design a Fraud Detection System | Hard | Event ingestion + real-time ML scoring, anomaly detection | Stripe/ML |
| Design a Global Shopping Cart Service | Medium | Session replication, conflict resolution | Amazon Cart |
| Design a High-throughput Search Auction System | Hard | Real-time bidding, latency <50ms | Google Ads |
| Design an Idempotent Payment Processing Pipeline | Hard | Transactions, retries | Stripe-like |
| Design an Inventory Consistency System | Hard | Distributed locking/sagas | E-commerce |
| Design an E-commerce Service | Medium | Orders, inventory, recommendations | Shopify |
| Design an Online Payment Service | Medium | PCI compliance, gateways | PayPal |
| Design an Auction System | Hard | Bidding, winner determination | eBay |

---

## E. High-Performance Distributed Compute & ML Infra – 9

| Topic | Complexity | Key Aspects | Resources/Notes |
|-------|------------|------------|----------------|
| Design a Batch Compute Engine | Advanced | Job DAGs, fault recovery | MapReduce/Spark-like |
| Design a GPU/TPU Orchestration Cluster | Advanced | Scheduling, multi-tenancy | Kubernetes+GPUs |
| Design a Priority-based Job Scheduler | Advanced | Fairness, SLAs, backfilling & preemption | Borg/K8s |
| Design a Stream Processing Pipeline | Advanced | Windowing, state management | Flink/Spark Streaming |
| Design a Workflow Orchestration System | Medium | DAGs, retries | Airflow/Dagster-like |
| Design an Autoscaling System for Real-time Compute | Advanced | Metrics-based, predictive | AWS ASG |
| Design an LLM Fine-Tuning & Serving Platform | Hard | Distributed training, quantization-aware training | 2026 AI |
| Design a Big Data Processing Pipeline | Advanced | ETL at scale | — |
| Design a High-Performance Computing Cluster | Advanced | MPI, low-latency networks | SpaceX simulations |

---

## F. Networking, Reliability & Multi-Region Ops – 10

| Topic | Complexity | Key Aspects | Resources/Notes |
|-------|------------|------------|----------------|
| Design a Disaster Recovery System | Advanced | Backups, failover, RPO/RTO | AWS DR |
| Design a Global Edge Compute Platform | Advanced | Serverless, anycast | Cloudflare Workers/AWS Lambda@Edge |
| Design a Multi-region Load Balancing System | Hard | Geo-routing, health checks | Akamai |
| Design a Multi-tenant SaaS System | Advanced | Namespaces, quotas | Salesforce |
| Design a Real-time CDN with edge compute | Hard | Purging, prefetching | Fastly |
| Design a Zero-downtime Deployment Framework | Medium | Canary, blue/green, feature flags | CI/CD |
| Design an API Gateway | Medium | AuthN/authZ, rate-limits, quotas & observability | Envoy/Istio |
| Design a Hybrid Cloud Infrastructure | Advanced | On-prem + cloud sync | — |
| Design a Virtualization System | Advanced | Hypervisors, VMs | VMware |
| Design a Feature Flag & A/B Testing Platform | Medium | User segmentation, statistical analysis | LaunchDarkly-like |

---

## G. Search, Recommendations & Personalization (AI-Heavy in 2025) – 9

| Topic | Complexity | Key Aspects | Resources/Notes |
|-------|------------|------------|----------------|
| Design a Large-scale Embedding Store | Advanced | FAISS, vector indexes | ML serving |
| Design a Real-time Clickstream Analytics System | Medium | Event processing, dashboards | Kafka+ELK |
| Design a Recommendation Ranking Pipeline | Hard | Retrieval + scoring + rerank + two-tower | Amazon Recs |
| Design a Vector Search Engine | Advanced | Semantic search, ANN hybrid retrieval | Pinecone-like, RAG |
| Design an Agentic AI System | Hard | Tool calling, multi-agent orchestration, memory | LangChain-style |
| Design Google Search | Hard | Crawling, indexing, ranking | Full search engine |
| Design Google Search Autocomplete | Medium | Trie + ML | Google |
| Design Typeahead Suggestion | Medium | Predictive search | — |
| Design a RAG Pipeline for LLM Applications | Hard | Indexing, retrieval, reranking, context window management | Retrieval-Augmented Generation |

---

## H. SpaceX/Tesla-Specific Flavor (High-Reliability & Real-Time) – 8

| Topic | Complexity | Key Aspects | Resources/Notes |
|-------|------------|------------|----------------|
| Design a Charging Station Network | Medium | Load balancing, availability | Tesla Superchargers |
| Design a Fleet Management System for Autonomous Vehicles | Hard | Tracking, routing, charging | Tesla-specific |
| Design a Low-Latency OTA Update System for Fleet | Advanced | Delta updates, rollback safety | Tesla vehicles |
| Design a Safety-Critical Real-Time Control System | Hard | Redundancy, voting, fault tolerance | SpaceX engines |
| Design a Telemetry Pipeline for Millions of Sensors | Advanced | Low-latency ingest, anomaly detection | SpaceX: stream processing |
| Design Edge AI Inference on Vehicles/Satellites | Advanced | On-device ML, cloud fallback | Tesla Autopilot |
| Design Satellite Constellation Management | Advanced | Inter-satellite links, dynamic routing | Starlink-like |
| Design a Model Update & Versioning System for Edge Devices | Advanced | Delta updates, rollback, version tracking | Tesla / IoT |

---

## I. Security & Privacy – 3

| Topic | Complexity | Key Aspects | Resources/Notes |
|-------|------------|------------|----------------|
| Design a Secure Authentication/Authorization System | Hard | OAuth2, OIDC, RBAC/ABAC, identity federation, token management | Google/Auth0 |
| Design a Secrets Management System | Medium/Hard | Encryption at rest and in transit, audit logging, dynamic secrets | Vault-like |
| Design a Data Anonymization / Differential Privacy Pipeline | Advanced | k-anonymity, l-diversity, differential privacy budgets, GDPR | Privacy-preserving data sharing |

---

## J. Observability & Monitoring – 3

| Topic | Complexity | Key Aspects | Resources/Notes |
|-------|------------|------------|----------------|
| Design a Distributed Tracing System | Medium | Trace context propagation, sampling, storage | Jaeger/Zipkin-like, OpenTelemetry |
| Design a Metrics & Alerting System | Medium | Time-series DB, aggregation, alert routing | Prometheus + Alertmanager at scale |
| Design a Log Aggregation Pipeline | Medium | Log shippers, indexing, storage tiers, query engine | ELK/Loki at petabyte scale |
