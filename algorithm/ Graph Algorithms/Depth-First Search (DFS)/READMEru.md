# Алгоритм обхода в глубину (DFS)

## Что такое DFS?

**Обход в глубину (Depth-First Search, DFS)** — это базовый алгоритм для обхода графа или дерева.  
Он углубляется по каждому пути до конца, а затем возвращается назад (backtracking).

---

## Как работает DFS?

1. Начинает с корневой (или любой) вершины.
2. Помечает её как посещённую.
3. Рекурсивно (или с помощью стека) посещает всех непосещённых соседей.
4. Продолжает, пока не посетит все доступные вершины.

---

## Примеры на Python

### Рекурсивный

```python
def dfs(graph, start, visited=None):
    if visited is None:
        visited = set()
    visited.add(start)
    for neighbor in graph[start]:
        if neighbor not in visited:
            dfs(graph, neighbor, visited)
    return visited

# Пример графа
graph = {
    'A': ['B', 'C'],
    'B': ['A', 'D', 'E'],
    'C': ['A', 'F'],
    'D': ['B'],
    'E': ['B', 'F'],
    'F': ['C', 'E']
}
print(dfs(graph, 'A'))  # Вывод: {'A', 'B', 'E', 'F', 'C', 'D'}
```

### Итеративный

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

print(dfs_iterative(graph, 'A'))  # Вывод: {'A', 'B', 'E', 'F', 'C', 'D'}
```

---

## Свойства DFS

- **Временная сложность:** O(V + E)
- **Память:** O(V)
- Используется для ориентированных и неориентированных графов.
- Можно реализовать рекурсивно или итеративно.

---

## Применение

- Поиск путей
- Топологическая сортировка
- Поиск циклов
- Поиск компонент связности
- Решение головоломок и лабиринтов

---

## Полезные ссылки

- [Объяснение DFS (GeeksforGeeks, RU)](https://translated.turbopages.org/proxy_u/en-ru.ru.2a2a0a7a-6511b8f3-d9c728e0-74722d776562/https/www.geeksforgeeks.org/depth-first-search-or-dfs-for-a-graph/)
- [Wikipedia — DFS](https://ru.wikipedia.org/wiki/Обход_в_глубину)
- [DFS Visualization (Brilliant.org, EN)](https://brilliant.org/wiki/depth-first-search-dfs/)

---

## Задачи для практики

- [LeetCode 200. Number of Islands](https://leetcode.com/problems/number-of-islands/)
- [LeetCode 547. Number of Provinces](https://leetcode.com/problems/number-of-provinces/)