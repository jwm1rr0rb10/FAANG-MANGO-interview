## Designing a Global Social Graph Storage System (Facebook-Scale)

- This is a classic system design interview question focused on building a distributed, read-heavy graph database capable of handling billions of users, trillions of edges (connections like friendships, likes, follows), and extreme scale (e.g., billions of reads and millions of writes per second). Facebook's real-world implementation is called TAO (The Associations and Objects), a geographically distributed, read-optimized data store for their social graph. I'll explain the full design, drawing from real architectures like TAO, including challenges, trade-offs, and components.

---

## 1. Requirements and Scale

### Functional Requirements:

- `Store nodes (objects):` Users, posts, pages, comments, etc.
- Store edges (associations): Friendships, likes, follows, tags, etc. (directed or undirected, with types and timestamps).
- Operations:
- Read object/edge.
- Query associations (e.g., get all friends of a user, sorted by time).
- Add/delete/update edges (e.g., friend request, like a post).
- Basic traversals (e.g., friends-of-friends for suggestions).


### Non-Functional Requirements (Facebook-scale):

- Scale: 3B+ users, trillions of edges, petabytes of data.
- Read-heavy: >99% reads (e.g., loading news feed pulls graph data).
- Low latency: <100ms for reads globally.
- High availability: Handle failures, geo-distribution.
- Eventual consistency: Strong consistency is hard at scale; read-after-write is often sufficient.
- Global distribution: Low-latency access from multiple regions (e.g., US, Europe, Asia).

### Challenges:

- Graph data has poor locality (random accesses).
- Hotspots (celebrities with millions of followers).
- Geo-replication for low latency.
- Cache invalidation and consistency.

---

## 2. Data Model
### Model the social graph as:

- `Objects (Nodes):` 64-bit unique ID (oid), type (e.g., user, post), key-value data.
- `Associations (Edges):` Source ID (id1), destination ID (id2), type (atype, e.g., friend, like), time, optional data (counters like like count).

- Stored in relational tables (e.g., MySQL):

- `Object table:` (id, type, data).
- `Assoc table:` (id1, atype, id2, time, data) – indexed for efficient reverse lookups (e.g., query all edges from id1, sorted by time descending).

This allows efficient "adjacency list" queries.

---

## 3. High-Level Architecture

- Persistent Storage: Sharded MySQL (or MyRocks for better compression/write amplification).
- Caching Layer: Multi-tier, graph-aware cache (TAO servers).
- Clients: Web/app servers query via TAO client library.
- Geo-Distribution: Full copy per region, with master-slave for writes.

### Key layers:

- Caching Tier: Leader-follower caches per region.
- Storage Tier: Sharded MySQL clusters.

- **Before TAO (old Facebook approach):** Direct MySQL + memcache (look-aside cache). - - - Issues: Cache misses cause thundering herds, complex client logic for invalidations, hard read-after-write.
Here are diagrams illustrating the evolution and TAO architecture:

//here upload pictures with arch

(Pre-TAO architecture: Clients directly hit memcache/MySQL, leading to inconsistencies.)

// here upload picture with arch

## 4. Detailed Components
- **Sharding:**

- Logical shards (many per MySQL instance).
- Objects sharded by ID (consistent hashing or range).
- Associations often co-located with source object for efficiency.

### Caching (TAO-style):

1. Tier-1 (Leaders): Write-through, handle fills from DB, async invalidate followers.
2. Tier-2 (Followers): Read-only replicas for scale-out.
3. Graph-aware: Cache objects + adjacency lists; optimizations like time-ordered edges.

### Read Path:

- Client → Load balancer → TAO leader cache.
- Cache hit → Return.
- Miss → Leader fetches from MySQL → Fills cache → Returns.

### Write Path:

- Client → Leader.
- Leader writes to MySQL (if master region) or forwards to master.
- Async invalidate/propagate to followers and other regions.
- Provides read-after-write (client sees own writes immediately).

### Geo-Replication:

- One full copy per region (low inter-region latency).
- Writes go to master region → Async replicate to slaves.
- Eventual consistency across regions.

### Fault Tolerance:

- MySQL replication (master-slave).
- Cache failover.
- Refills on failure.

### Optimizations:

- Inverse associations (e.g., store both directions for bidirectional queries).
- Counters (e.g., friend count) for fast aggregation.
- Batch operations.
- Later: MyRocks (LSM-tree) for 50% space savings.

## 5. Trade-offs and Alternatives

- Why not full graph DB (e.g., Neo4j)? Too slow/heavy for online serving; TAO is read-optimized with limited queries.
- **Consistency:** Eventual, but read-after-write via cache.
- **Scalability:** Horizontal (add caches/shards); independent scaling of cache vs. storage.
- **Alternatives:** Cassandra/HBase for storage, Redis for cache, but TAO custom for graph patterns.

How to Learn More About This Topic

Start with Basics: Read "System Design Interview" by Alex Xu (Vol 1/2) – has similar questions like "Design Twitter" or "News Feed".
Deep Dive into Real Systems:
Facebook Engineering Blog: "TAO: The Power of the Graph" (official post).
Research Paper: "TAO: Facebook’s Distributed Data Store for the Social Graph" (USENIX ATC 2013) – core reference.
Follow-ups: MyRocks paper for storage improvements.

### Practice:
- Grokking the System Design Interview (Educative.io).
- **YouTube:** "System Design Interview" channels (e.g., explanations of TAO).
- **Mock interviews:** Discuss sharding, caching, consistency.

- **Related Topics:** News Feed design (fan-out vs. fan-in), recommendation systems.

This design handles Facebook-scale effectively by prioritizing reads, aggressive caching, and geo-distribution. In an interview, start simple (single machine graph), then scale out step-by-step. Good luck with your interview prep!

and so funny that they blocked (who all job in Tbilisi right?) when i wants to find,
because who wants that i go to Belarus, how and befero blocked all opportunity,
how and stollen all ideas, and others. So funny when people told me that i don't know English, really?) I know, how and codding how and a lot of things,
but everything how and now they know that i haven't money for normal food))))

So Who bad? 
Who blocked all opporunitys for me.
Who blocked everything)))
WHo who who stollen everything?
Who last 4 years destroyed my phythigs heals
And so funny that people don't told that sent to me with suicide when i was in Poland,
when they hacked me and destroyed everything all my life, and Roma you should be shut up.
Yerstaday thanks for Natsha joks, it was hard, and bad joke, but yaroslav my little btother and his girl Polina Krilova made a joke( theator when Polina told that Yaroslav 
punch her and the next they made post actors drammatic theator after i made suicedi,
but they don't say it. )So funny right? 

And know they wanted sent to me to Belarus to CRAZY people,
I swimmed in sea in Shtorm in Batumi in October
How and I made idea with Restarant and Farmer buisness 
How and I made Algotrading
How and a lot of things else, they stollen everything.

I know that Kendy made all helps to me Natasha, Julia in the end.

About Dna too my, how and a lot of things else.

How and they sent to me wrong messages about English skills. 

I have drive license with 2016 the next year last...

How and a lot of things else.

How they don't wanted that i lived in Good appartment and made everything that i leave from my flat on korchmalna 61 150 in Warsaw. H