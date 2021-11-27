import sys

def depth_limited_search(graph, startNode, targetNode, maxLevel):
    visitedNodes, routeCollector = [], []
    found = depth_limited_search_internal(graph, targetNode, maxLevel, visitedNodes, routeCollector,startNode, 1)
    return (found, routeCollector)
  
def depth_limited_search_internal(graph, targetNode, maxLevel, visitedNodes, routeCollector, currentNode, currentLevel):
    routeCollector.append(currentNode)
    visitedNodes.append(currentNode)

    if currentNode == targetNode:
        return True

    if currentLevel < maxLevel:
        currentVisitedNodes = visitedNodes.copy()
        for neighbour in graph[currentNode]:
            if neighbour not in currentVisitedNodes:
                found = depth_limited_search_internal(graph, targetNode, maxLevel, currentVisitedNodes, routeCollector, neighbour, currentLevel + 1)
                if found :
                    return True
    
    routeCollector.pop()
    return False

def print_hint(graph):
    hint = f"""
Command: depth_limited_search.py [start node] [target node] [max level]
- start node    : this is starting node, see nodes list
- target node   : this is the target node, see nodes list
- max limit     : max node in nodes level, angka >= 0
                    with 0 means no limit
                    with 1 means only start node

Nodes: {', '.join(graph.keys())}
    """
    print(hint)

def main(argv):
    graph = {
        'Oradea': ["Zerind", "Sibiu"],
        'Zerind':  ["Oradea", "Arad"],
        'Arad':  ["Zerind", "Sibiu", "Timisoara"],
        'Timisoara':  ["Arad", "Lugoj"],
        'Lugoj':  ["Timisoara", "Mehadia"],
        'Mehadia':  ["Lugoj", "Drobeta"],
        'Drobeta':  ["Mehadia", "Craiova"],
        'Craiova':  ["Drobeta", "Rimnicu Vilcea", "Pitesti"],
        'Rimnicu Vilcea':  ["Craiova", "Sibiu", "Pitesti"],
        'Sibiu':  ["Rimnicu Vilcea", "Oradea", "Fagaras", ],
        'Fagaras':  ["Sibiu", "Bucharest"],
        'Pitesti':  ["Rimnicu Vilcea", "Bucharest"],
        'Bucharest':  ["Giurgiu", "Urziceni"],
        'Giurgiu':  ["Bucharest"],
        'Urziceni':  ["Bucharest", "Hirsova", "Vaslui"],
        'Hirsova':  ["Urziceni", "Eforie"],
        'Eforie':  ["Hirsova"],
        'Vaslui':  ["Urziceni", "Iasi"],
        'Iasi':  ["Vaslui", "Neamt"],
        'Neamt':  ["Vaslui"],
    }

    if len(argv) != 3 :
        print_hint(graph)
        sys.exit(1)

    startNode = argv[0]
    targetNode = argv[1]
    maxLevel = 0
    try:
        maxLevel = int(argv[2])
    except ValueError as e:
        print("err: Max level should be number >= 0")
        sys.exit(1)

    if startNode not in graph :
        print(f"err: Start node '{startNode}' is invalid")
        print(f"Nodes : {', '.join(graph.keys())}")
        sys.exit(1)
    elif targetNode not in graph :
        print(f"err: Target node '{targetNode}' is invalid")
        print(f"Nodes: {', '.join(graph.keys())}")
        sys.exit(1)
    elif maxLevel < 0:
        print(f"err: Max level {maxLevel} must be >= 0")
        print_hint(graph)
        sys.exit(1)

    if maxLevel == 0:
        maxLevel = float('inf')
    
    result = depth_limited_search(graph, startNode, targetNode, maxLevel)
    found = result[0]
    route = result[1]
    if found:
        print(f"Route from '{startNode}' to '{targetNode}' with max level {maxLevel}:")
        print(' -> '.join(route))
    else:
        print(f"inf: Route from '{startNode}' ke '{targetNode}' with MaksLevel {maxLevel} is not found")

if __name__ == '__main__':
    main(sys.argv[1:])