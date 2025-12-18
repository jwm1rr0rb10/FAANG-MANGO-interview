# Алгоритм Дейкстры (нахождение кратчайшего пути)

## Что такое алгоритм Дейкстры?

**Алгоритм Дейкстры** — классический способ поиска кратчайших путей от одной вершины ко всем остальным в взвешенном графе с неотрицательными весами рёбер.

---

## Как работает алгоритм Дейкстры?

1. Всем вершинам присваивается начальное расстояние: источнику — 0, остальным — бесконечность.
2. Источник выбирается текущей вершиной, остальные считаются непосещёнными.
3. Для текущей вершины рассматриваются все непосещённые соседи и пересчитывается их расстояние через текущую вершину. Если расстояние стало меньше, обновляем.
4. После обработки вершина помечается как посещённая (больше не рассматривается).
5. Выбираем следующую вершину с минимальным расстоянием среди непосещённых и повторяем шаги 3-5, пока все вершины не будут посещены или пока не найдём кратчайший путь.

---

## Пример на Python

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

# Пример графа:
graph = {
    'A': [('B', 1), ('C', 4)],
    'B': [('A', 1), ('C', 2), ('D', 5)],
    'C': [('A', 4), ('B', 2), ('D', 1)],
    'D': [('B', 5), ('C', 1)]
}
print(dijkstra(graph, 'A'))  # Вывод: {'A': 0, 'B': 1, 'C': 3, 'D': 4}
```

---

## Свойства

- **Временная сложность:** O((V + E) log V) с приоритетной очередью
- **Память:** O(V)
- Работает для графов с неотрицательными весами рёбер

---

## Применение

- Навигация GPS (расчёт кратчайшего маршрута)
- Маршрутизация в сетях
- Робототехника и планирование движения
- Навигация игровых AI

---

## Полезные ссылки

- [Алгоритм Дейкстры (GeeksforGeeks, RU)](https://translated.turbopages.org/proxy_u/en-ru.ru.2a2a0a7a-6511b8f3-d9c728e0-74722d776562/https/www.geeksforgeeks.org/dijkstras-shortest-path-algorithm-graph/)
- [Wikipedia — Алгоритм Дейкстры](https://ru.wikipedia.org/wiki/Алгоритм_Дейкстры)
- [Визуализация алгоритма (Brilliant.org, EN)](https://brilliant.org/wiki/dijkstras-shortest-path-algorithm/)

---

## Задачи для практики

- [LeetCode 743. Network Delay Time](https://leetcode.com/problems/network-delay-time/)
- [LeetCode 787. Cheapest Flights Within K Stops](https://leetcode.com/problems/cheapest-flights-within-k-stops/)