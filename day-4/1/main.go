package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func valid_combinations(x int, y int) [][][]int {
	combinations := make([][][]int, 8)
	additives := [][]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
	for index := range len(additives) {
		xmas := "XMAS"
		combo := make([][]int, 4)
		for i := 0; i < len(xmas); i++ {
			combo[i] = []int{x + additives[index][0]*i, y + additives[index][1]*i}
		}
		if is_valid_combo(combo) {
			combinations[index] = combo
		}
	}
	return combinations
}

func is_valid_combo(combo [][]int) bool {
	flatt_list := []int{}
	for _, point := range combo {
		// fmt.Println(p)
		flatt_list = append(flatt_list, point...)
	}
	allowed_min := 0
	allowed_max := 140
	if allowed_min <= slices.Min(flatt_list) && slices.Max(flatt_list) < allowed_max {
		return true
	}
	return false
}

func get_chars(matrix *[][]rune, combination [][]int) string {
	combination_string := ""
	for _, point := range combination {
		combination_string += string((*matrix)[point[0]][point[1]])
	}
	return combination_string
}

func main() {
	fmt.Println("Solving Day 4 Problem 1")
	fmt.Println("Reading files from the current directory")

	root := "/Users/gat786/projects/personal/golang/advent-of-code-2024/input"
	file_name := "data-4"

	fp, err := os.Open(fmt.Sprintf("%s/%s", root, file_name))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	lineNum := 0
	scanner := bufio.NewScanner(fp)
	matrix := make([][]rune, 0)
	for scanner.Scan() {
		lineContent := scanner.Text()
		lineList := make([]rune, 0)
		for char_index, _ := range lineContent {
			lineList = append(lineList, rune(lineContent[char_index]))
		}
		lineNum++
		matrix = append(matrix, lineList)
	}

	xmas_count := 0
	for row_index, row := range matrix {
		for column_index, _ := range row {
			fmt.Println(row_index, column_index)
			combinations := valid_combinations(row_index, column_index)
			for _, combo := range combinations {
				if len(combo) > 0 {
					result_chars := get_chars(&matrix, combo)
					if result_chars == "XMAS" {
						xmas_count++
					}
				}
			}
		}
		lineNum++
	}

	fmt.Println(xmas_count)
	// for _, combo := range combinations {
	// 	// valid := is_valid_combo(combo)
	// 	fmt.Println(valid)
	// }
}
