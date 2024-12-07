package main

import (
	"bufio"
	"fmt"
	"os"
)

func get_diagonals(matrix *[][]rune, x int, y int) [][][]int {
	diagonals := make([][][]int, 3)
	diagonals[0][0] = []int{x - 1, y - 1}
	diagonals[0][1] = []int{x, y}
	diagonals[0][2] = []int{x + 1, y + 1}
	diagonals[1][0] = []int{x - 1, y + 1}
	diagonals[1][1] = []int{x, y}
	diagonals[1][2] = []int{x + 1, y + 1}
	return diagonals
}

func main() {
	fmt.Println("Solving Day 4 Problem 1")
	fmt.Println("Reading files from the current directory")

	root := "/Users/ganeshtiwari/projects/personal/go-projects/advent-of-code-2024/input"
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
			diagonals := get_diagonals(&matrix, row_index, column_index)
			fmt.Println(diagonals)
		}
		lineNum++
	}

	fmt.Println(xmas_count)
}
