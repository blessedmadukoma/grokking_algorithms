# Dijkstra's Algorithm

# Using 3 hash tables for the graph, costs and parents

graph = {}

graph["start"] = {}
# storing the costs of the neighbors
graph["start"]["a"] = 6
graph["start"]["b"] = 2

# adding the rest of the nodes and their neighbors
graph["a"] = {}
graph["a"]["finish"] = 1
graph["b"] = {}
graph["b"]["a"] = 3
graph["b"]["finish"] = 5

# the finish node does not have any neighbours
graph["finish"] = {}

# A hashtable to store the costs. If we do not know the cost, we can use infinity in Python
infinity = float("inf")
costs = {}
costs["a"] = 6
costs["b"] = 2
costs["finish"] = infinity

# A parents hashtable
parents = {}
parents["a"] = "start"
parents["b"] = "start"
parents["finish"] = None

# Array to keep track of all the processed nodes
processed = []

def find_lowest_cost_node(costs):
  lowest_cost = float("inf")
  lowest_cost_node = None
  for node in costs: # go through each node
    cost = costs[node]
    if cost < lowest_cost and node not in processed: # if it's the lowest cost and hasn't been processed yet
      lowest_cost = cost # set as the new lowest cost
      lowest_cost_node = node
  return lowest_cost_node

def search():
  # find the lowest cost node
  node = find_lowest_cost_node(costs)
  # while the node is not none
  while node is not None:
    cost = costs[node]
    neighbors = graph[node]
    # go through all the neighbors of this node
    for n in neighbors.keys():
      new_cost = cost + neighbors[n] # if it is cheaper to get to thsi neighbor by:
      if costs[n] > new_cost: # going through this node
        costs[n] = new_cost # update the cost for this node
        parents[n] = node # this node becomes the new parent for this neighbor
    processed.append(node) # mark the node as processed
    node = find_lowest_cost_node(costs) # find the next node to process, and loop
  return f"Cost is: {cost}"

print(search())