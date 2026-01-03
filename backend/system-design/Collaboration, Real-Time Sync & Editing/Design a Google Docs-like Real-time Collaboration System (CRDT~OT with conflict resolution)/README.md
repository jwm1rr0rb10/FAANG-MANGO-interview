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

![Mermaid UML](https://github.com/ogamor69wm1rr0rb/senior_question_interview/blob/main/images/backend/system-design/google-doc-1.svg)

This diagram shows client connections funneling through the gateway, with the server orchestrating locks, storage, and modes. Problems in traditional setups (e.g., single-point failure in server) are mitigated by sharding and replication.

--- 

## Key Features and Improvements

### 1. Segment-Based IDs for Granular Locking

#### How it Works:

- Each editable segment (e.g., word, sentence, paragraph, table cell) gets a unique ID (UUID or Snowflake) generated on creation or split.
- Operations are tied to segment IDs instead of positions: e.g., `{op: 'insert', segmentID: 'abc123', text: 'new', userID: 'user456'}.`
- Locking: When a user starts editing a segment, the Lock Manager assigns ownership (with TTL timeout, e.g., 30s inactivity).
- If locked, other users see read-only mode with visual indicators (e.g., highlighted border).

---

#### Why Better:

- **Prevents Interference:** Unlike pure OT (where transforms can garble intent), locks ensure atomic edits on segments, reducing merge errors.
- **Scalability for Large Groups:** For 1000+ users, divide document into segments and assign owners, turning chaos into coordinated work (like Git branches but real-time).
- **Problems Addressed:** Traditional position-based edits fail with shifts (e.g., insert shifts all positions); IDs are stable.

---

#### Potential Issues:

- Overhead: Storing IDs bloats data (mitigate with compression).
- Deadlocks: If users lock too many segments (solve with per-user limits).
- Granularity: Too fine (per char) → high overhead; too coarse (per page) → blocks collaboration. Adaptive: Auto-merge small segments.

---

### 2. Local Caching with Autosave

#### How it Works:

- Each client maintains a local cache (IndexedDB) mirroring the document state.
- Edits are applied optimistically locally, queued, and autosaved every 5-10s.
- On reconnect/sync: Send queue to server, which merges using OT/CRDT + locks.
- Tied to user ID: Cache includes `{segmentID, pendingOps, userID}` for personalized views.

---

#### Why Better:

- **Offline Resilience:** Users continue editing without internet; sync resolves conflicts later (better than Google Docs' limited offline).
- **Reduces Server Load:** Not every keystroke hits the server; batch autosaves.
- **Data Safety:** Autosave prevents loss from crashes.

--- 

#### Problems Addressed:

- **Network Flakiness:** Traditional systems drop changes on disconnect; here, queue persists.
- **Latency:** Local applies feel instant.

---

#### Potential Issues:

- **Divergence:** Long offline leads to large merges (mitigate with conflict previews).
- **Storage Limits:** Browser caps (e.g., 5MB IndexedDB); compress or evict old data.

--- 

### 3. Dynamic Editing Modes

#### How it Works:

- Modes auto-switch based on active users (monitored via WebSocket heartbeats).
- Manual override via UI/API.
- Each mode adjusts rules: locking strictness, broadcast frequency, merge strategy.

**Mode Details** (Table):

| User Count | Mode Name             | Key Rules                                                                  | Benefits                                | Trade-Offs                                |
|-----------:|:----------------------|:---------------------------------------------------------------------------|:----------------------------------------|:------------------------------------------|
| 2–10       | Free Collaboration    | Full OT/CRDT, no mandatory locks. Optional manual segment claims.          | Fast, creative, low friction            | Potential minor conflicts                 |
| 10–50      | Soft Locking          | Auto-lock on edit start. Unlock on save or timeout. Queue pending ops.     | Balances collaboration with protection  | Slight delay on locked segments           |
| 50–100     | Granular Ownership    | Assign owners per segment by user ID. FIFO queue for edit requests.        | Coordinated for medium teams            | More waiting under high contention        |
| 100–500    | Turn-Based            | Server processes ops in batches (e.g., every 10s). Local preview enabled.  | Stable for large crowds                 | Higher latency, less real-time feel       |
| 500–1000+  | Strict Moderated      | Moderator approval required. Batch syncs. Full local preview and caching.  | Secure for enterprise-scale editing     | Slowest, requires hierarchy and oversight |

### Why Better:

- **Adaptive:** Traditional systems are one-size-fits-all (e.g., always eventual consistency), leading to overload in large groups. This scales rules dynamically.
- **Customization:** Users choose modes for context (e.g., "Strict" for legal docs).
- **Over-Insurance:** Layers extra safety without killing usability.

---

### Problems Addressed:

- **Over-Broadcasting:** In large modes, reduce to batches → lower network use.
- **User Overwhelm:** Visual cues and queues prevent "edit wars".

---

### Potential Issues:

- **Mode Thrashing:** Frequent switches if users join/leave (debounce with hysteresis, e.g., wait 1min).
- **Complexity:** More code paths (test thoroughly).
- **Fairness:** Queueing might favor early users (add priorities by role).

### Diagram: Editing Flow in Dynamic Modes (Mermaid Flowchart)

![Mermaid Flowchart](https://github.com/ogamor69wm1rr0rb/senior_question_interview/blob/main/images/backend/system-design/google-doc2.svg)

This flowchart illustrates how modes alter the flow, adding checks for larger groups.

---

### Conflict Resolution: Hybrid OT/CRDT with IDs

- **Base:** Use OT for immediate consistency in small modes; CRDT for eventual in large/offline.
- **Enhancement:** IDs make operations commutative where possible, with locks preventing most conflicts.
- **Merge Logic:** If conflicts (e.g., two edits on locked segment), prioritize by timestamp/user role, or use AI (e.g., LLM for semantic merge).
- **Problems Addressed:** OT's complexity in transforms; CRDT's metadata bloat (IDs are lightweight).

--- 

### Technologies and Implementation

- **Frontend:** React + ProseMirror (for rich text with custom ID extensions).
- **Backend:** Node.js/Go + Socket.io for WebSockets.
- **Locks/Pub-Sub:** Redis (Redlock for distributed locks).
- **Storage:** MongoDB for op logs + PostgreSQL for metadata.
- **Libs:** Yjs (CRDT with ID support) or ShareDB (OT).
- **Deployment:** Kubernetes for sharding; AWS/GCP for geo-replication.
- **Monitoring:** Prometheus for user counts/mode switches.

### Prototype Steps:

1. Set up basic OT/CRDT with Yjs.
2. Add segment ID generation.
3. Implement Lock Manager.
4. Build mode logic as a state machine.
5. Test with simulated users (e.g., Locust).

#### Potential Problems and Further Improvements

- **Performance:** High user counts → latency spikes (improve with edge computing/CDNs).
- **Security:** ID spoofing (use JWT for auth).
- **Accessibility:** Ensure locks don't frustrate (add voice-over cues).
- **Future:** Integrate AI for auto-segmentation or conflict suggestions. Add versioning like Git for branches.

This design creates a more robust, user-friendly system by focusing on proactive improvements over reactive fixes.

---

### References

- Inspired by OT/CRDT papers and tools like Yjs.
- For diagrams: Use tools like Mermaid.js for rendering in docs.