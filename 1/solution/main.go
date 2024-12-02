package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	root := "input"

	fp, err := os.Open(fmt.Sprintf("%s/%s", root, "data"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	firstList := make([]int, 0)
	secondList := make([]int, 0)
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
				secondList = append(secondList, number)
				isFirst = true
			}
		}
	}

	sort.Ints(firstList)
	sort.Ints(secondList)

	total := 0
	for i := 0; i < len(firstList); i++ {
		dist := abs(firstList[i] - secondList[i])
		total += dist
	}
	fmt.Println(total)
}
