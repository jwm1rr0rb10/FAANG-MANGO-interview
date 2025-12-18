# A* (A-Star) Search Algorithm

## Description

A* (A-Star) is an informed search algorithm used to find the optimal path between two points in a graph. It combines the strengths of Dijkstra's algorithm and greedy best-first search using a heuristic, making it fast and efficient.

---

## How it works

A* uses two functions:
- **g(n)** — the cost from the start node to the current node n.
- **h(n)** — a heuristic estimate of the cost from n to the goal (often the distance).

The algorithm selects the node with the lowest sum f(n) = g(n) + h(n) and expands it until it reaches the goal or runs out of options.

---

## Example in Python

```python
import heapq

def a_star_search(graph, start, goal, h):
    open_set = []
    heapq.heappush(open_set, (0 + h(start), 0, start, [start]))
    visited = set()

    while open_set:
        f, g, current, path = heapq.heappop(open_set)
        if current == goal:
            return path

        if current in visited:
            continue
        visited.add(current)

        for neighbor, cost in graph.get(current, []):
            if neighbor not in visited:
                heapq.heappush(open_set, (g + cost + h(neighbor), g + cost, neighbor, path + [neighbor]))
    return None

# Example graph and heuristic
graph = {
    'A': [('B', 1), ('C', 3)],
    'B': [('D', 1)],
    'C': [('D', 1)],
    'D': []
}
def heuristic(node):
    h_map = {'A': 3, 'B': 2, 'C': 1, 'D': 0}
    return h_map[node]

path = a_star_search(graph, 'A', 'D', heuristic)
print(path)  # ['A', 'B', 'D']
```

---

## Code explanation

- `a_star_search` is the main function implementing A*.
- A priority queue (`heapq`) is used to select the node with the lowest f(n).
- `visited` is a set of visited nodes.
- `graph` is a dictionary of nodes and their neighbors.
- `heuristic` is a function estimating distance to the goal.

---

## Time and space complexity

- **Time:** O(E), where E is the number of edges, but can be higher with a poor heuristic.
- **Space:** O(V), where V is the number of nodes, as the open set and visited set are stored.

---

## Where to use

A* is commonly used for:
- robot and game navigation (finding paths on a map),
- route planning (GPS, logistics),
- solving puzzles.

---

## Where using is bad

- If the heuristic is weak (does not guide towards the goal), A* may become slow and degrade to breadth-first search.
- On very large graphs with limited memory resources.

---

### Useful Links & Topics

- [A* Search Algorithm Explained (GeeksforGeeks)](https://www.geeksforgeeks.org/a-search-algorithm/)
- [A* Pathfinding Visualization (Red Blob Games)](https://www.redblobgames.com/pathfinding/a-star/)
- [Wikipedia - A* Search Algorithm](https://en.wikipedia.org/wiki/A*_search_algorithm)
- [Heuristics in A* Search (Brilliant.org)](https://brilliant.org/wiki/a-star-search/)
- [A* vs Dijkstra’s Algorithm](https://towardsdatascience.com/a-star-vs-dijkstras-algorithm-c3a5ebc56a94)
- [Video: A* Search Algorithm (Computerphile)](https://www.youtube.com/watch?v=ySN5Wnu88nI)

---

### Practice Problems

- **LeetCode Task:**  
  [LeetCode 773. Sliding Puzzle](https://leetcode.com/problems/sliding-puzzle/)  
  *Difficulty: Hard*  
  This problem can be solved efficiently with A* Search, using the number of misplaced tiles or Manhattan distance as the heuristic.

- **LeetCode 1091. Shortest Path in Binary Matrix**  
  [Problem Link](https://leetcode.com/problems/shortest-path-in-binary-matrix/)  
  *Difficulty: Medium*  
  Though BFS is typical, A* can be applied for optimal pathfinding.

- **LeetCode 127. Word Ladder**  
  [Problem Link](https://leetcode.com/problems/word-ladder/)  
  *Difficulty: Hard*  
  A* can be used to optimize the search for transformation sequences.

---

