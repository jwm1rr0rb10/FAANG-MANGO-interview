# Real-Time Collaborative Editing System: Enhanced Design with Segment IDs and Dynamic Modes


## Overview

This `README` provides a comprehensive design for a real-time collaborative editing system inspired by `Google Docs`,
but enhanced with innovative features like segment-based IDs for granular locking, dynamic editing modes based on user count,
local caching with autosave, and adaptive conflict resolution. The focus is on building a "better" system from the ground up,
addressing common pitfalls in existing implementations (e.g., high conflict rates in large groups, scalability issues, and poor offline support).
We'll highlight problems in traditional approaches `(OT/CRDT)` and how this design mitigates them.

---

## Key goals:

- `Scalability:` Handle 2 to 1000+ concurrent users per document without degrading performance.
- `Usability:` Prevent "editing chaos" by introducing locks and modes, while maintaining real-time feel.
- `Reliability:` Use local caches for offline resilience and autosave to minimize data loss.
- `Flexibility:` Allow manual or auto mode switching for different scenarios (e.g., brainstorming vs. enterprise editing).

---

## Traditional problems addressed:

- `Conflict Overload:` In pure OT/CRDT, concurrent edits on the same text lead to unpredictable merges (e.g., lost intent in rich text).
- `Scalability Bottlenecks:` Broadcasting every change to 1000+ users overwhelms servers/networks.
- `User Frustration:` No way to "claim" sections, leading to overwrites in large teams.
- `Offline Gaps:` Delays in sync cause data divergence.

This design builds "better" by layering granular controls on top of OT/CRDT, making it adaptive and robust.

---

## System Architecture

### High-Level Components

The system follows a client-server architecture with distributed elements for scalability. Core components:

- `Client (Frontend):` Handles UI, local edits, and optimistic updates. Uses a local cache for offline work.
- `Collaboration Server:` Central hub for coordinating operations, resolving conflicts, and broadcasting changes. Sharded by document ID.
- `Lock Manager:` A dedicated service (e.g., on Redis) for managing segment locks and ownership.
- `Storage Layer:` Persistent store for document logs, snapshots, and metadata.
- `Real-Time Gateway:` WebSocket-based for low-latency communication.
- `Mode Controller:` Monitors user count and switches editing modes dynamically.

**Data Flow:**

1. User edits locally → Queues operation with segment ID and user ID.
2. Client checks/sends to server via WebSocket.
3. Server validates lock → Applies OT/CRDT → Broadcasts if approved.
4. Clients apply updates, respecting locks.

## Diagram: High-Level Architecture (Mermaid UML) 

