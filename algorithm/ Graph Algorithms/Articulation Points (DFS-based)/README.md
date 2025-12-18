# Articulation Points (DFS-based)

## Description

Articulation points (or cut vertices) in a graph are nodes whose removal increases the number of connected components, effectively disconnecting parts of the graph. Identifying these points is crucial for understanding vulnerabilities in networks, such as critical routers, servers, or system components whose failure would split the network.

---

## How it works

The classic algorithm for finding articulation points uses Depth-First Search (DFS), tracking discovery times and the earliest reachable ancestor for each vertex:

1. Perform DFS traversal, keeping track of:
   - `disc[v]`: discovery time of vertex `v`.
   - `low[v]`: earliest discovery time reachable from `v` or its descendants.
2. For each node:
   - If it is the root and has more than one child, it is an articulation point.
   - If it is not the root and for any child `u`, `low[u] >= disc[v]`, then `v` is an articulation point.

---

## Example in Python

```python
def find_articulation_points(graph):
    def dfs(u, parent):
        nonlocal time
        visited[u] = True
        disc[u] = low[u] = time
        time += 1
        children = 0

        for v in graph[u]:
            if not visited[v]:
                children += 1
                dfs(v, u)
                low[u] = min(low[u], low[v])
                if parent is None and children > 1:
                    articulation_points.add(u)
                if parent is not None and low[v] >= disc[u]:
                    articulation_points.add(u)
            elif v != parent:
                low[u] = min(low[u], disc[v])

    n = len(graph)
    visited = {node: False for node in graph}
    disc = {}
    low = {}
    articulation_points = set()
    time = 0
    for u in graph:
        if not visited[u]:
            dfs(u, None)
    return articulation_points

# Example graph as adjacency list
graph = {
    0: [1, 2],
    1: [0, 2],
    2: [0, 1, 3, 5],
    3: [2, 4],
    4: [3],
    5: [2, 6, 7],
    6: [5, 7],
    7: [5, 6]
}
print(find_articulation_points(graph))  # Output: {2, 3, 5}
```

---

## Code explanation

- Uses DFS to traverse the graph.
- Tracks discovery and low times for each node.
- Adds nodes to `articulation_points` set if they meet articulation conditions.

---

## Time and space complexity

- **Time:** O(V + E), where V is the number of vertices and E is the number of edges.
- **Space:** O(V), for storing times and visited nodes.

---

## Where to use

Articulation points are useful for:
- Network reliability analysis (finding critical routers, servers).
- Identifying weak points in communication systems or social networks.
- Analyzing connectivity in road or utility networks.

---

## Where using is bad

- In dense graphs where most nodes are connected, articulation points may not provide useful insights.
- Not applicable for disconnected graphs without modification.
- For very large graphs, recursive DFS may hit Python’s recursion limit (consider an iterative approach).

---

### Useful Links & Topics

- [Articulation Point (GeeksforGeeks)](https://www.geeksforgeeks.org/articulation-points-or-cut-vertices-in-a-graph/)
- [Wikipedia — Articulation Point](https://en.wikipedia.org/wiki/Biconnected_component)
- [Finding Articulation Points (Brilliant.org)](https://brilliant.org/wiki/articulation-points/)
- [Video: Tarjan’s Algorithm for Articulation Points](https://www.youtube.com/watch?v=2kREIkF9UAs)

---

### Practice Problems

- **LeetCode Task:**  
  [LeetCode 1192. Critical Connections in a Network](https://leetcode.com/problems/critical-connections-in-a-network/)  
  *Difficulty: Hard*  
  This problem requires finding critical edges (bridges), but the articulation point algorithm is closely related and useful for understanding the solution.

- **Codeforces Problem:**  
  [Codeforces 796B. Articulation Points](https://codeforces.com/problemset/problem/796/B)  
  *Difficulty: Medium*  
  Practice finding articulation points in competitive programming settings.

---