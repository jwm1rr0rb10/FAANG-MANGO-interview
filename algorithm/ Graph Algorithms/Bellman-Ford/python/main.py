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