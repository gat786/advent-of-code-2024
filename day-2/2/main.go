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

func removeIndex(s *[]int, index int) []int {
	newCopy := make([]int, len(*s)-1)
	copiedTill := 0
	for currentIndex := range *s {
		if index == currentIndex {
			continue
		} else {
			newCopy[copiedTill] = (*s)[currentIndex]
			copiedTill++
		}
	}
	return newCopy
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
		items_length := len(items_int)
		combinations := make([][]int, 0)
		combinations = append(combinations, items_int)
		for i := 0; i < items_length; i++ {
			comb := removeIndex(&items_int, i)
			combinations = append(combinations, comb)
		}
		fmt.Println(combinations)
		valid_combo_found := false
		for i := 0; i < len(combinations); i++ {
			is_combination_valid := walk(combinations[i])
			if is_combination_valid {
				fmt.Printf("Combination: %v is valid\n", combinations[i])
				valid_combo_found = true
			}
			if valid_combo_found {
				break
			}
		}
		if valid_combo_found {
			fmt.Println("Path is valid")
			valid_count++
		} else {
			fmt.Println("Path is invalid")
		}
	}

	fmt.Printf("Valid Count: %d\n", valid_count)
}
