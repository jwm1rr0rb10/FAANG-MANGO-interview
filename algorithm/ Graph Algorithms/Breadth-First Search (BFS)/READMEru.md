# Breadth-First Search (BFS) Algorithm

## What is BFS?

Breadth-First Search (BFS) — это алгоритм для обхода графов и поиска кратчайших путей в невзвешенных графах. Он исследует все соседние вершины на каждом уровне, прежде чем переходить к следующему уровню.

---

## How Does BFS Work?

1. Начинает с одной вершины (начальной).
2. Помещает её в очередь.
3. Пока очередь не пуста:
    - Извлекает вершину из очереди.
    - Посещает всех соседей этой вершины, которые ещё не были посещены, и добавляет их в очередь.

---

## Python Example

```python
from collections import deque

def bfs(graph, start):
    visited = set()
    queue = deque([start])
    result = []

    while queue:
        vertex = queue.popleft()
        if vertex not in visited:
            visited.add(vertex)
            result.append(vertex)
            queue.extend([n for n in graph[vertex] if n not in visited])
    return result

# Example graph representation
graph = {
    'A': ['B', 'C'],
    'B': ['A', 'D', 'E'],
    'C': ['A', 'F'],
    'D': ['B'],
    'E': ['B', 'F'],
    'F': ['C', 'E']
}
print(bfs(graph, 'A'))  # Output: ['A', 'B', 'C', 'D', 'E', 'F']
```

---

## BFS Properties

- **Time Complexity:** O(V + E)
- **Space Complexity:** O(V)
- **Works for:** Graphs, trees, undirected and directed graphs
- **Finds:** Shortest path in unweighted graphs

---

## Practical Applications

- Поиск кратчайших путей (например, маршруты в навигаторах)
- Социальные сети — поиск друзей через связи
- Веб-краулинг — обход ссылок на сайтах
- Решение головоломок, например, поиск минимального количества шагов

---

## Useful Links

- [BFS Explanation (GeeksforGeeks)](https://www.geeksforgeeks.org/breadth-first-search-or-bfs-for-a-graph/)
- [Wikipedia — Breadth-First Search](https://en.wikipedia.org/wiki/Breadth-first_search)
- [BFS Visualization (Brilliant.org)](https://brilliant.org/wiki/breadth-first-search-bfs/)

---

## Practice Problems

- [LeetCode 102. Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/)
- [LeetCode 542. 01 Matrix](https://leetcode.com/problems/01-matrix/)