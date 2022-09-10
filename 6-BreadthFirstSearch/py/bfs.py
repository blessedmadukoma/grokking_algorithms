# Bread First Search in Python: using Dequeue for Queue

from collections import deque


graph = {}
graph["me"] = ["John", "Alice", "Martha"]
graph["John"] = [] # John does not have any neighbor
graph["Alice"] = ["Emmanuel", "Don", "Luke"]
graph["Emmanuel"] = [] # Emmanuel does not have any neighbor
graph["Don"] = [] # Don does not have any neighbor
graph["Luke"] = [] # Luke does not have any neighbor
graph["Martha"] = ["Bernice", "James"]
graph["Bernice"] = [] # Bernice does not have any neighbor
graph["James"] = [] # James does not have any neighbor
graph["Don"] = ["Francis", "Luke", "Daniel"]
graph["Francis"] = [] # Francis does not have any neighbor
def search(name):
 search_queue = deque()
 search_queue += graph[name]
 searched = []
 while search_queue:
  person = search_queue.popleft()
  if not person in searched:
   if person_to_find(person):
    print(person + " is found!")
    return True
   else:
    # check if the person has neighbours
    # if graph[person]:
     # search_queue += graph[person]
    search_queue += graph[person]
    # else:
     # print(f"{person} doesn't have neighbours")
    searched.append(person)

  # continue if the person isn't found
  continue
  # return False

def person_to_find(name):
 if name == "Daniel":
  return name

search("me")