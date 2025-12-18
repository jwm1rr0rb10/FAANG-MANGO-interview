# Design a Global Social Graph Storage(Facebook-scale)





## 1. Goal & Score

- `Problem`: design a global, highly-available, low-latency storage system to persist and serve the social graph(users and relationships) at Facebook-scale (billions of users, hundreds of billions of edges).

- `What to support`:
    - Friendships(mutual) and follow relationships(directed)
    - Edge metadata(type, timestamp, privacy, status)
    - Fast reads: getFriends(user), areFriends(u,v), mutualFriends(u,v), follower/following lists, paginated edges.
    - Efficient writes for friend requests, accepts, unfriends, follow/unfollow
    - Graph traversals(short paths, two-hop neighbors, recommendations) -- both online
    (low-latency) and offline/batch.
    - Global scale: multi-DC deployment, geo-replication, SLA for reads(~tens of ms) and writes
    - Privacy & ACL enforcement
    - Analytics/graph processing for recommendations.

## 2. Key Requirements (Functional & Non-functional)

- `Functional:`
    - CRUD edges and nodes.
    - Query adjacency (list edges with paging + filters)
    - Query small traversals (1-2 hops) quickly.
    - Edge attributes and indexing.

- `Non-functional:`
    - Scale: billions of users, 100s billions to trillions of edges.
    - High read QPS(reads >> writes), low latency(mc-level)
    - High availability (multi-DC, tolerate partitions)
    - Cost-effective storage
    - Eventual consistency acceptable for many read paths; some opearations require stronger guarantees(friend accept flow)
    - Secure and privacy-compliant

 ## 3. Workload Characteristics and Petterns.

 - Read-heavy: fetching friend lists, profile visits, feed generation, social graphs for UI components.
 - Writes: friend/follow actions are frequent but far fewer than reads: some users(celebruties) have extreme fan-out writes and reads.
 - Access skew: power-law: few users(celebrities) have millions of edges; most users have small-to-medium degree.
 - Query types;
    - Single-user adjacency(get friends/followers)
    - edge existence (areFriends)
    - Set operations (mutual friends -> intersection)
    - Two-hop traversals for recommendtaions
    - Long-path queries for graph analytics (offline).

4.  Data Model Options
- Adjacency list(pre user): store lsit of neighbor IDs for each user (typical).
    - Key: userID, Column: neighbor-list (sorted by ID or timestamp), Edge metadata stored inline or in separate structure.
- Edge-centric (edge table): each edge as a row: (src, dst, type, status, created_at, attrs). Good for queries filtered by time/attrs.
- Hybrid: adjacency lists for fast read of neighbors + edge table for searching/metadata.
- Representations for adjacency lists:
    - Ordered arrays of 64-bit IDs (sorted) - good for intersections.
    - Compressed delta-encoded 