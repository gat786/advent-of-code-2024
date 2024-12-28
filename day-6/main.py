from typing import List
from dataclasses import dataclass
from enum import Enum

up_arrow    = "^"
down_arrow  = "v"
right_arrow = ">"
left_arrow  = "<"

arrows_set = { up_arrow: True, down_arrow: True, right_arrow: True, left_arrow: True}
output_file_path = "../output/updated_matrix-day-6-1.txt"

class Orientation(Enum):
  Left    = 1
  Right   = 2
  Top     = 3
  Bottom  = 4

@dataclass
class CharachterPoint:
  line_index: int
  char_index: int

def read_file():
  matrix: List[List[str]] = []
  with open("../input/data-6","r") as fp:
    while True:
      line = fp.readline()
      line = line.strip()
      if line == "":
        break
      line_list = []
      for char_index, char in enumerate(line):
        line_list.append(char)
      matrix.append(line_list)
      
  return matrix

def is_inbounds(matrix: List[List[str]], point: CharachterPoint):
  if point.line_index < len(matrix) and point.char_index < len(matrix[0]):
    return True
  return False

def read_example():
  example = """....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#..."""
  print(example)
  matrix: List[List[str]] = []
  for line in example.splitlines():
    line = line.strip()
    if line == "":
      break
    line_list = []
    for char_index, char in enumerate(line):
      line_list.append(char)
    matrix.append(line_list)
  return matrix

def matrix_to_file(matrix: List[List[str]]):
  content = ""
  for line in matrix:
    for charachter in line:
      content += charachter
    content += "\n"
  write_to_file(content=content)

def count_x():
  x_count = 0
  with open(output_file_path, "r") as fp:
    lines = fp.readlines()
    for line in lines:
      for char in line:
        if char == "X":
          x_count += 1
  
  return x_count

def find_current_position(matrix: List[List[str]]) -> CharachterPoint | None:
  for line_index, line in enumerate(matrix):
    for item_index, item in enumerate(line):
      if item in arrows_set:
        print("Item found")
        point = CharachterPoint(line_index=line_index, char_index=item_index)
        return point

def get_next_step(curr_position: CharachterPoint, soldier_facing: str) -> CharachterPoint | None:
  print(f"Soldier is facing: {soldier_facing}, getting next step")
  match soldier_facing:
    case ">":
      new_point = CharachterPoint(
        line_index=curr_position.line_index,
        char_index=curr_position.char_index + 1
      )
      return new_point
    case "<":
      new_point = CharachterPoint(
        line_index=curr_position.line_index,
        char_index=curr_position.char_index - 1
      )
      return new_point
    case "^":
      new_point = CharachterPoint(
        line_index=curr_position.line_index - 1,
        char_index=curr_position.char_index
      )
      return new_point
    case "v":
      new_point = CharachterPoint(
        line_index=curr_position.line_index + 1,
        char_index=curr_position.char_index
      )
      return new_point
    case _:
      raise Exception("Wrong direction given")

def get_changed_direction(soldier_facing: str):
  print(f"Soldier is facing: {soldier_facing}, and faces a block on next step, changing direction")
  match soldier_facing:
    case ">":
      return "v"
    case "<":
      return "^"
    case "^":
      return ">"
    case "v":
      return "<"
    case _:
      return soldier_facing

def is_blockage(matrix: List[List[str]], curr_position: List):
  if matrix[curr_position[0]][curr_position[1]] == "#":
    return True

def get_next_move(matrix: List[List[str]], curr_position: List):
  pass

def write_to_file(content: str):
  with open(output_file_path, "w") as fp:
    fp.write(content)

def main():
  matrix = read_file()
  # matrix = read_example()
  arrow_index = find_current_position(matrix=matrix)
  if arrow_index is not None:
    print(f"Initial soldier position found at {arrow_index}")
    soldier_facing = matrix[arrow_index.line_index][arrow_index.char_index]
    while True:
      next_step = get_next_step(arrow_index, soldier_facing=soldier_facing)
      print(f"Next step will be {next_step}")
      if next_step is not None:
        if not is_inbounds(matrix=matrix, point=next_step):
          print("The next step will be outside bounds")
          matrix[arrow_index.line_index][arrow_index.char_index] = "X"
          break
        if matrix[next_step.line_index][next_step.char_index] == "#":
          soldier_facing = get_changed_direction(soldier_facing=soldier_facing)
          matrix[arrow_index.line_index][arrow_index.char_index] = soldier_facing
        else:
          matrix[arrow_index.line_index][arrow_index.char_index] = "X"
          matrix[next_step.line_index][next_step.char_index] = soldier_facing
      matrix_to_file(matrix=matrix)
      updated_arrow_index = find_current_position(matrix=matrix)
      if updated_arrow_index is not None:
        arrow_index = updated_arrow_index

    print(f"Marking {arrow_index} as X and closing evaluation")
    matrix_to_file(matrix=matrix)
    x_count = count_x()
    print(f"Count of distinct steps taken {x_count}")



if __name__ == "__main__":
  main()