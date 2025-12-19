### Updated Advanced System Design Topics for FAANG/MANGA, SpaceX & Tesla Interviews (2025 Edition)

#### A. Consumer-Scale Systems & Real-Time Features

- Design a Global Social Graph Storage (Facebook-scale)
- Design a Multi-region Active-Active Newsfeed System
- Design Global Messaging with Exactly-Once Semantics
- **Design TikTok-Style Video Feed** (ranking, embedding stores, A/B pipelines)
- Design a Low-Latency Multi-region Notification System
- Design Time-series Storage for Billions of Events/sec
- Design a Real-time Inference Service (10M RPS, <10 ms latency)
- **Design a ChatGPT-like LLM Serving Architecture** (incl. autoscaling, quantization, fallback)
- Design a Large-scale Feature Store (for ML training/serving)
- Design Privacy-Aware User Data Storage (GDPR, erasure, auditing, differential privacy)
- Design a Global User Profile Service with caching & write coalescing
- **NEW 2025**: Design a Real-time Personalized Recommendation Engine with Online Learning

#### B. Distributed Storage, Consistency & Databases

- Design a Distributed Key-Value Store (Dynamo/Bigtable-style)
- Design a Multi-zone Conflict-free Replication System (CRDT-based)
- Design a Columnar Storage System for Analytics (Parquet/BigQuery-like)
- Design a Large-scale Blob Storage (S3-like with multi-region replication)
- Design a Sharded SQL Database with Online Resharding
- Design a Distributed Metadata Service (e.g., HDFS NameNode or etcd)
- Design Log-as-Database Architecture (Kafka/Pravega/Lakehouse)
- Design a Global Locking/Coordination Service (Zookeeper/Chubby-like)
- **NEW 2025**: Design a Vector Database for Semantic Search/ANN (Pinecone/Weaviate-scale)

#### C. High-Performance Distributed Compute & ML Infra

- Design a Stream Processing Pipeline (Flink/Spark Streaming)
- Design a Workflow Orchestration System (Airflow/Dagster-like)
- Design an Autoscaling System for Real-time Compute
- Design a Batch Compute Engine (MapReduce/Spark-like)
- Design a GPU/TPU Orchestration Cluster (for ML training/inference)
- Design a Priority-based Job Scheduler with backfilling & preemption
- **NEW 2025**: Design an LLM Fine-Tuning & Serving Platform (LoRA, quantization-aware training)

#### D. Networking, Reliability & Multi-Region Ops

- Design a Multi-region Load Balancing System (GSLB with anycast)
- Design a Zero-downtime Deployment Framework (canary, blue/green, feature flags)
- Design a Disaster Recovery System with RPO/RTO guarantees
- Design an API Gateway with authN/authZ, rate-limits, quotas & observability
- Design a Multi-tenant SaaS System with noisy-neighbor isolation
- Design a Real-time CDN with edge compute and cache invalidation
- **NEW 2025**: Design a Global Edge Compute Platform (Cloudflare Workers/AWS Lambda@Edge)

#### E. E-commerce, Fintech & Mission-Critical Systems

- Design an Idempotent Payment Processing Pipeline (Stripe-like)
- Design a Fraud Detection System (event ingestion + real-time ML scoring)
- Design a Global Shopping Cart Service (session replication, conflict resolution)
- Design a High-throughput Search Auction System (AdWords/Meta Ads)
- Design an Inventory Consistency System with distributed locking/sagas

#### F. Search, Recommendations & Personalization (AI-Heavy in 2025)

- Design a Vector Search Engine (semantic search / ANN with hybrid retrieval)
- Design Google Search Autocomplete (prefix index + ranking + personalization)
- Design a Large-scale Embedding Store (for ranking/retrieval models)
- Design a Recommendation Ranking Pipeline (retrieval + scoring + rerank + two-tower)
- Design a Real-time Clickstream Analytics System
- **NEW 2025**: Design an Agentic AI System (tool calling, multi-agent orchestration, memory)

#### G. Collaboration, Real-Time Sync & Editing

- Design a Google Docs-like Real-time Collaboration System (CRDT/OT with conflict resolution)
- Design a Multiplayer Game Backend (state sync, tick model, lag compensation)
- Design a Real-time Presence and Status Service
- Design a Realtime Document Versioning System with merge semantics

#### H. SpaceX/Tesla-Specific Flavor (High-Reliability & Real-Time)

- Design a Telemetry Pipeline for Millions of Sensors (rocket/vehicle → ground/cloud)
- Design a Low-Latency Over-the-Air (OTA) Update System for Fleet
- Design a Safety-Critical Real-Time Control System (redundancy, voting)
- Design Satellite Constellation Management (inter-satellite links, routing)
- **NEW 2025**: Design Edge AI Inference on Vehicles/Satellites (quantized models, fallback)