def find_bridges(graph):
    n = len(graph)
    time = [0]  # Mutable counter
    disc = [-1] * n
    low = [-1] * n
    bridges = []

    def dfs(u, parent):
        disc[u] = low[u] = time[0]
        time[0] += 1
        for v in graph[u]:
            if disc[v] == -1:
                dfs(v, u)
                low[u] = min(low[u], low[v])
                if low[v] > disc[u]:
                    bridges.append((u, v))
            elif v != parent:
                low[u] = min(low[u], disc[v])

    for i in range(n):
        if disc[i] == -1:
            dfs(i, -1)
    return bridges

# Example usage:
# Graph as adjacency list
graph = [
    [1, 2],    # 0
    [0, 2],    # 1
    [0, 1, 3], # 2
    [2]        # 3
]
print(find_bridges(graph))  # Output: [(2, 3)]