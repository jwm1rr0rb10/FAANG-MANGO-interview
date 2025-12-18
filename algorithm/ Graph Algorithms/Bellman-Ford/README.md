# Bellman-Ford Algorithm

## Description

Bellman-Ford is a single-source shortest path algorithm that works on graphs with negative edge weights and can detect negative-weight cycles. Unlike Dijkstra’s algorithm, Bellman-Ford does not require all edge weights to be non-negative and is suitable for graphs where negative cycles need to be detected (e.g., in currency exchange, cost optimization).

---

## How it works

1. Initialize the distance to the source as 0 and all other nodes as infinity.
2. Relax all edges V-1 times (where V is the number of vertices), updating the shortest paths.
3. After V-1 passes, check all edges again:
   - If a shorter path is still found, a negative cycle exists.

---

## Example in Python

```python
def bellman_ford(graph, source):
    # graph: list of edges (u, v, weight)
    distance = {v: float('inf') for v in graph['vertices']}
    distance[source] = 0

    for _ in range(len(graph['vertices']) - 1):
        for u, v, w in graph['edges']:
            if distance[u] + w < distance[v]:
                distance[v] = distance[u] + w

    # Detect negative cycles
    for u, v, w in graph['edges']:
        if distance[u] + w < distance[v]:
            raise Exception("Graph contains a negative-weight cycle")

    return distance

# Example graph
graph = {
    'vertices': ['A', 'B', 'C', 'D'],
    'edges': [
        ('A', 'B', 1),
        ('B', 'C', 3),
        ('A', 'C', 10),
        ('C', 'D', 2),
        ('D', 'B', -10)
    ]
}
print(bellman_ford(graph, 'A'))
```

---

## Code explanation

- `distance` dictionary stores shortest distance to each vertex from the source.
- The relaxation step updates the shortest known paths.
- A final check detects negative-weight cycles.

---

## Time and space complexity

- **Time:** O(V * E)
- **Space:** O(V)

---

## Where to use

Bellman-Ford is useful for:
- Graphs with negative weights.
- Detecting arbitrage opportunities in currency exchange.
- Finding shortest paths in routing protocols (e.g., RIP).
- Cost optimization problems.

---

## Where using is bad

- Not efficient for dense graphs; slower than Dijkstra’s for non-negative weights.
- Not suitable when graph size is very large and performance is critical.

---

### Useful Links & Topics

- [Bellman-Ford Algorithm Explained (GeeksforGeeks)](https://www.geeksforgeeks.org/bellman-ford-algorithm-dp-23/)
- [Wikipedia — Bellman-Ford Algorithm](https://en.wikipedia.org/wiki/Bellman%E2%80%93Ford_algorithm)
- [Bellman-Ford Visualization (Brilliant.org)](https://brilliant.org/wiki/bellman-ford/)
- [Negative Weight Cycle Detection (YouTube)](https://www.youtube.com/watch?v=obWXjtg0L64)

---

### Practice Problems

- **LeetCode Task:**  
  [LeetCode 787. Cheapest Flights Within K Stops](https://leetcode.com/problems/cheapest-flights-within-k-stops/)  
  *Difficulty: Medium*  
  Bellman-Ford is an efficient solution for this problem with limited stops.

- **LeetCode 853. Car Fleet II**  
  [Problem Link](https://leetcode.com/problems/car-fleet-ii/)  
  *Difficulty: Hard*  
  Can be modeled with edge relaxations.

---