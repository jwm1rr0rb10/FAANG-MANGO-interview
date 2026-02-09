# Updated Advanced System Design Topics for FAANG/MANGA, SpaceX & Tesla Interviews (2025 Edition)

# [A Collaboration, Real-Time Sync & Editing]()

| Name | Level | Description | Resource |
|:-----|:------|:------------|:---------|
| 1. [**Design a Collaborative Online Spreadsheet System**](link) | Medium | Cell locking, formula sync. | FAANG-like (Google Sheets). |
| 2. [**Design a Google Docs-like Real-time Collaboration System (CRDT/OT with conflict resolution)**](link) | Hard/Advanced | Real-time sync, multi-user editing, offline support, merge semantics. | Medium article in list; add AI for auto-complete. |
| 3. [**Design a Live Comment System**](link) | Medium | Threaded replies, notifications. | Meta-style (Facebook comments). |
| 4. [**Design a Multiplayer Game Backend (state sync, tick model, lag compensation)**](link) | Medium | Low-latency state, anti-cheat, matchmaking. | From Medium; For Tesla — the same vehicle sync. |
| 5. [**Design a Real-time Presence and Status Service**](link) | Easy | User online/offline, typing indicators. | Common for chat apps; integrate with WebSockets. |
| 6. [**Design a Realtime Document Versioning System with merge semantics**](link) | Advanced | Git-like merges, history rollback. | Add CRDT for conflict-free. |

---

## [B Consumer-Scale Systems & Real-Time Features]()

- 1.  [**Design a ChatGPT-like LLM Serving Architecture (incl. autoscaling, quantization, fallback)**]() 
- 2.  [**Design a Global Social Graph Storage (Facebook-scale)**]() 
- 3.  [**Design a Global User Profile Service with caching & write coalescing**]() 
- 4.  [**Design a Large-scale Feature Store (for ML training/serving)**]() 
- 5.  [**Design a Low-Latency Multi-region Notification System**]() 
- 6.  [**Design a Multi-region Active-Active Newsfeed System**]() 
- 7.  [**Design a Real-time Inference Service (10M RPS, <10 ms latency)**]() 
- 8.  [**Design a Real-time Personalized Recommendation Engine with Online Learning**]() 
- 9.  [**Design Privacy-Aware User Data Storage (GDPR, erasure, auditing, differential privacy)**]() 
- 10. [**Design TikTok-Style Video Feed (ranking, embedding stores, A/B pipelines)**]() 
- 11. [**Design Time-series Storage for Billions of Events/sec**]() 

| Name | Level | Description | Resource |
|:---|:---|:---|:---|
| 7. [**Design a ChatGPT-like LLM Serving Architecture (autoscaling, quantization, fallback)**](link) | Hard | Model inference, token limits, cost optimization. | 2026 тренд: AI-heavy; добавь agentic tools. |
| 8. [**Design a Global Social Graph Storage (Facebook-scale)**](link) | Hard | Graph DBs, sharding, privacy. | Meta-specific. |
| 9. [**Design a Global User Profile Service with caching & write coalescing**](link) | Medium | Multi-region sync, GDPR compliance. | Из твоего. |
| 10. [**Design a Large-scale Feature Store (for ML training/serving)**](link) | Advanced | Offline/online features, vector embeddings. | ML infra. |
| 11. [**Design a Low-Latency Multi-region Notification System**](link) | Medium | Push, fan-out, reliability. | Apple/Google push. |
| 12. [**Design a Multi-region Active-Active Newsfeed System**](link) | Hard | Ranking, A/B tests. | Facebook Newsfeed. |
| 13. [**Design a Real-time Inference Service (10M RPS, <10 ms latency)**](link) | Advanced | GPU scaling, quantization. | Tesla AI (autopilot inference). |
| 14. [**Design a Real-time Personalized Recommendation Engine with Online Learning**](link) | Hard | Two-tower models, feedback loops. | Netflix/Amazon. |
| 15. [**Design Privacy-Aware User Data Storage (GDPR, erasure, auditing, differential privacy)**](link) | Advanced | Encryption, access logs. | 2026 фокус: privacy. |
| 16. [**Design TikTok-Style Video Feed (ranking, embedding stores, A/B pipelines)**](link) | Hard | Short-form content, ML ranking. | ByteDance-style. |
| 17. [**Design Time-series Storage for Billions of Events/sec**](link) | Advanced | InfluxDB-like, compression. | Monitoring tools. |
| 18. [**Design Instagram**](link) | Medium | Feeds, stories, media upload. | FAANG staple<br>igotanoffer.com |
| 19. [**Design YouTube or Netflix**](link) | Medium | Streaming, recommendations. | Content delivery<br>tryexponent.com |

--- 

## [C Distributed Storage, Consistency & Databases]()

- 1. [**Design a Columnar Storage System for Analytics (Parquet/BigQuery-like)**]() 
- 2. [**Design a Distributed Key-Value Store (Dynamo/Bigtable-style)**]() 
- 3. [**Design a Distributed Metadata Service (e.g., HDFS NameNode or etcd)**]() 
- 4. [**Design a Global Locking/Coordination Service (Zookeeper/Chubby-like)**]() 
- 5. [**Design a Large-scale Blob Storage (S3-like with multi-region replication)**]() 
- 6. [**Design a Multi-zone Conflict-free Replication System (CRDT-based)**]() 
- 7. [**Design a Sharded SQL Database with Online Resharding**]() 
- 8. [**Design a Vector Database for Semantic Search/ANN (Pinecone/Weaviate-scale)**]() 
- 9. [**Design Log-as-Database Architecture (Kafka/Pravega/Lakehouse)**]() 

| Topic | Сложность | Ключевые аспекты | Ресурсы/Примечания |
|:---|:---|:---|---|
| 1. [**Design a Columnar Storage System for Analytics (Parquet/BigQuery-like)**](link) | Advanced | Query optimization, partitioning. | Big Data. |
| 2. [**Design a Distributed Key-Value Store (Dynamo/Bigtable-style)**](link) | Hard | CAP trade-offs, replication. | Amazon Dynamo. |
| 3. [**Design a Distributed Metadata Service (e.g., HDFS NameNode or etcd)**](link) | Advanced | High availability, consensus. | Из твоего. |
| 4. [**Design a Global Locking/Coordination Service (Zookeeper/Chubby-like)**](link) | Hard | Leader election, watches. | Google Chubby. |
| 5. [**Design a Large-scale Blob Storage (S3-like with multi-region replication)**](link) | Hard | Durability, geo-replication. | AWS S3. |
| 6. [**Design a Multi-zone Conflict-free Replication System (CRDT-based)**](link) | Advanced | Eventual consistency. | Из твоего. |
| 7. [**Design a Sharded SQL Database with Online Resharding**](link) | Hard | Zero-downtime migrations. | Vitess-style. |
| 8. [**Design a Vector Database for Semantic Search/ANN (Pinecone/Weaviate-scale)**](link) | Advanced | HNSW indexes, hybrid search. | AI search 2026. |
| 9. [**Design Log-as-Database Architecture (Kafka/Pravega/Lakehouse)**](link) | Hard | Stream processing, immutability. | Delta Lake. |
| 10. [**Design a Wide Column Database**](link) | Hard | Cassandra-like, sparse data. | Из твоего 1-го. |
| 11. [**Design a Distributed File System**](link) | Hard | Fault tolerance, replication. | HDFS. |

---

## [D E-commerce, Fintech & Mission-Critical Systems]() 

- 1. [**Design a Fraud Detection System (event ingestion + real-time ML scoring)**]() 
- 2. [**Design a Global Shopping Cart Service (session replication, conflict resolution)**]() 
- 3. [**Design a High-throughput Search Auction System (AdWords/Meta Ads)**]() 
- 4. [**Design an Idempotent Payment Processing Pipeline (Stripe-like)**]() 
- 5. [**Design an Inventory Consistency System with distributed locking/sagas**]() 
- 6. [**Design a High-throughput Search Auction System (AdWords/Meta Ads)**]() 

| Тема | Сложность | Ключевые аспекты | Ресурсы/Примечания |
|:---|:---|:---|---|
| 1. [**Design a Fraud Detection System (event ingestion + real-time ML scoring)**](link) | Hard | Anomaly detection, rules+ML. | Stripe/ML. |
| 2. [**Design a Global Shopping Cart Service (session replication, conflict resolution)**](link) | Medium | Distributed sessions, sagas. | Amazon Cart. |
| 3. [**Design a High-throughput Search Auction System (AdWords/Meta Ads)**](link) | Hard | Real-time bidding, latency <50ms. | Google Ads. |
| 4. [**Design an Idempotent Payment Processing Pipeline (Stripe-like)**](link) | Hard | Transactions, retries. | Fintech. |
| 5. [**Design an Inventory Consistency System with distributed locking/sagas**](link) | Hard | 2PC or sagas for consistency. | E-commerce. |
| 6. [**Design an E-commerce Service**](link) | Medium | Orders, inventory, recommendations. | Shopify. |
| 7. [**Design an Online Payment Service**](link) | Medium | PCI compliance, gateways. | PayPal. |
| 8. [**Design an Auction System**](link) | Hard | Bidding, winner determination. | eBay. |
---

## [E High-Performance Distributed Compute & ML Infra]()

- 1. [**Design a Batch Compute Engine (MapReduce/Spark-like)**]() 
- 2. [**Design a GPU/TPU Orchestration Cluster (for ML training/inference)**]() 
- 3. [**Design a Priority-based Job Scheduler with backfilling & preemption**]() 
- 4. [**Design a Stream Processing Pipeline (Flink/Spark Streaming)**]()
- 5. [**Design a Workflow Orchestration System (Airflow/Dagster-like)**]()
- 6. [**Design an Autoscaling System for Real-time Compute**]()
- 7. [**Design an LLM Fine-Tuning & Serving Platform (LoRA, quantization-aware training)**]()

| Тема | Сложность | Ключевые аспекты | Ресурсы/Примечания |
|:---|:---|:---|---|
| 1. [**Design a Batch Compute Engine (MapReduce/Spark-like)**](link) | Advanced | Job DAGs, fault recovery. | Hadoop. |
| 2. [**Design a GPU/TPU Orchestration Cluster (for ML training/inference)**](link) | Advanced | Scheduling, multi-tenancy. | Kubernetes+GPUs. |
| 3. [**Design a Priority-based Job Scheduler with backfilling & preemption**](link) | Advanced | Fairness, SLAs. | Borg/K8s. |
| 4. [**Design a Stream Processing Pipeline (Flink/Spark Streaming)**](link) | Advanced | Windowing, state management. | Kafka integration. |
| 5. [**Design a Workflow Orchestration System (Airflow/Dagster-like)**](link) | Medium | DAGs, retries. | ETL/ML pipelines. |
| 6. [**Design an Autoscaling System for Real-time Compute**](link) | Advanced | Metrics-based, predictive. | AWS ASG. |
| 7. [**Design an LLM Fine-Tuning & Serving Platform (LoRA, quantization-aware training)**](link) | Hard | Distributed training, serving. | 2026 AI. |
| 8. [**Design a Big Data Processing Pipeline**](link) | Advanced | ETL at scale. | Из твоего 1-го. |
| 9. [**Design a High-Performance Computing Cluster**](link) | Advanced | MPI, low-latency networks. | SpaceX simulations. |

---

## [F Networking, Reliability & Multi-Region Ops]()

- 1. [**Design a Disaster Recovery System with RPO/RTO guarantees**]()
- 2. [**Design a Global Edge Compute Platform (Cloudflare Workers/AWS Lambda@Edge)**]()
- 3. [**Design a Multi-region Load Balancing System (GSLB with anycast)**]()
- 4. [**Design a Multi-tenant SaaS System with noisy-neighbor isolation**]()
- 5. [**Design a Real-time CDN with edge compute and cache invalidation**]()
- 6. [**Design a Zero-downtime Deployment Framework (canary, blue/green, feature flags)**]()
- 7. [**Design an API Gateway with authN/authZ, rate-limits, quotas & observability**]()

| Тема | Сложность | Ключевые аспекты | Ресурсы/Примечания |
|:---|:---|:---|---|
| 1. [**Design a Disaster Recovery System with RPO/RTO guarantees**](link) | Advanced | Backups, failover. | AWS DR. |
| 2. [**Design a Global Edge Compute Platform (Cloudflare Workers/AWS Lambda@Edge)**](link) | Advanced | Serverless, anycast. | 2026 edge. |
| 3. [**Design a Multi-region Load Balancing System (GSLB with anycast)**](link) | Hard | Geo-routing, health checks. | Akamai. |
| 4. [**Design a Multi-tenant SaaS System with noisy-neighbor isolation**](link) | Advanced | Namespaces, quotas. | Salesforce. |
| 5. [**Design a Real-time CDN with edge compute and cache invalidation**](link) | Hard | Purging, prefetching. | Fastly. |
| 6. [**Design a Zero-downtime Deployment Framework (canary, blue/green, feature flags)**](link) | Medium | Rollouts, monitoring. | CI/CD. |
| 7. [**Design an API Gateway with authN/authZ, rate-limits, quotas & observability**](link) | Medium | Envoy/Istio. | Microservices. |
| 8. [**Design a Hybrid Cloud Infrastructure**](link) | Advanced | On-prem + cloud sync. | Из твоего 1-го. |
| 9. [**Design a Virtualization System**](link) | Advanced | Hypervisors, VMs. | VMware. |
---

## [G Search, Recommendations & Personalization (AI-Heavy in 2025)]()

- 1. [**Design a Large-scale Embedding Store (for ranking/retrieval models)**]()
- 2. [**Design a Real-time Clickstream Analytics System**]()
- 3. [**Design a Recommendation Ranking Pipeline (retrieval + scoring + rerank + two-tower)**]()
- 4. [**Design a Vector Search Engine (semantic search / ANN with hybrid retrieval)**]()
- 5. [**Design an Agentic AI System (tool calling, multi-agent orchestration, memory)**]()
- 6. [**Design Google Search Autocomplete (prefix index + ranking + personalization)**]()

| Тема | Сложность | Ключевые аспекты | Ресурсы/Примечания |
|:---|:---|:---|---|
| 1. [**Design a Large-scale Embedding Store (for ranking/retrieval models)**](link) | Advanced | FAISS, vector indexes. | ML serving. |
| 2. [**Design a Real-time Clickstream Analytics System**](link) | Medium | Event processing, dashboards. | Kafka+ELK. |
| 3. [**Design a Recommendation Ranking Pipeline (retrieval + scoring + rerank + two-tower)**](link) | Hard | Online/offline. | Amazon Recs. |
| 4. [**Design a Vector Search Engine (semantic search / ANN with hybrid retrieval)**](link) | Advanced | Pinecone-like. | RAG for AI. |
| 5. [**Design an Agentic AI System (tool calling, multi-agent orchestration, memory)**](link) | Hard | LangChain-style. | 2026 тренд. |
| 6. [**Design Google Search Autocomplete (prefix index + ranking + personalization)**](link) | Medium | Trie + ML. | Google. |
| 7. [**Design Typeahead Suggestion**](link) | Medium | Predictive search. | Из твоего 1-го. |
| 8. [**Design Google Search**](link) | Hard | Crawling, indexing, ranking. | Full search engine. |

---

## [H. SpaceX/Tesla-Specific Flavor (High-Reliability & Real-Time)]()

| Тема | Сложность | Ключевые аспекты | Ресурсы/Примечания |
|:---|:---|:---|---|
| 1. [**Design a Charging Station Network**](link) | Medium | Load balancing, availability. | Tesla Superchargers |
| 2. [**Design a Fleet Management System for Autonomous Vehicles**](link) | Hard | Tracking, routing, charging. | Tesla-specific |
| 3. [**Design a Telemetry Pipeline for Millions of Sensors (rocket/vehicle → ground/cloud)**](link) | Advanced | Low-latency ingest, anomaly detection. | SpaceX: Stream processing |
| 4. [**Design a Low-Latency Over-the-Air (OTA) Update System for Fleet**](link) | Advanced | Delta updates, rollback safety. | Tesla vehicles |
| 5. [**Design a Safety-Critical Real-Time Control System (redundancy, voting)**](link) | Hard | Fault tolerance, real-time OS. | SpaceX engines |
| 6. [**Design Satellite Constellation Management (inter-satellite links, routing)**](link) | Advanced | Starlink-like, dynamic routing. | SpaceX |
| 7. [**Design Edge AI Inference on Vehicles/Satellites (quantized models, fallback)**](link) | Advanced | On-device ML, cloud fallback. | Tesla Autopilot |


