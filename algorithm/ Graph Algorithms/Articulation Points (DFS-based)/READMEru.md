# Точки сочленения (DFS-based)

## Описание

Точки сочленения (или "cut vertices") — это вершины графа, удаление которых увеличивает количество компонент связности, то есть разрывает граф на отдельные части. Выявление таких точек важно для анализа уязвимости сетей: например, критические серверы или маршрутизаторы, выход из строя которых разделяет сеть.

---

## Как работает

Классический алгоритм поиска точек сочленения основан на обходе графа в глубину (DFS), с отслеживанием времени открытия и самой ранней достижимой вершины для каждой вершины:

1. Выполняем обход DFS, при этом для каждой вершины храним:
   - `disc[v]` — время открытия вершины `v`.
   - `low[v]` — минимальное время открытия среди всех достижимых вершин из `v` и её потомков.
2. Для каждой вершины:
   - Если она является корнем и имеет более одного ребёнка, то это точка сочленения.
   - Если не корень, и для любого ребёнка `u` выполняется `low[u] >= disc[v]`, то `v` — точка сочленения.

---

## Пример на Python

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

# Пример графа (список смежности)
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
print(find_articulation_points(graph))  # Вывод: {2, 3, 5}
```

---

## Описание кода

- Используется обход графа в глубину (DFS).
- Хранятся времена открытия и минимальные достижимые времена для каждой вершины.
- Узлы добавляются в множество `articulation_points`, если выполняется условие точки сочленения.

---

## Временная и пространственная сложность

- **Время:** O(V + E), где V — количество вершин, E — количество рёбер.
- **Память:** O(V), для хранения времён и посещённых вершин.

---

## Применение

Точки сочленения полезны для:
- Анализа надёжности сетей (поиск критических маршрутизаторов, серверов).
- Выявления слабых мест в коммуникационных или социальных сетях.
- Анализа связности дорожных или коммунальных сетей.

---

## Где использовать — плохо

- В плотных графах, где большинство вершин связаны, точки сочленения могут быть малоинформативны.
- Неприменимо к несвязным графам без модификации алгоритма.
- Для очень больших графов рекурсивный DFS может привести к переполнению стека (лучше использовать итеративный подход).

---

### Полезные ссылки и темы

- [Articulation Point (GeeksforGeeks, RU)](https://translated.turbopages.org/proxy_u/en-ru.ru.2a2a0a7a-6511b8f3-6c06cb0b-74722d776562/https/www.geeksforgeeks.org/articulation-points-or-cut-vertices-in-a-graph/)
- [Articulation Point — Википедия](https://ru.wikipedia.org/wiki/Двусвязный_граф)
- [Finding Articulation Points (Brilliant.org, EN)](https://brilliant.org/wiki/articulation-points/)
- [Видео: Алгоритм Тарджана для точек сочленения (EN)](https://www.youtube.com/watch?v=2kREIkF9UAs)

---

### Задачи для практики

- **LeetCode задача:**  
  [LeetCode 1192. Critical Connections in a Network](https://leetcode.com/problems/critical-connections-in-a-network/)  
  *Сложность: сложно*  
  Здесь нужно искать критические рёбра (мосты), но алгоритм поиска точек сочленения тесно связан и полезен для понимания задачи.

- **Codeforces задача:**  
  [Codeforces 796B. Articulation Points](https://codeforces.com/problemset/problem/796/B)  
  *Сложность: средняя*  
  Практика поиска точек сочленения в задачах соревнований.

---