# Алгоритм A* (A-Star) поиска

## Описание

A* (A-Star) — это информированный алгоритм поиска, который используется для нахождения оптимального пути между двумя точками в графе. Он сочетает в себе преимущества алгоритма Дейкстры и жадного поиска по эвристике, что обеспечивает быструю и эффективную работу.

---

## Как работает

A* использует две функции:
- **g(n)** — стоимость пути от начальной точки до текущей вершины n.
- **h(n)** — эвристическая оценка стоимости от n до целевой вершины (обычно — расстояние).

Алгоритм выбирает вершину с минимальной суммой f(n) = g(n) + h(n) и расширяет её, пока не достигнет цели или не исчерпает варианты.

---

## Пример на Python

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

# Пример графа и эвристики
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

## Описание кода

- `a_star_search` — основная функция, реализующая A*.
- Используется очередь с приоритетом (`heapq`), чтобы выбирать вершину с минимальным f(n).
- `visited` — множество посещённых вершин.
- `graph` — словарь с вершинами и их соседями.
- `heuristic` — функция оценки расстояния до цели.

---

## Временная и пространственная сложность

- **Время:** O(E), где E — количество рёбер, но может быть больше при слабой эвристике.
- **Память:** O(V), где V — количество вершин, так как хранится очередь открытых вершин и множество посещённых.

---

## Применение

A* часто используется в:
- навигации роботов и игр (поиск пути на карте),
- планировании маршрутов (GPS, логистика),
- решении головоломок.

---

## Где использовать — плохо

- Если эвристика неадекватна (не приближает к цели), алгоритм работает медленно и может деградировать до поиска в ширину.
- В очень больших графах с ограниченными ресурсами памяти.

---

### Полезные ссылки и темы

- [A* Search Algorithm Explained (GeeksforGeeks, RU)](https://translated.turbopages.org/proxy_u/en-ru.ru.2a2a0a7a-6511b8f3-d9c728e0-74722d776562/https/www.geeksforgeeks.org/a-search-algorithm/)
- [Визуализация поиска пути A* (Red Blob Games, RU)](https://translated.turbopages.org/proxy_u/en-ru.ru.2a2a0a7a-6511b8f3-6c06cb0b-74722d776562/https/www.redblobgames.com/pathfinding/a-star/)
- [A* Search Algorithm — Википедия](https://ru.wikipedia.org/wiki/A*_search_algorithm)
- [Эвристики в A* Search (Brilliant.org, EN)](https://brilliant.org/wiki/a-star-search/)
- [A* vs Dijkstra’s Algorithm (EN)](https://towardsdatascience.com/a-star-vs-dijkstras-algorithm-c3a5ebc56a94)
- [Видео: A* Search Algorithm (Computerphile, EN)](https://www.youtube.com/watch?v=ySN5Wnu88nI)

---

### Задачи для практики

- **LeetCode задача:**  
  [LeetCode 773. Sliding Puzzle](https://leetcode.com/problems/sliding-puzzle/)  
  *Сложность: сложно*  
  Эту задачу можно эффективно решить с помощью A* Search, используя количество неверно размещённых плиток или манхэттенское расстояние в качестве эвристики.

- **LeetCode 1091. Shortest Path in Binary Matrix**  
  [Ссылка на задачу](https://leetcode.com/problems/shortest-path-in-binary-matrix/)  
  *Сложность: средняя*  
  Обычно используется BFS, но A* также применяется для поиска оптимального пути.

- **LeetCode 127. Word Ladder**  
  [Ссылка на задачу](https://leetcode.com/problems/word-ladder/)  
  *Сложность: сложно*  
  A* Search позволяет оптимизировать поиск последовательности преобразований.

---