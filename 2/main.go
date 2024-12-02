package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println("Reading files from the current directory")
	root := "/Users/gat786/projects/personal/golang/advent-of-code-2024/input"

	fp, err := os.Open(fmt.Sprintf("%s/%s", root, "data-1"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	firstList := make([]int, 0)
	occurenceInSecond := make(map[int]int)

	for scanner.Scan() {
		isFirst := true
		lineContent := scanner.Text()
		items := strings.Split(lineContent, " ")

		for i := 0; i < len(items); i++ {
			if items[i] == "" {
				continue
			}
			number, err := strconv.Atoi(items[i])
			if err != nil {
				fmt.Println("Error converting to number ", err)
			}
			if isFirst {
				firstList = append(firstList, number)
				isFirst = false
			} else {
				if _, ok := occurenceInSecond[number]; ok {
					occurenceInSecond[number]++
				} else {
					occurenceInSecond[number] = 1
				}
			}
		}
	}

	totalDistance := 0
	for _, number := range firstList {
		if _, ok := occurenceInSecond[number]; ok {
			fmt.Println("Number ", number, " is present ", occurenceInSecond[number], " times")
			totalDistance = totalDistance + (number * occurenceInSecond[number])
		} else {
			continue
		}
	}
	fmt.Println("Total distance is ", totalDistance)
}
