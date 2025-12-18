1. Architecture & Internals
   Explain the PostgreSQL multi-process architecture. How does it differ from a multi-threaded model?

What is the role of the Postmaster process and background workers?

How does Shared Buffers work, and how does it interact with the OS Page Cache?

Describe the Write-Ahead Logging (WAL) mechanism. Why is it critical for ACID compliance?

What are Checkpoints, and how do they impact system performance during high write loads?

Explain the concept of TOAST (The Oversized-Attribute Storage Technique). How does Postgres store large field values?

What is a CTID, and how does it change during an UPDATE operation?

Describe the purpose of the Free Space Map (FSM) and Visibility Map (VM).

What is the difference between pg_xlog (or pg_wal in newer versions) and pg_clog (now pg_xact)?

How does PostgreSQL handle Catalog Bloat, and why is it dangerous?

2. Concurrency & Locking
   How does MVCC (Multi-Version Concurrency Control) work in PostgreSQL?

Explain the four Transaction Isolation Levels supported by Postgres. Which one is the default?

What is Transaction ID (XID) Wraparound, and how does the autovacuum process prevent it?

Compare Advisory Locks with traditional row/table locks. When would you use them at the application level?

What are Deadlocks, and how does Postgres detect and resolve them?

Explain the difference between FOR UPDATE and FOR SHARE locking clauses.

What is Predicate Locking in the context of Serializable isolation?

How does SKIP LOCKED help in implementing high-throughput job queues?

3. Performance Tuning & Indexing
   Explain how the Query Planner/Optimizer works. How does it use statistics?

What is the difference between B-Tree, GIN, GIST, and BRIN indexes? Provide a use case for each.

How do you identify Index Bloat, and how do you fix it without downtime (REINDEX CONCURRENTLY)?

When would you use a Partial Index vs. a Functional Index?

Describe Covering Indexes (the INCLUDE clause) and how they enable Index-Only Scans.

How do you interpret the output of EXPLAIN (ANALYZE, BUFFERS)? What does "Shared Hit" signify?

What is Parallel Querying? What parameters (e.g., max_parallel_workers) control it?

How do you tune work_mem, maintenance_work_mem, and effective_cache_size?

Describe the impact of Fillfactor on table and index performance.

How would you debug a "slow" query that only performs poorly in production but not in staging?

4. Advanced SQL & Schema Design
   Compare Recursive CTEs with standard joins for hierarchical data (e.g., an org chart).

What are Window Functions? Explain the difference between RANK(), DENSE_RANK(), and ROW_NUMBER().

When should you use JSONB over a structured relational schema? What are the indexing trade-offs?

Explain Table Partitioning (Declarative). Contrast Range, List, and Hash partitioning.

What is a Materialized View, and how do you handle the "refresh" bottleneck?

Describe Foreign Data Wrappers (FDW). How can they be used for cross-database querying?

What are Lateral Joins, and how do they differ from standard correlated subqueries?

Explain the usage of the FILTER clause in aggregate functions.

5. High Availability & Administration
   Explain the difference between Physical Replication and Logical Replication.

What is Synchronous vs. Asynchronous Replication, and how does it affect the RPO/RTO?

How do you perform Point-in-Time Recovery (PITR)?

What is Streaming Replication, and how do you monitor the replication lag?

Describe the role of a Connection Pooler like PgBouncer or Pgpool-II. Why is it necessary?

How do you handle a Major Version Upgrade with minimal downtime?

What is the difference between VACUUM, VACUUM ANALYZE, and VACUUM FULL?

How would you design a multi-tenant database: one schema per tenant or one database per tenant?

What are Tablespaces, and when would you move an index to a different physical disk?

6. Problem Solving & Scenarios
   A table has 1 billion rows. You need to add a non-nullable column with a default value. How do you do this safely?

Your autovacuum is not keeping up with a high-update table. What settings do you tweak first?

You see a spike in "Idle in Transaction" connections. How do you investigate and what is the risk to the DB?

A query suddenly switched from using an index to a sequential scan. What could be the cause?

How would you implement a Distributed Lock using only PostgreSQL?