from typing import List
from collections import deque

def get_rules_and_updates():
  rules   = {}
  updates = []
  with open("../../input/data-5","r") as input_data:
    lines = input_data.readlines()
    reading_conditions = True
    for l in lines:
      l = l.strip()
      if l == "":
        reading_conditions = False
        continue
      if reading_conditions:
        numbers = [int(x) for x in l.split("|")]
        for number in numbers:
          if number not in rules:
            rules[number] = set()
        
        rules[numbers[0]].add(numbers[1])
      else:
        numbers = [int(x) for x in l.split(",")]
        updates.append(numbers)
  return rules, updates

def check_invalid(rules: dict[int,set], update: List[int]):
  for i in range(len(update)):
    previous_items = update[0:i]
    items_that_should_not_be_behind = rules[update[i]]
    if len(items_that_should_not_be_behind.intersection(previous_items)) > 0:
      return True
  return False

def build_graph(rules: dict[int,set], update: List[int]):
  graph = { x: { "visited": False, "indegree": 0, "children": set()} for x in update }

  for item in update:
    updates_set = set(update)
    updates_set.remove(item)
    rule_for_element = rules[item]

    children_in_current_update = rule_for_element.intersection(updates_set)
    for child in children_in_current_update:
      graph[item]["children"].add(child)
      graph[child]["indegree"] += 1

  return graph

def get_zero_indegree_item(graph):
  for index,item in graph.items():
    if item["indegree"] == 0:
      return index

if __name__ == "__main__":
  rules, updates = get_rules_and_updates()

  valid_count = 0
  sum_middle  = 0

  sum_middle_of_sorted_updates = 0
  
  for update in updates:
    is_invalid = check_invalid(rules=rules,update=update)
    if is_invalid:
      print(f"Update: {update}, items_count: {len(update)}, is_invalid: {is_invalid}")
    if not is_invalid:
      valid_count += 1
      middle_item = update[len(update)//2]
      sum_middle += middle_item
    else:
      graph = build_graph(rules=rules, update=update)
      # print(graph)
      solved_queue = deque()
      
      while len(solved_queue) != len(update):
        while len(graph.items()) > 0:
          zero_indegree_item_index = get_zero_indegree_item(graph=graph)
          if zero_indegree_item_index:
            zero_indegree_item = graph[zero_indegree_item_index]
            solved_queue.append(zero_indegree_item_index)
            children_in_current_update = zero_indegree_item["children"]
            for child in children_in_current_update:
              graph[child]["indegree"] -= 1
            graph.pop(zero_indegree_item_index)
      
      solved_update = list(solved_queue)
      print(f"Sorted Queue: {solved_queue}, is_invalid: {check_invalid(rules=rules,update=solved_update)}")
      middle_item = solved_update[len(solved_update)//2]
      sum_middle_of_sorted_updates += middle_item


  
  print(f"Found {valid_count}, valid entries, having total sum of {sum_middle}")
  print(f"Sum of middle elements of invalid items: {sum_middle_of_sorted_updates}")