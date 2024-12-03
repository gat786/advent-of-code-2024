package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Order struct {
	Increasing bool
	Decreasing bool
	Invalid    bool
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func isIncreasingValidly(num1 int, num2 int) bool {
	if num1 > num2 {
		difference := abs(num1 - num2)
		if difference > 0 && difference < 4 {
			return true
		}
	}
	return false
}

func isDecreasingValidly(num1 int, num2 int) bool {
	if num1 < num2 {
		difference := abs(num1 - num2)
		if difference > 0 && difference < 4 {
			return true
		}
	}
	return false
}

func walk(arr []int) bool {
	var orderOfElements Order
	previousNumber := 0
	currentNumber := 0
	orderSet := false
	for i := 1; i < len(arr); i++ {
		previousNumber = arr[i-1]
		currentNumber = arr[i]
		if !orderSet {
			if currentNumber > previousNumber {
				orderOfElements.Increasing = true
				orderSet = true
			} else if currentNumber < previousNumber {
				orderOfElements.Decreasing = true
				orderSet = true
			} else {
				orderOfElements.Invalid = true
				orderSet = true
			}
		}
		if orderOfElements.Invalid {
			return false
		} else {
			if orderOfElements.Increasing {
				if isIncreasingValidly(currentNumber, previousNumber) {
					continue
				} else {
					return false
				}
			} else if orderOfElements.Decreasing {
				if isDecreasingValidly(currentNumber, previousNumber) {
					continue
				} else {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Reading files from the current directory")
	root := "/Users/gat786/projects/personal/golang/advent-of-code-2024/input"

	fp, err := os.Open(fmt.Sprintf("%s/%s", root, "data-2"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	valid_count := 0
	for scanner.Scan() {
		lineContent := scanner.Text()
		items := strings.Split(lineContent, " ")
		items_int := make([]int, len(items))
		for i := 0; i < len(items); i++ {
			num, err := strconv.Atoi(items[i])
			if err != nil {
				fmt.Printf("Error converting str to number %s", err)
				return
			}
			items_int[i] = num
		}
		is_valid := walk(items_int)
		fmt.Printf("Is Valid: %v\n", is_valid)
		if is_valid {
			valid_count++
		}
	}

	fmt.Printf("Valid Count: %d\n", valid_count)
}
