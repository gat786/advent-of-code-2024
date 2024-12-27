from typing import List

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
  graph = { x: { "visited": False, "indegree": 0, "children": []} for x in update }

  for item in update:
    updates_set = set(update)
    updates_set.remove(item)
    rule_for_element = rules[item]

    children_in_current_update = rule_for_element.intersection(updates_set)
    for child in children_in_current_update:
      graph[item]["children"].append(child)
      graph[child]["indegree"] += 1

  return graph

if __name__ == "__main__":
  rules, updates = get_rules_and_updates()

  valid_count = 0
  sum_middle  = 0
  for update in updates:
    is_invalid = check_invalid(rules=rules,update=update)
    print(f"Update: {update}, items_count: {len(update)}, is_invalid: {is_invalid}")
    if not is_invalid:
      valid_count += 1
      middle_item = update[len(update)//2]
      sum_middle += middle_item
    else:
      graph = build_graph(rules=rules, update=update)
      print(graph)
  
  print(f"Found {valid_count}, valid entries, having total sum of {sum_middle}")