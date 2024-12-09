package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Solving Day 5 Problem 1")
	fmt.Println("Reading files from the current directory")

	root := "../../input"
	file_name := "data-5"

	page_rules := make(map[string][]string, 0)
	page_updates := make([][]string, 0)

	fp, err := os.Open(fmt.Sprintf("%s/%s", root, file_name))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	rules_completed := false

	lineNum := 0
	scanner := bufio.NewScanner(fp)
	valid_count := 0
	valid_updates := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			rules_completed = true
			continue
		} else {
			if rules_completed {
				page_updates_list := strings.Split(line, ",")
				page_update_numbers := make([]string, 0)
				for _, update := range page_updates_list {
					page_update_numbers = append(page_update_numbers, update)
				}
				page_updates = append(page_updates, page_update_numbers)
			} else {
				rule_items := strings.Split(line, "|")
				if _, ok := page_rules[rule_items[0]]; ok {
					fmt.Printf("Rule for %s already exists, appending to the rules\n", rule_items[0])
					page_rules[rule_items[0]] = append(page_rules[rule_items[0]], rule_items[1])
				} else {
					fmt.Printf("Rule does not exist for %s, creating a new rule\n", rule_items[0])
					page_rules[rule_items[0]] = []string{rule_items[1]}
				}
			}
		}
		lineNum++
	}
	fmt.Println(page_rules)

	for _, update := range page_updates {
		previous_entries := make([]string, 0)
		update_invalid := false
		fmt.Println("Current update that we are processing is ", update)

		for _, update_number := range update {
			fmt.Printf("Current update number: %v\n", update_number)

			if _, ok := page_rules[update_number]; !ok {
				fmt.Printf("No rules found for %s, ignoring this entry\n", update_number)
				continue
			} else {
				fmt.Printf("Rules found for %s\n", update_number)
				rule_numbers := page_rules[update_number]
				for _, previous_entry_item := range previous_entries {
					contains := slices.Contains(rule_numbers, previous_entry_item)
					fmt.Printf("Current %s, Previous Entries: %s, Rules for current %s, Checking for %s, Contains?: %t\n", update_number, previous_entries, rule_numbers, previous_entry_item, contains)
					if contains {
						update_invalid = true
						break
					}
				}
				fmt.Println("Completed checking for previous entries")
			}
			if update_invalid {
				break
			}
			// update previous entries fields
			previous_entries = append(previous_entries, update_number)
		}
		if update_invalid {
			fmt.Println("Invalid")
		} else {
			fmt.Println("Valid")
			valid_updates = append(valid_updates, update)
			valid_count += 1
		}

		update_invalid = false
	}

	fmt.Printf("Found %d valid updates\n", valid_count)

	total_middle := 0
	for update := range valid_updates {
		middle_index := len(valid_updates[update]) / 2
		fmt.Printf("Middle Index: %d, Middle Value: %s\n", middle_index, valid_updates[update][middle_index])
		number, err := strconv.Atoi(valid_updates[update][middle_index])
		if err != nil {
			fmt.Printf("Error converting middle index value from valid updates: %s\n", err)
			return
		}
		total_middle += number
	}
	fmt.Printf("Total middle value: %d\n", total_middle)
}
