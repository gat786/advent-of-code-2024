package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	seq = map[rune]string{109: "m", 117: "u", 108: "l",
		40: "(", 41: ")", 48: "0", 49: "1", 50: "2", 51: "3", 52: "4",
		53: "5", 54: "6", 55: "7", 56: "8", 57: "9", 44: ",", 100: "d",
		111: "o", 110: "n", 116: "t"}
	seq_numbers   = map[rune]int{48: 0, 49: 1, 50: 2, 51: 3, 52: 4, 53: 5, 54: 6, 55: 7, 56: 8, 57: 9}
	comma_rune    = 44
	m             = 109
	d             = 100
	closing_brace = 41
	do_string     = "do()"
	dont_string   = "don't()"

	consider_next_multiplications = true
)

func get_number(lineContent *string, index int) (int, error) {
	number_in_string := ""
	for {
		c := rune((*lineContent)[index])
		_, ok := seq_numbers[c]
		if ok {
			number_in_string += string(c)
			index++
		} else if c == rune(comma_rune) || c == rune(41) {
			number_int, err := strconv.Atoi(number_in_string)
			if err != nil {
				return 0, fmt.Errorf("error converting string to number %s", number_in_string)
			}
			return number_int, nil
		} else {
			return 0, fmt.Errorf("invalid character found: %s", string(c))
		}
	}
}

func get_invalid_number(lineContent *string, index int) string {
	number_in_string := ""
	for {
		c := rune((*lineContent)[index])
		_, ok := seq_numbers[c]
		if ok {
			number_in_string += string(c)
			index++
		} else {
			number_in_string += string(c)
			return number_in_string
		}
	}
}

func get_condition(lineContent *string, index int) (string, error) {
	condition_string := ""
	c := (*lineContent)[index : index+3]
	c = fmt.Sprintf("%s%s", "d", c)
	if c == do_string {
		condition_string += c
		return condition_string, nil
	} else {
		c := (*lineContent)[index : index+6]
		c = fmt.Sprintf("%s%s", "d", c)
		if c == dont_string {
			condition_string += c
			return condition_string, nil
		} else {
			return "", fmt.Errorf("invalid condition found: %s", c)
		}
	}
}

func main() {
	fmt.Println("Solving Day 3 Problem 1")
	fmt.Println("Reading files from the current directory")

	root := "/Users/gat786/projects/personal/golang/advent-of-code-2024/input"
	file_name := "data-3"

	fp, err := os.Open(fmt.Sprintf("%s/%s", root, file_name))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	// mul() == 109, 117, 108, 40, 41
	// numbers == 48, 49, 50, 51, 52, 53, 54, 55, 56, 57
	// comma == 44

	scanner := bufio.NewScanner(fp)
	mul_brace_size := len("ul(")
	comb_numbers := make([][]int, 0)

	for scanner.Scan() {
		lineContent := scanner.Text()
		curr_index := 0
		// current_place := 0
		for {
			if curr_index >= len(lineContent) {
				fmt.Println("End of line reached, reading next line")
				scanner.Scan()
				lineContent = scanner.Text()
				curr_index = 0
			}
			if len(lineContent) == 0 {
				fmt.Println("Empty line found, we've reached the end of the file")
				break
			}
			c := rune(lineContent[curr_index])
			curr_index++
			if curr_index < len(lineContent) {
				_, ok := seq[c]
				// fmt.Println(index, ok)
				if ok {
					if c == rune(d) {
						fmt.Println("d found")

						condition, err := get_condition(&lineContent, curr_index)
						if err != nil {
							fmt.Println(err)
							curr_index += 1
						} else {
							fmt.Println("Condition found: ", condition)
							curr_index += len(condition) - 1
							if condition == do_string {
								consider_next_multiplications = true
							} else if condition == dont_string {
								consider_next_multiplications = false
							}
						}
					} else if c == rune(m) {
						// current letter is m, get all the next continous valid runes
						fmt.Println("m found")
						// take the next smallest_combination_size runes
						next_runes := lineContent[curr_index : curr_index+mul_brace_size]
						if next_runes == "ul(" {
							curr_index += mul_brace_size
							// get the firstNumber
							firstNumber, err := get_number(&lineContent, curr_index)
							if err != nil {
								invalid_number := get_invalid_number(&lineContent, curr_index)
								fmt.Println("Invalid number found: ", invalid_number)
								curr_index += len(invalid_number)
							} else {
								fmt.Println("Number found: ", firstNumber)
								curr_index += len(strconv.Itoa(firstNumber))
								next_char := rune(lineContent[curr_index])
								if next_char == rune(comma_rune) {
									curr_index++
									// get the secondNumber
									secondNumber, err := get_number(&lineContent, curr_index)
									if err != nil {
										invalid_number := get_invalid_number(&lineContent, curr_index)
										fmt.Println("Invalid number found: ", invalid_number)
										curr_index += len(invalid_number)
									} else {
										fmt.Println("Number found: ", secondNumber)
										curr_index += len(strconv.Itoa(secondNumber))

										// get the closing brace
										next_char := rune(lineContent[curr_index])
										if rune(closing_brace) == rune(next_char) {
											fmt.Println("Closing brace found")

											// Print complete sequence
											fmt.Printf("Complete sequence would be: mul(%d, %d)\n", firstNumber, secondNumber)
											if consider_next_multiplications {
												fmt.Println("Adding this combination to the list")
												combination := []int{firstNumber, secondNumber}
												comb_numbers = append(comb_numbers, combination)
												curr_index = curr_index + 1
											} else {
												fmt.Println("Skipping this combination, since we found dont()")
											}
										} else {
											fmt.Println("Closing brace not found, it is not a complete sequence")
										}
									}
								} else {
									fmt.Println("Comma not found, it is not a complete sequence")
								}
							}
						} else {
							curr_index++
						}
					}
				}
			} else {
				if scanner.Scan() {
					lineContent = scanner.Text()
					curr_index = 0
				} else {
					break
				}
			}
		}
	}

	fmt.Println("All the valid combinations are: ", comb_numbers)

	total := 0
	for _, comb := range comb_numbers {
		multiplication := comb[0] * comb[1]
		total += multiplication
	}

	fmt.Println("Total: ", total)
}
