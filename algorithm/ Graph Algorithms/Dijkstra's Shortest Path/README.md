# Dijkstra's Shortest Path Algorithm

## What is Dijkstra's Algorithm?

**Dijkstra's Algorithm** is a classic algorithm for finding the shortest path between nodes in a weighted graph (with non-negative edge weights).  
It efficiently computes the minimum cost to reach every vertex from a source node.

---

## How Does Dijkstra's Algorithm Work?

1. Assign each node a tentative distance value: set it to zero for the initial node and infinity for all others.
2. Set the initial node as current. Mark all other nodes as unvisited.
3. For the current node, consider all its unvisited neighbors and calculate their tentative distances through the current node. Update if the new distance is smaller.
4. Once processed, mark the current node as visited. Visited nodes are not checked again.
5. Select the unvisited node with the smallest tentative distance as the new current node, and repeat steps 3-5 until all nodes are visited or the shortest path is found.

---

## Python Example

```python
import heapq

def dijkstra(graph, start):
    # graph: {node: [(neighbor, weight), ...]}
    distances = {node: float('inf') for node in graph}
    distances[start] = 0
    queue = [(0, start)]

    while queue:
        current_dist, current_node = heapq.heappop(queue)
        if current_dist > distances[current_node]:
            continue
        for neighbor, weight in graph[current_node]:
            distance = current_dist + weight
            if distance < distances[neighbor]:
                distances[neighbor] = distance
                heapq.heappush(queue, (distance, neighbor))
    return distances

# Example usage:
graph = {
    'A': [('B', 1), ('C', 4)],
    'B': [('A', 1), ('C', 2), ('D', 5)],
    'C': [('A', 4), ('B', 2), ('D', 1)],
    'D': [('B', 5), ('C', 1)]
}
print(dijkstra(graph, 'A'))  # Output: {'A': 0, 'B': 1, 'C': 3, 'D': 4}
```

---

## Properties

- **Time Complexity:** O((V + E) log V) with a min-heap (priority queue)
- **Space Complexity:** O(V)
- Works for weighted graphs with non-negative edge weights.

---

## Practical Applications

- GPS navigation (shortest route calculation)
- Network routing
- Robotics and path planning
- Game AI navigation

---

## Useful Links

- [Dijkstra's Algorithm (GeeksforGeeks)](https://www.geeksforgeeks.org/dijkstras-shortest-path-algorithm-graph/)
- [Wikipedia — Dijkstra's Algorithm](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm)
- [Dijkstra Visualization (Brilliant.org)](https://brilliant.org/wiki/dijkstras-shortest-path-algorithm/)

---

## Practice Problems

- [LeetCode 743. Network Delay Time](https://leetcode.com/problems/network-delay-time/)
- [LeetCode 787. Cheapest Flights Within K Stops](https://leetcode.com/problems/cheapest-flights-within-k-stops/)