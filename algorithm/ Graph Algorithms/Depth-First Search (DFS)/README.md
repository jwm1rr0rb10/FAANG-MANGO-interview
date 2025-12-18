# Depth-First Search (DFS) Algorithm

## What is DFS?

**Depth-First Search (DFS)** is a fundamental graph traversal algorithm.  
It explores as far as possible along each branch before backtracking, making it ideal for tasks like pathfinding, connectivity checking, and cycle detection.

---

## How Does DFS Work?

1. Start from a root (or any) node.
2. Mark the node as visited.
3. Recursively (or using a stack) visit each unvisited neighbor.
4. Continue until all nodes reachable from the start are visited.

---

## Python Example

### Recursive

```python
def dfs(graph, start, visited=None):
    if visited is None:
        visited = set()
    visited.add(start)
    for neighbor in graph[start]:
        if neighbor not in visited:
            dfs(graph, neighbor, visited)
    return visited

# Example graph
graph = {
    'A': ['B', 'C'],
    'B': ['A', 'D', 'E'],
    'C': ['A', 'F'],
    'D': ['B'],
    'E': ['B', 'F'],
    'F': ['C', 'E']
}
print(dfs(graph, 'A'))  # Output: {'A', 'B', 'E', 'F', 'C', 'D'}
```

### Iterative

```python
def dfs_iterative(graph, start):
    visited = set()
    stack = [start]
    while stack:
        vertex = stack.pop()
        if vertex not in visited:
            visited.add(vertex)
            stack.extend([n for n in graph[vertex] if n not in visited])
    return visited

print(dfs_iterative(graph, 'A'))  # Output: {'A', 'B', 'E', 'F', 'C', 'D'}
```

---

## DFS Properties

- **Time Complexity:** O(V + E)
- **Space Complexity:** O(V)
- Works for both directed and undirected graphs.
- Can be implemented recursively or iteratively.

---

## Practical Applications

- Pathfinding algorithms
- Topological sorting
- Cycle detection
- Connected components detection
- Maze and puzzle solving

---

## Useful Links

- [DFS Explanation (GeeksforGeeks)](https://www.geeksforgeeks.org/depth-first-search-or-dfs-for-a-graph/)
- [Wikipedia — Depth-First Search](https://en.wikipedia.org/wiki/Depth-first_search)
- [DFS Visualization (Brilliant.org)](https://brilliant.org/wiki/depth-first-search-dfs/)

---

## Practice Problems

- [LeetCode 200. Number of Islands](https://leetcode.com/problems/number-of-islands/)
- [LeetCode 547. Number of Provinces](https://leetcode.com/problems/number-of-provinces/)