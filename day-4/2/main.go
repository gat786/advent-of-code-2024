package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	xmas_map = map[string]bool{"MAS": true, "SAM": true}
)

type Point struct {
	x int
	y int
}

func get_diagonals(matrix *[][]rune, x int, y int) [][]Point {
	diagonals := make([][]Point, 2)
	diagonal_1 := make([]Point, 3)
	diagonal_2 := make([]Point, 3)
	diagonal_1_additive := []Point{{-1, -1}, {0, 0}, {1, 1}}
	diagonal_2_additive := []Point{{-1, 1}, {0, 0}, {1, -1}}
	for i, _ := range diagonal_1 {
		diagonal_1[i] = Point{x + diagonal_1_additive[i].x, y + diagonal_1_additive[i].y}
		diagonal_2[i] = Point{x + diagonal_2_additive[i].x, y + diagonal_2_additive[i].y}
	}
	diagonals[0] = diagonal_1
	diagonals[1] = diagonal_2
	return diagonals
}

func get_chars(matrix *[][]rune, combination []Point) string {
	combination_string := ""
	for _, point := range combination {
		if point.x < 0 || point.y < 0 || point.x >= len(*matrix) || point.y >= len((*matrix)[0]) {
			combination_string = ""
			break
		}
		combination_string += string((*matrix)[point.x][point.y])
	}
	return combination_string
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
			diagonal_1_string := get_chars(&matrix, diagonals[0])
			diagonal_2_string := get_chars(&matrix, diagonals[1])
			if xmas_map[diagonal_1_string] && xmas_map[diagonal_2_string] {
				fmt.Println("Both strings are XMAS, ", diagonal_1_string, diagonal_2_string)
				xmas_count++
			}
		}
		lineNum++
	}

	fmt.Println(xmas_count)
}
