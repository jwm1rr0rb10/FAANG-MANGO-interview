# Bridge Finding in Graphs (DFS-Based Algorithm)

## What is a Bridge?

A **bridge** (or cut-edge) in a graph is an edge whose removal increases the number of connected components. In other words, it’s an edge that, if cut, would disconnect part of the graph.

---

## Why Bridge Finding is Important

- Network reliability analysis.
- Detecting critical connections in distributed systems.
- Identifying weak points in social or transport networks.

---

## Bridge Finding Algorithm (DFS-based)

### Steps

1. Perform DFS traversal of the graph.
2. For every node, record:
    - **Discovery time**: when the node is first visited.
    - **Lowest discovery time reachable**: including back edges.
3. An edge (u, v) is a bridge if `low[v] > disc[u]` after DFS on v.

---

## Python Example

```python
def find_bridges(graph):
    n = len(graph)
    time = [0]  # Mutable counter
    disc = [-1] * n
    low = [-1] * n
    bridges = []

    def dfs(u, parent):
        disc[u] = low[u] = time[0]
        time[0] += 1
        for v in graph[u]:
            if disc[v] == -1:
                dfs(v, u)
                low[u] = min(low[u], low[v])
                if low[v] > disc[u]:
                    bridges.append((u, v))
            elif v != parent:
                low[u] = min(low[u], disc[v])

    for i in range(n):
        if disc[i] == -1:
            dfs(i, -1)
    return bridges

# Example usage:
# Graph as adjacency list
graph = [
    [1, 2],    # 0
    [0, 2],    # 1
    [0, 1, 3], # 2
    [2]        # 3
]
print(find_bridges(graph))  # Output: [(2, 3)]
```

---

## Properties

- **Time Complexity:** O(V + E)
- **Space Complexity:** O(V)
- Works for undirected graphs.
- Finds all bridges in a single DFS traversal.

---

## Practical Applications

- Analyzing computer networks for single points of failure.
- Planning redundancy in infrastructure.
- Social network analysis (finding critical relationships).

---

## References

- [Tarjan's Algorithm (Wikipedia)](https://en.wikipedia.org/wiki/Bridge_(graph_theory))
- [GeeksforGeeks — Bridge Finding](https://www.geeksforgeeks.org/bridge-in-a-graph/)