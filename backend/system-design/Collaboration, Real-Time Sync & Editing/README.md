# System Design: Collaboration, Real-Time Sync & Editing

## Topic Overview

Real-time collaborative systems enable multiple users to edit shared content simultaneously with immediate visibility of changes. This domain requires solving fundamental challenges of consistency, latency, conflict resolution, and state synchronization across distributed clients. The core tension is between providing instant feedback (optimistic local updates) and maintaining eventual consistency across all participants. These systems power modern collaborative tools, multiplayer experiences, and live status services.

## Algorithm Descriptions

### 1. CRDT (Conflict-Free Replicated Data Type)

- **Purpose:** Achieve eventual consistency without coordination

- **Core Principle:** Mathematical structures where operations commute (order doesn't matter)

- **Types:** State-based (send full state), operation-based (send operations)

- **Use Case:** Collaborative text editing, counters, sets in distributed systems

- **Key Feature:** Automatic conflict resolution through idempotent operations

---

### 2. OT (Operational Transformation)
   
- **Purpose:** Transform operations to maintain consistency

- **Core Principle:** Apply transformations to operations so they can be applied in different orders

- **Workflow:** Client operations transformed against server history before application

- **Use Case:** Google Docs, Etherpad (historical approach)

- **Complexity:** Requires central coordination or sophisticated transformation logic

---

### 3. State Synchronization (for Gaming)

- **Purpose:** Keep game clients in consistent state

- **Models:**

    - **Lockstep:** Deterministic simulation with synchronized inputs

    - **Client-Server:** Single source of truth with state snapshots

    - **Peer-to-Peer:** Distributed authority with consensus

- **Techniques:** Entity interpolation, prediction, reconciliation

---

### 4. Lag Compensation

- **Purpose:** Mitigate network latency in real-time interactions

- **Methods:**

    - **Client-side Prediction:** Act immediately, reconcile with server

    - **Server Rewind:** Process actions based on past game state

    - **Interpolation:** Smooth display of other entities' movements

- **Use Case:** First-person shooters, fast-paced multiplayer games

---
 

### 5. Presence & Status Algorithms

- **Purpose:** Track user availability and activity in real-time

- **Techniques:**

- **Heartbeat:** Regular pings to indicate liveliness

- **Last-Seen Timestamps:** Efficient but less immediate

- **WebSocket Pub/Sub:** Real-time status propagation

- **Grace Periods:** Prevent flickering between states

---

### 6. Versioning with Merge Semantics

- **Purpose:** Maintain document history with branching/merging

- **Approaches:**

    - **DAG-based:** Graph structure for non-linear history

    - **Patch-based:** Store differences between versions

    - **Three-way Merge:** Use common ancestor to resolve conflicts

    - **Operational Merging:** Apply OT/CRDT principles to version trees

---

### System Tradeoffs

|Algorithm   | Consistency Model    | Latency Tolerance | 	Conflict Resolution   | Best For |
|:-----------|:---------------------|:------------------|:-----------------------|:---------|
| CRDT	     | Eventual	       	    | Hight             | Automatic	 |    Text, simple structures|
| OT	     | Strong (eventual)    | Medium	           | Transformation rules	|Complex documents| 
| State Sync | Eventual/Strong	     |  Very Low         |   Games, simulations|
| Presence	 | Eventual	Medium	     | Last-write-wins	  | Status indicators   |   

These algorithms represent different points in the design space of distributed consistency, each optimized for specific collaboration scenarios from document editing to real-time gaming.
