# Database Interview Questions (SQL/NoSQL)

## 📊 SQL (Relational Databases)

### 👶 Junior (Beginner)

#### SQL Basics
- Core concepts
  - What is a database? What is a DBMS?
  - What is a relational database?
  - Main SQL components: DDL, DML, DCL, TCL

#### DDL (Data Definition Language)
- What are CREATE, ALTER, DROP?
- How to create a table? Common data types (INT, VARCHAR, DATE, etc.)
- What is PRIMARY KEY, FOREIGN KEY?
- What are UNIQUE, NOT NULL constraints?

#### DML (Data Manipulation Language)
- Main operators: SELECT, INSERT, UPDATE, DELETE
- SELECT syntax: SELECT, FROM, WHERE, ORDER BY, LIMIT
- How to use WHERE with conditions (=, <>, >, <, BETWEEN, IN, LIKE)?
- What is DISTINCT?

#### JOIN Operations
- What is JOIN? Why is it needed?
- Types of JOIN: INNER JOIN, LEFT JOIN, RIGHT JOIN
- Difference between INNER JOIN and LEFT JOIN
- What is CROSS JOIN?

#### Aggregate Functions
- Main aggregate functions: COUNT(), SUM(), AVG(), MIN(), MAX()
- What is GROUP BY?
- What is HAVING and how is it different from WHERE?
- Can WHERE be used with aggregate functions?

#### Basic Concepts
**Indexes**
- What is an index in a database?
- Why are indexes needed?
- How to create an index?
- What is a composite index?

**Transactions**
- What is a transaction?
- Transaction ACID properties:
  - Atomicity
  - Consistency
  - Isolation
  - Durability
- Basic operators: BEGIN, COMMIT, ROLLBACK

**Normalization**
- What is database normalization?
- First Normal Form (1NF)
- Second Normal Form (2NF)
- Third Normal Form (3NF)
- Why is normalization needed?

---

### 🧑 Middle (Intermediate)

#### Advanced SQL
**Complex Queries**
- Subqueries: correlated and non-correlated
- Common Table Expressions (CTE)
- Recursive queries
- Window Functions:
  - ROW_NUMBER(), RANK(), DENSE_RANK()
  - LAG(), LEAD()
  - SUM() OVER(), AVG() OVER()
- What is PARTITION BY in window functions?

**Query Optimization**
- What is EXPLAIN and how to use it?
- How to read a query execution plan?
- What is a full table scan and why is it bad?
- How indexes work with different queries

**Indexes (Advanced)**
- Types: B-tree, Hash, GiST, GIN, BRIN
- How does a B-tree index work?
- What is a covering index?
- When is an index not used?
- What is a functional index?
- Index problems: fragmentation, bloat

**Transactions and Isolation**
- Transaction isolation levels:
  - READ UNCOMMITTED
  - READ COMMITTED
  - REPEATABLE READ
  - SERIALIZABLE
- Concurrency problems:
  - Dirty Read
  - Non-repeatable Read
  - Phantom Read
- What are locks? Types of locks
- Deadlock – what is it and how to avoid?

**Performance**
- How to measure query performance?
- What is a query plan and how to analyze it?
- JOIN optimization methods
- Materialized Views
- Partitioning

**Database Design**
- Denormalization: when and why?
- Sharding vs Replication
- Database design patterns
- Schema optimization for specific workloads

**Working with Specific DBMS**
- **PostgreSQL specifics**
  - Advanced data types (JSON, JSONB, ARRAY, HSTORE)
  - Table inheritance
  - Extensions
  - MVCC (Multi-Version Concurrency Control)
  - VACUUM and AUTOVACUUM
  - Prepared statements

- **MySQL specifics**
  - Storage engines: InnoDB vs MyISAM
  - Query optimization in MySQL
  - Replication in MySQL
  - System variables and tuning

---

### 🧙‍♂️ Senior (Advanced)

#### Architecture and Scaling
- Database scaling
  - Vertical vs Horizontal scaling
  - Sharding strategies (range, hash, directory-based)
  - Replication: master-slave, master-master, multi-master
  - Read replicas
- Caching strategies: write-through, write-back, cache-aside

**High Availability and Fault Tolerance**
- What is High Availability (HA)?
- Failover and failback strategies
- Database clustering
- Backups and recovery: full, incremental, differential
- Point-in-Time Recovery (PITR)

**Distributed Transactions**
- Two-Phase Commit (2PC)
- Saga pattern
- CAP theorem and its application
- PACELC theorem
- Distributed locks

#### Performance and Monitoring
**Advanced Optimization**
- Optimization for high-load systems
- Connection pooling
- Query rewriting and application-level optimization
- Using materialized views
- OLAP vs OLTP optimization

**Monitoring and Diagnostics**
- Key performance indicators (KPIs) for databases
- Real-time performance monitoring
- Slow query analysis
- Using pg_stat_statements, performance_schema
- Alerting and automated response

#### Security
- SQL Injection and protection methods
- Roles and privileges
- Data encryption: at-rest and in-transit
- Audit and compliance
- Row Level Security (RLS)

#### Patterns and Best Practices
**Schema Migrations**
- Migration management tools
- Zero-downtime migrations
- Backward compatible changes
- Strategies for updating large tables

**Data Modeling for Scaling**
- Time-series data
- Graph data
- Geo data
- Sharding large tables
- Data archiving

**Distributed Databases**
- Consistency in distributed systems
- Consensus algorithms (Paxos, Raft)
- Eventual consistency
- CRDT (Conflict-Free Replicated Data Types)
