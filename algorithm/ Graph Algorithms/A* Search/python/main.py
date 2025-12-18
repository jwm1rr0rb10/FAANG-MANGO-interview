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