# System Design Interview Prep for FAANG/MANGA, SpaceX & Tesla (2026–2027 Edition)

A prioritized roadmap, not just a catalog. Start with the **Answer Framework** and **Fundamentals**, then work through **Tier 1** (must-know). After that, pick the **Tier 2** domains that match your target company and role. **Tier 3** is optional and only useful for niche infra roles.

**Status legend:** ⬜ not started · 🟨 in progress · ✅ done (can explain end-to-end in 45 min) · 🎤 done as a mock interview

**Levels:** Easy · Medium · Hard · Advanced

---

## Table of Contents

| Section | Topics | Priority |
|:--------|:------:|:---------|
| [0. Answer Framework](#0-answer-framework) | — | Must |
| [1. Fundamentals](#1-fundamentals) | 16 | Must |
| [2. Back-of-the-Envelope Numbers](#2-back-of-the-envelope-numbers) | — | Must |
| [3. Tier 1 Core Classics](#3-tier-1-core-classics) | 20 | Must |
| [4. Tier 2 Domain Tracks](#4-tier-2-domain-tracks) | 70 | By target company |
| [4A. Collaboration and Real-Time Sync](#4a-collaboration-and-real-time-sync) | 5 | Google, Figma, Notion, Meta |
| [4B. Consumer-Scale Social and Media](#4b-consumer-scale-social-and-media) | 6 | Meta, TikTok, Netflix |
| [4C. Distributed Storage and Databases](#4c-distributed-storage-and-databases) | 7 | Google, Amazon, infra roles |
| [4D. E-commerce, Fintech and Mission-Critical](#4d-e-commerce-fintech-and-mission-critical) | 7 | Amazon, Stripe, banks |
| [4E. AI, LLM and ML Infrastructure](#4e-ai-llm-and-ml-infrastructure) | 11 | All companies in 2026–2027 |
| [4F. Data Processing and Distributed Compute](#4f-data-processing-and-distributed-compute) | 6 | Data/infra roles |
| [4G. Networking, Reliability and Multi-Region](#4g-networking-reliability-and-multi-region) | 9 | Infra, SRE, platform roles |
| [4H. Observability](#4h-observability) | 3 | Infra, SRE |
| [4I. Security and Privacy](#4i-security-and-privacy) | 4 | All, especially fintech |
| [4J. SpaceX and Tesla Specific](#4j-spacex-and-tesla-specific) | 8 | SpaceX, Tesla, automotive, IoT |
| [5. Tier 3 Niche and Bonus](#5-tier-3-niche-and-bonus) | 7 | Optional |
| [6. 12-Week Study Plan](#6-12-week-study-plan) | — | — |
| [7. Resources](#7-resources) | — | — |

---

## 0. Answer Framework

Use the same structure for every problem. Interviewers score the **process and trade-offs**, not just the final diagram.

| Step | Time (45 min) | What to do |
|:-----|:-------------:|:-----------|
| 1. Requirements | 5 min | Functional requirements (3–5 core features). Non-functional: scale, latency, availability vs consistency, durability. Explicitly state what is out of scope. |
| 2. Estimates | 3–5 min | DAU, QPS (avg and peak), read/write ratio, storage per year, bandwidth. Only do the math that changes design decisions. |
| 3. API | 3–5 min | Core endpoints or events (REST / gRPC / WebSocket), request and response shape, pagination, idempotency keys. |
| 4. Data Model | 3–5 min | Entities, access patterns, SQL vs NoSQL choice and why, partition key. |
| 5. High-Level Design | 10 min | Clients, LB, services, DBs, caches, queues. Walk through the main read and write paths. |
| 6. Deep Dives | 10–15 min | 2–3 hardest parts: hot keys, fan-out, consistency, failure handling. Let the interviewer steer. |
| 7. Wrap-up | 2–3 min | Bottlenecks, single points of failure, monitoring, what you would do at 10x scale. |

**Level expectations:**
- **L4 / Mid:** working end-to-end design, reasonable component choices.
- **L5 / Senior:** drives the conversation, explains trade-offs proactively, handles failures and scaling.
- **L6+ / Staff:** identifies the real problem, compares multiple architectures, covers operations, cost, migration and org impact.

---

## 1. Fundamentals

Learn these before the problems. Every design is a combination of them.

| # | Topic | What to know | Status |
|:-:|:------|:-------------|:------:|
| 1 | Scaling basics | Vertical vs horizontal scaling, stateless services, load balancing (L4 vs L7, algorithms) | ⬜ |
| 2 | Caching | Cache-aside, write-through, write-back, TTL, eviction (LRU/LFU), invalidation, cache stampede, hot keys | ⬜ |
| 3 | CDN | Push vs pull CDN, cache keys, purging, edge caching of dynamic content | ⬜ |
| 4 | Databases | SQL vs NoSQL, indexes (B-tree vs LSM-tree), transactions, isolation levels | ⬜ |
| 5 | Replication | Leader-follower, multi-leader, leaderless, sync vs async, replication lag, read-your-writes | ⬜ |
| 6 | Partitioning | Range vs hash sharding, consistent hashing, rebalancing, hot partitions | ⬜ |
| 7 | Consistency models | Strong, eventual, causal, linearizability, CAP and PACELC | ⬜ |
| 8 | Consensus | Raft/Paxos basics, quorums, leader election, split brain, fencing tokens | ⬜ |
| 9 | Messaging | Queues vs logs (SQS vs Kafka), at-most/at-least/exactly-once, ordering, DLQ, backpressure | ⬜ |
| 10 | Idempotency | Idempotency keys, deduplication, retries with exponential backoff and jitter | ⬜ |
| 11 | Distributed transactions | 2PC, saga (choreography vs orchestration), outbox pattern, CDC | ⬜ |
| 12 | Communication | REST, gRPC, GraphQL, WebSocket, SSE, long polling, webhooks | ⬜ |
| 13 | Reliability patterns | Timeouts, circuit breaker, bulkhead, rate limiting, graceful degradation | ⬜ |
| 14 | Storage types | Block vs file vs object storage, blob + metadata split | ⬜ |
| 15 | Search basics | Inverted index, tokenization, ranking, vector search (ANN, HNSW) | ⬜ |
| 16 | Observability basics | Metrics, logs, traces, SLI/SLO/error budgets | ⬜ |

---

## 2. Back-of-the-Envelope Numbers

| Quantity | Approx value |
|:---------|:-------------|
| L1 cache / RAM reference | ~1 ns / ~100 ns |
| Read 1 MB sequentially from RAM | ~10 µs |
| SSD random read | ~100 µs |
| Round trip within a datacenter | ~0.5 ms |
| Read 1 MB sequentially from SSD | ~1 ms |
| Round trip across continents | ~100–150 ms |
| Seconds per day | ~86,400 ≈ 10^5 |
| 1M requests/day | ≈ 12 QPS |
| 100M requests/day | ≈ 1,200 QPS (peak ≈ 2–3x) |
| Single well-tuned DB node | ~1K–10K write QPS, more for reads |
| Single Redis node | ~100K+ ops/sec |
| Powers of 2 | 2^10 ≈ 1K, 2^20 ≈ 1M, 2^30 ≈ 1B, 2^40 ≈ 1T |

---

## 3. Tier 1 Core Classics

The most frequently asked problems. Know every one of these end-to-end.

| # | Topic | Level | Key aspects | Status |
|:-:|:------|:-----:|:------------|:------:|
| 1 | Design a URL Shortener (TinyURL) | Easy | ID generation (base62, counters, Snowflake), redirect latency, caching, analytics | ⬜ |
| 2 | Design a Rate Limiter / Abuse Protection | Medium | Token bucket, sliding window, distributed counters in Redis, per-user and per-IP limits | ⬜ |
| 3 | Design a Distributed Cache | Medium | Consistent hashing, eviction, replication, hot keys, stampede protection | ⬜ |
| 4 | Design a Unique ID Generator | Easy | Snowflake, UUID v7, ordering, clock skew | ⬜ |
| 5 | Design a Notification System (push, email, SMS) | Medium | Fan-out, templates, user preferences, retries, dedup, provider failover | ⬜ |
| 6 | Design a Chat System (WhatsApp / Messenger) | Hard | WebSocket gateways, message ordering, delivery receipts, offline sync, group chat, presence | ⬜ |
| 7 | Design a News Feed (Facebook / X) | Hard | Fan-out on write vs read, hybrid for celebrities, ranking, feed cache | ⬜ |
| 8 | Design Instagram | Medium | Media upload, object storage, CDN, feed, stories with TTL | ⬜ |
| 9 | Design YouTube / Netflix | Hard | Upload pipeline, transcoding, adaptive bitrate (HLS/DASH), CDN, view counts | ⬜ |
| 10 | Design Dropbox / Google Drive | Hard | Chunking, deduplication, sync conflicts, metadata service, notifications | ⬜ |
| 11 | Design Uber / Ride-Sharing | Hard | Location updates, geo index (geohash, quadtree, H3), matching, trip state machine | ⬜ |
| 12 | Design Yelp / Proximity Search | Medium | Geospatial indexing, read-heavy caching, ranking by distance and rating | ⬜ |
| 13 | Design Ticketmaster / Hotel Booking | Hard | Seat reservation with holds, double-booking prevention, virtual waiting queue, flash traffic | ⬜ |
| 14 | Design a Web Crawler | Medium | URL frontier, politeness, dedup (Bloom filter, simhash), robots.txt, distributed workers | ⬜ |
| 15 | Design Typeahead / Search Autocomplete | Medium | Trie or prefix index, top-K per prefix, offline aggregation, personalization | ⬜ |
| 16 | Design Top-K / Leaderboard / Trending | Medium | Count-min sketch, heavy hitters, sliding windows, Redis sorted sets | ⬜ |
| 17 | Design an Ad Click Aggregator | Hard | Stream processing, exactly-once counting, late events, reconciliation with batch | ⬜ |
| 18 | Design a Distributed Message Queue (Kafka-like) | Hard | Partitioned log, consumer groups, offsets, replication (ISR), retention | ⬜ |
| 19 | Design a Distributed Key-Value Store (Dynamo-like) | Hard | Consistent hashing, quorums, vector clocks, hinted handoff, anti-entropy (Merkle trees) | ⬜ |
| 20 | Design an Idempotent Payment System (Stripe-like) | Hard | Idempotency keys, state machine, double-entry ledger, reconciliation, PSP integration | ⬜ |

---

## 4. Tier 2 Domain Tracks

Choose 2–4 tracks based on your target company and role.

### 4A. Collaboration and Real-Time Sync

| # | Topic | Level | Key aspects | Status |
|:-:|:------|:-----:|:------------|:------:|
| 1 | Design Google Docs (real-time collaborative editing) | Advanced | OT vs CRDT (Yjs, Automerge), cursors and presence, offline edits, snapshots, version history and merge | ⬜ |
| 2 | Design a Collaborative Spreadsheet (Google Sheets) | Hard | Cell-level conflicts, formula dependency graph, recalculation, sync | ⬜ |
| 3 | Design a Live Comments / Reactions System | Medium | Threaded replies, real-time push, fan-out for viral posts, sampling reactions | ⬜ |
| 4 | Design a Multiplayer Game Backend | Hard | Authoritative server, tick model, lag compensation, matchmaking, anti-cheat | ⬜ |
| 5 | Design a Presence and Typing Indicator Service | Medium | Heartbeats, TTL, pub/sub fan-out, scaling to millions of connections | ⬜ |

### 4B. Consumer-Scale Social and Media

| # | Topic | Level | Key aspects | Status |
|:-:|:------|:-----:|:------------|:------:|
| 1 | Design a TikTok-Style Video Feed | Hard | Candidate generation, ML ranking, prefetching, cold start, CDN | ⬜ |
| 2 | Design Live Streaming with Real-Time Reactions (Twitch / TikTok Live) | Hard | Ingest, low-latency delivery (<2 s), chat fan-out to millions, backpressure | ⬜ |
| 3 | Design a Social Graph Storage (Meta TAO-like) | Hard | Graph model on sharded stores, caching tiers, friends-of-friends queries | ⬜ |
| 4 | Design a Global User Profile Service | Medium | Multi-region replication, read-your-writes, data residency | ⬜ |
| 5 | Design a Multi-Region Active-Active News Feed | Advanced | Extension of Tier 1 feed: cross-region replication, conflict handling, failover | ⬜ |
| 6 | Design a Photo / Media Processing Pipeline | Medium | Async processing, thumbnails, content moderation, dedup | ⬜ |

### 4C. Distributed Storage and Databases

| # | Topic | Level | Key aspects | Status |
|:-:|:------|:-----:|:------------|:------:|
| 1 | Design Blob / Object Storage (S3-like) | Hard | Metadata vs data plane, erasure coding, durability (11 nines), geo-replication | ⬜ |
| 2 | Design a Distributed Lock / Coordination Service (Chubby / ZooKeeper) | Hard | Consensus, leader election, leases, watches, fencing tokens | ⬜ |
| 3 | Design a Sharded SQL Database with Online Resharding (Vitess-like) | Hard | Query routing, zero-downtime migrations, cross-shard queries | ⬜ |
| 4 | Design a Wide-Column Database (Cassandra / Bigtable) | Hard | LSM trees, SSTables, compaction, tunable consistency | ⬜ |
| 5 | Design a Distributed File System (GFS / HDFS) | Hard | Chunk servers, master metadata, replication, failure recovery | ⬜ |
| 6 | Design Time-Series Storage for Billions of Events/sec | Advanced | Compression (Gorilla), downsampling, retention tiers, high-cardinality | ⬜ |
| 7 | Design a Multi-Region Conflict-Free Replication System | Advanced | CRDTs, version vectors, causal consistency, convergence | ⬜ |

### 4D. E-commerce, Fintech and Mission-Critical

| # | Topic | Level | Key aspects | Status |
|:-:|:------|:-----:|:------------|:------:|
| 1 | Design an E-commerce Platform (Amazon / Shopify) | Medium | Catalog, cart, checkout, order workflow, search | ⬜ |
| 2 | Design a Shopping Cart Service | Medium | Guest vs logged-in carts, merging, availability over consistency | ⬜ |
| 3 | Design an Inventory Consistency System | Hard | Reservations, overselling prevention, sagas, flash sales | ⬜ |
| 4 | Design a Double-Entry Ledger / Wallet | Hard | Immutable entries, balances, strong consistency, audit, reconciliation | ⬜ |
| 5 | Design a Real-Time Fraud Detection System | Hard | Streaming features, low-latency ML scoring, rules engine, feedback loop | ⬜ |
| 6 | Design an Auction System (eBay) | Hard | Bid ordering, concurrency, auction close, anti-sniping | ⬜ |
| 7 | Design an Ad Auction / Real-Time Bidding System | Advanced | <50–100 ms budget, bid requests fan-out, budget pacing | ⬜ |

### 4E. AI, LLM and ML Infrastructure

The fastest-growing interview category for 2026–2027.

| # | Topic | Level | Key aspects | Status |
|:-:|:------|:-----:|:------------|:------:|
| 1 | Design a ChatGPT-like LLM Serving System | Hard | Token streaming, continuous batching, KV cache, GPU scheduling, cost per token | ⬜ |
| 2 | Design an LLM Gateway / Router | Medium | Multi-provider routing, fallbacks, rate limits, semantic caching, cost tracking | ⬜ |
| 3 | Design a RAG Pipeline | Hard | Chunking, embeddings, hybrid retrieval, reranking, context management, freshness | ⬜ |
| 4 | Design an Agentic AI System | Hard | Tool calling, multi-step orchestration, memory, sandboxing, guardrails, observability | ⬜ |
| 5 | Design a Vector Database / Embedding Store | Advanced | HNSW / IVF indexes, filtering, sharding, hybrid search, updates | ⬜ |
| 6 | Design an LLM Evaluation and Guardrails Platform | Medium | Offline evals, online monitoring, safety filters, regression detection | ⬜ |
| 7 | Design a Recommendation System (Netflix / Amazon) | Hard | Retrieval (two-tower), ranking, re-ranking, real-time feedback, A/B tests | ⬜ |
| 8 | Design a Feature Store | Advanced | Offline/online consistency, point-in-time correctness, low-latency serving | ⬜ |
| 9 | Design a Real-Time ML Inference Service (10M RPS, <10 ms) | Advanced | Model serving, batching, quantization, autoscaling GPUs | ⬜ |
| 10 | Design an LLM Fine-Tuning and Training Platform | Advanced | Distributed training, checkpointing, job scheduling, dataset versioning | ⬜ |
| 11 | Design a GPU Cluster Orchestration System | Advanced | Gang scheduling, multi-tenancy, preemption, utilization | ⬜ |

### 4F. Data Processing and Distributed Compute

| # | Topic | Level | Key aspects | Status |
|:-:|:------|:-----:|:------------|:------:|
| 1 | Design a Stream Processing Pipeline (Flink-like) | Advanced | Windowing, watermarks, state, checkpointing, exactly-once | ⬜ |
| 2 | Design a Batch Compute Engine (MapReduce / Spark) | Advanced | DAG execution, shuffle, data locality, fault recovery | ⬜ |
| 3 | Design a Distributed Job Scheduler (cron at scale) | Hard | Priorities, fairness, preemption, exactly-once triggering, retries | ⬜ |
| 4 | Design a Workflow Orchestration System (Airflow / Temporal) | Medium | DAGs, durable execution, retries, backfills | ⬜ |
| 5 | Design a Real-Time Clickstream Analytics System | Medium | Event ingestion, aggregation, OLAP storage, dashboards | ⬜ |
| 6 | Design an Autoscaling System | Hard | Metrics-based and predictive scaling, cooldowns, oscillation | ⬜ |

### 4G. Networking, Reliability and Multi-Region

| # | Topic | Level | Key aspects | Status |
|:-:|:------|:-----:|:------------|:------:|
| 1 | Design an API Gateway | Medium | AuthN/AuthZ, rate limits, routing, quotas, observability | ⬜ |
| 2 | Design a Global Load Balancing System | Hard | GeoDNS, Anycast, health checks, latency-based routing, failover | ⬜ |
| 3 | Design a CDN with Edge Compute | Hard | Cache hierarchy, purging, edge functions, TLS termination | ⬜ |
| 4 | Design a Service Discovery System | Medium | Registry, health checks, client vs server-side discovery | ⬜ |
| 5 | Design a Disaster Recovery Strategy | Advanced | RPO/RTO, backups, active-active vs active-passive, failover drills | ⬜ |
| 6 | Design a Zero-Downtime Deployment System | Medium | Blue-green, canary, rollback, schema migrations | ⬜ |
| 7 | Design a Feature Flag and A/B Testing Platform | Medium | Targeting, consistent bucketing, low-latency evaluation, statistics | ⬜ |
| 8 | Design a Multi-Tenant SaaS Platform | Advanced | Isolation models, noisy neighbors, quotas, per-tenant data | ⬜ |
| 9 | Design a Serverless / Edge Compute Platform | Advanced | Isolation (V8 isolates, microVMs), cold starts, placement | ⬜ |

### 4H. Observability

| # | Topic | Level | Key aspects | Status |
|:-:|:------|:-----:|:------------|:------:|
| 1 | Design a Metrics and Alerting System | Medium | Time-series DB, aggregation, cardinality, alert routing | ⬜ |
| 2 | Design a Log Aggregation Pipeline | Medium | Shippers, indexing, storage tiers, query engine | ⬜ |
| 3 | Design a Distributed Tracing System | Medium | Context propagation, sampling, span storage (OpenTelemetry) | ⬜ |

### 4I. Security and Privacy

| # | Topic | Level | Key aspects | Status |
|:-:|:------|:-----:|:------------|:------:|
| 1 | Design an Authentication and Authorization System | Hard | OAuth2, OIDC, sessions vs tokens, RBAC/ABAC, token revocation | ⬜ |
| 2 | Design a Secrets Management System (Vault-like) | Hard | Encryption, key rotation, dynamic secrets, audit logs | ⬜ |
| 3 | Design Privacy-Aware User Data Storage | Advanced | Encryption at rest, access logs, GDPR deletion, data residency | ⬜ |
| 4 | Design a Data Anonymization / Differential Privacy Pipeline | Advanced | k-anonymity, differential privacy budgets, compliance | ⬜ |

### 4J. SpaceX and Tesla Specific

| # | Topic | Level | Key aspects | Status |
|:-:|:------|:-----:|:------------|:------:|
| 1 | Design a Telemetry Pipeline for Millions of Sensors | Advanced | High-throughput ingest, compression, anomaly detection, storage tiers | ⬜ |
| 2 | Design an OTA Update System for a Vehicle Fleet | Advanced | Delta updates, staged rollout, atomic install, rollback safety | ⬜ |
| 3 | Design a Fleet Management System | Hard | Vehicle tracking, routing, charging, remote commands | ⬜ |
| 4 | Design a Charging Station Network (Superchargers) | Medium | Availability, reservations, load balancing, payments | ⬜ |
| 5 | Design a Safety-Critical Real-Time Control System | Hard | Redundancy, voting (triple modular), determinism, fault tolerance | ⬜ |
| 6 | Design Edge AI Inference on Vehicles / Satellites | Advanced | On-device models, latency budgets, cloud fallback, model versioning | ⬜ |
| 7 | Design a Model Update and Versioning System for Edge Devices | Advanced | Versioning, compatibility, canary on fleet, rollback | ⬜ |
| 8 | Design Satellite Constellation Management (Starlink-like) | Advanced | Dynamic routing, inter-satellite links, handoffs, ground stations | ⬜ |

---

## 5. Tier 3 Niche and Bonus

Rarely asked in general SWE interviews. Study them only for specialized infra roles or after finishing Tier 1–2.

| # | Topic | Level | Key aspects | Status |
|:-:|:------|:-----:|:------------|:------:|
| 1 | Design Google Search (full engine) | Advanced | Crawl, index, rank, serve at web scale | ⬜ |
| 2 | Design a Columnar Analytics Database (BigQuery-like) | Advanced | Columnar formats, vectorized execution, partition pruning | ⬜ |
| 3 | Design a Data Lakehouse (Iceberg / Delta Lake) | Advanced | ACID on object storage, time travel, schema evolution | ⬜ |
| 4 | Design a Log-as-Database Architecture | Advanced | Event sourcing, materialized views, replay | ⬜ |
| 5 | Design a High-Performance Computing Cluster | Advanced | MPI, low-latency interconnects, job placement | ⬜ |
| 6 | Design a Virtualization System | Advanced | Hypervisors, VM scheduling, live migration | ⬜ |
| 7 | Design a Hybrid Cloud Infrastructure | Advanced | On-prem and cloud sync, networking, identity federation | ⬜ |

---

## 6. 12-Week Study Plan

| Weeks | Focus | Output |
|:-----:|:------|:-------|
| 1–2 | Answer Framework + Fundamentals 1–8 | Notes for each fundamental, estimation practice |
| 3 | Fundamentals 9–16 | Notes, 2 easy problems (URL Shortener, ID Generator) |
| 4–6 | Tier 1 problems 2–12 | 3–4 problems per week, 1 mock interview per week |
| 7–8 | Tier 1 problems 13–20 | 4 problems per week, 1–2 mocks per week |
| 9–10 | Tier 2 tracks for your target company | 2 tracks, 4–6 problems, 2 mocks per week |
| 11 | AI/LLM track (4E) | LLM serving, RAG, agentic AI, recommendations |
| 12 | Review and mocks | 3–4 full mock interviews, revisit weak topics |

**Per-problem template.** Save one file per problem with these sections: requirements, estimates, API, data model, diagram, 2–3 deep dives, trade-offs, failure scenarios, what changes at 10x.

**Rule:** a problem is ✅ only when you can explain it out loud in 45 minutes without notes.

---

## 7. Resources

**Books**
- *Designing Data-Intensive Applications*, Martin Kleppmann
- *System Design Interview*, Vol. 1 and 2, Alex Xu
- *Machine Learning System Design Interview*, Ali Aminian and Alex Xu
- *Database Internals*, Alex Petrov

**Courses and practice**
- ByteByteGo
- Hello Interview (guides and mock interviews)
- Grokking the System Design Interview
- Peer mock interviews (Pramp, interviewing.io, friends)

**Classic papers**
- Google: GFS, MapReduce, Bigtable, Chubby, Spanner
- Amazon: Dynamo
- Raft (*In Search of an Understandable Consensus Algorithm*)
- Kafka (LinkedIn)
- Facebook: TAO, Memcache at scale

**Engineering blogs**
- Meta, Netflix, Uber, Discord, Stripe, Cloudflare, Airbnb, Pinterest, LinkedIn engineering blogs
