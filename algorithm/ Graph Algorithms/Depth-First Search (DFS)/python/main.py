def dfs(graph, start, visited=None):
    if visited is None:
        visited = set()
    visited.add(start)
    for neighbor in graph[start]:
        if neighbor not in visited:
            dfs(graph, neighbor, visited)
    return visited

# Example graph
graph = {
    'A': ['B', 'C'],
    'B': ['A', 'D', 'E'],
    'C': ['A', 'F'],
    'D': ['B'],
    'E': ['B', 'F'],
    'F': ['C', 'E']
}
print(dfs(graph, 'A'))  # Output: {'A', 'B', 'E', 'F', 'C', 'D'}

### Iterative

def dfs_iterative(graph, start):
    visited = set()
    stack = [start]
    while stack:
        vertex = stack.pop()
        if vertex not in visited:
            visited.add(vertex)
            stack.extend([n for n in graph[vertex] if n not in visited])
    return visited

print(dfs_iterative(graph, 'A'))  # Output: {'A', 'B', 'E', 'F', 'C', 'D'}