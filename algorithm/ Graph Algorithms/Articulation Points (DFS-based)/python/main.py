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

# Example graph as adjacency list
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
print(find_articulation_points(graph))  # Output: {2, 3, 5}