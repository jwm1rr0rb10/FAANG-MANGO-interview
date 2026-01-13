# Graph Algorithms — Overview and Reference

This file is a reference for the "Graph Algorithms" section of the repository. I reviewed the list of algorithms in the folder, 
verified and clarified their typical time and space complexities (for common implementations), and added short descriptions, 
example application areas, and practical notes. After this, we can start examining algorithms one-by-one: code, 
implementation variants, tests, and optimizations.

---
Contents
- Introduction
- How to read entries (format)
- Algorithm groups
    - Traversals and basic searches
    - Shortest-path algorithms
    - Connectivity and components
    - Spanning trees and cuts
    - Flows and matchings
    - Special paths (Eulerian, Hamiltonian)
    - DAG algorithms
    - Miscellaneous
- Full list of algorithms (descriptions, time/space, applications)
- Notes and next steps

How to read entries
- Short description — what the algorithm does.
- Time complexity — typical for a common implementation (variants noted when relevant).
- Space complexity — additional memory beyond the input graph.
- Practical notes — when to use, constraints, and common variants.

---

Algorithm groups and applications (brief)
- Traversals and searches (BFS, DFS, A*): building blocks used in most graph problems — routing, connectivity analysis, pathfinding, game AI.
- Shortest-path algorithms (Dijkstra, Bellman-Ford, Floyd–Warshall, A*): routing, mapping, logistics, cost optimization.
- Connectivity/components (Tarjan, Kosaraju, articulation points, bridges): structural analysis, fault tolerance, component detection.
- MST / Min-Cut (Kruskal, Prim, Stoer–Wagner): network design, clustering, minimizing connection cost.
- Flows and matchings (Dinic, Edmonds–Karp, Ford–Fulkerson, Hopcroft–Karp): resource allocation, assignment problems, capacity-constrained routing.
- Special paths (Eulerian, Hamiltonian): routing tasks and traversal problems; Hamiltonian problems are NP-hard.
- DAG algorithms (Topological Sort): task scheduling, dependency resolution, build systems.
- Miscellaneous: algorithms with specialized but important applications.

---

Full list of algorithms (algorithm → description, time, space, notes)

1) A* Search
- Description: Heuristic search for a shortest path from a start to a goal using f = g + h (g = cost so far, h = heuristic).
- Time complexity: Depends on heuristic; worst-case can degrade to Dijkstra or worse. With a binary heap: O((V + E) log V) ≈ O(E log V) when E ≥ V. With a Fibonacci heap: O(E + V log V). Practical performance depends heavily on heuristic quality (admissible/consistent).
- Space complexity: O(V) (open/closed sets, g/h/f arrays, parent pointers).
- Applications: game AI, robotics, navigation.
- Note: Better heuristics reduce explored nodes significantly.

2) Articulation Points (DFS-based)
- Description: Finds vertices whose removal increases the number of connected components.
- Time complexity: O(V + E).
- Space complexity: O(V) (auxiliary arrays, recursion stack).
- Applications: finding critical nodes in networks.

3) Bellman–Ford
- Description: Computes single-source shortest paths with possible negative edge weights; detects negative cycles.
- Time complexity: O(V * E).
- Space complexity: O(V) (distances and predecessors).
- Applications: finance, routing with penalties; use when negative weights are possible.

4) Breadth-First Search (BFS)
- Description: Level-by-level traversal; finds shortest paths in unweighted graphs.
- Time complexity: O(V + E).
- Space complexity: O(V) (queue, visited, parent arrays).
- Applications: shortest unweighted path, component detection, minimum number of steps.

5) Bridge Finding (DFS-based)
- Description: Finds edges whose removal increases the number of connected components.
- Time complexity: O(V + E).
- Space complexity: O(V).
- Applications: detecting fragile links in networks.

6) Depth-First Search (DFS)
- Description: Depth-first traversal, used for cycle detection, component finding, topological sorting, etc.
- Time complexity: O(V + E).
- Space complexity: O(V) (recursion or explicit stack and visited set).
- Applications: foundational building block for many graph algorithms.

7) Dijkstra's Shortest Path
- Description: Single-source shortest paths for non-negative weights.
- Time complexity: With binary heap: O((V + E) log V) ≈ O(E log V) for sparse graphs; with Fibonacci heap: O(E + V log V); with adjacency matrix: O(V^2).
- Space complexity: O(V) (distances, predecessors, priority queue).
- Applications: routing, navigation, cost optimization.

8) Dinic's Algorithm (Max Flow)
- Description: Uses level graphs and blocking flows for faster maximum flow computation.
- Time complexity: Worst-case O(V^2 * E). In practice faster; for unit networks or special graphs there are tighter bounds.
- Space complexity: O(V + E).
- Applications: max-flow/min-cut problems, assignment via flow reductions.

9) Edmonds–Karp (BFS-based augmenting paths)
- Description: Implementation of Ford–Fulkerson using BFS to pick shortest augmenting paths (by edge count).
- Time complexity: O(V * E^2).
- Space complexity: O(V + E).
- Applications: max flow for moderate-sized graphs; predictable worst-case.

10) Eulerian Path / Circuit
- Description: Finds a path or circuit that traverses every edge exactly once (when conditions on degrees are met).
- Time complexity: O(V + E) with an efficient implementation (Hierholzer’s algorithm).
- Space complexity: O(V + E) (path stack and possibly a working copy of the edge lists).
- Applications: routing that needs to traverse each edge once (postman, routing inspections).

11) Floyd–Warshall (All-Pairs Shortest Path)
- Description: Dynamic programming algorithm computing shortest paths for all vertex pairs.
- Time complexity: O(V^3).
- Space complexity: O(V^2) (distance matrix; optionally next matrix for path reconstruction).
- Applications: small dense graphs, precomputing all-pairs distances.

12) Ford–Fulkerson (Max Flow)
- Description: Augmenting path method on the residual network.
- Time complexity: O(E * f) where f is the maximum flow value (can be large). For integer capacities it terminates; for non-integer capacities, it may not.
- Space complexity: O(V + E).
- Applications: conceptual basis for flow algorithms; practical implementations use Edmonds–Karp or Dinic.

13) Hamiltonian Path
- Description: Path visiting every vertex exactly once. NP-hard problem (decision version).
- Time complexity: Naive O(n!) in worst case; Held–Karp dynamic programming: O(n^2 * 2^n) time, O(n * 2^n) space.
- Space complexity: Naive recursion O(n); DP O(n * 2^n).
- Applications: TSP variants, ordering problems; use heuristics or exponential algorithms for small n.

14) Hopcroft–Karp (Bipartite Matching)
- Description: Maximum matching in bipartite graphs using layered BFS/DFS to find many augmenting paths per phase.
- Time complexity: O(√V * E).
- Space complexity: O(V + E).
- Applications: assignment problems, resource allocation.

15) Kosaraju's SCC
- Description: Finds strongly connected components by two DFS passes (order on reversed graph then DFS in original order).
- Time complexity: O(V + E).
- Space complexity: O(V + E).
- Applications: directed graph analysis, condensation of graphs.

16) Kruskal's MST
- Description: Minimum spanning tree by sorting edges and using union-find (disjoint set).
- Time complexity: O(E log E) = O(E log V) for typical cases (sorting dominates).
- Space complexity: O(V + E) (edge list and DSU).
- Applications: network design, building minimum-cost connections.

17) Min-Cut (Stoer–Wagner)
- Description: Global minimum cut algorithm for undirected graphs.
- Time complexity: O(V^3) in the classical analysis.
- Space complexity: O(V^2) or O(V + E) depending on representation; implementations often use adjacency structures to optimize.
- Applications: clustering, identifying weak links.

18) Prim's MST
- Description: Builds MST by growing a tree, similar to Dijkstra but using edge weights to a tree.
- Time complexity: With binary heap: O((V + E) log V) ≈ O(E log V) for sparse graphs; with adjacency matrix: O(V^2) (good for dense graphs).
- Space complexity: O(V + E).
- Applications: same as Kruskal; choose Prim for dense graphs or when you have fast decrease-key operations.

19) Tarjan's SCC
- Description: Single-pass algorithm using low-link values and a stack to extract strongly connected components.
- Time complexity: O(V + E).
- Space complexity: O(V) (stack, indices, low-link arrays).
- Applications: same as Kosaraju, often preferred due to single DFS pass.

20) Topological Sort
- Description: Produces a linear ordering of vertices in a DAG such that all edges go from earlier to later in the order (Kahn’s algorithm or DFS-based).
- Time complexity: O(V + E).
- Space complexity: O(V + E) (graph storage and auxiliary queue/stack).
- Applications: task scheduling, dependency resolution, build systems.

---

Practical notes about complexities and usage
- Time bounds often depend on the priority queue or other data structures used: Dijkstra and A* differ between binary heap and Fibonacci heap implementations; in practice binary heaps (std::priority_queue) are common, so O((V + E) log V) / O(E log V) is the practical rule.
- For max-flow algorithms, practical performance can differ greatly from worst-case asymptotics; Dinic is usually much faster than Edmonds–Karp and Ford–Fulkerson on medium/large inputs.
- For NP-hard problems (Hamiltonian), always note exponential behavior; prefer heuristics, approximations, or exponential algorithms only for small sizes.
- Space complexity here refers to additional memory beyond input graph representation. If the graph is stored as an adjacency matrix, the input already costs O(V^2); adjacency lists cost O(V + E).

---

What I did and what’s next
I translated and prepared this English README mirroring the Russian version: verified descriptions, clarified typical time/space complexities (with common-implementation variants), and added practical notes and application areas. Next, tell me which algorithm you'd like to dive into first (I can produce detailed explanations, reference implementations, test cases, and optimization suggestions). I can also prepare a commit/PR with this README in the repository if you want me to add it directly.