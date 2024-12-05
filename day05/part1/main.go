package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	example_input := read_input("./day05/part1/inputs/example_input.txt")
	example_total := part1(example_input)
	log.Printf("Example total: %v\n", example_total)

	input := read_input("./day05/part1/inputs/input.txt")
	total := part1(input)
	log.Printf("Total: %v\n", total)
}

type rule struct {
	page_number int
	before      map[int]bool
	after       map[int]bool
}

func read_input(inputFile string) []string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line_string := string(scanner.Bytes())
		lines = append(lines, line_string)
	}
	return lines
}

func part1(lines []string) int {
	pairs, orders := parse_input(lines)
	rules := create_rules(pairs)
	valid_orders := validate_order(orders, rules)
	total := sum_middle_page(valid_orders)
	return total
}

func parse_input(lines []string) ([][]int, [][]int) {
	before_contained_pages := true
	page_pairs := [][]int{}
	updates := [][]int{}
	for _, line := range lines {
		if line == "" {
			before_contained_pages = false
		} else if before_contained_pages {
			page_pair_as_string := strings.Split(line, "|")
			page_pair := []int{}
			for _, page := range page_pair_as_string {
				as_int, err := strconv.Atoi(page)
				if err != nil {
					log.Fatal(err)
				}
				page_pair = append(page_pair, as_int)
			}
			page_pairs = append(page_pairs, page_pair)
		} else if !before_contained_pages {
			contained_pages_as_strings := strings.Split(line, ",")
			new_update := []int{}
			for _, page_number := range contained_pages_as_strings {
				as_int, err := strconv.Atoi(page_number)
				if err != nil {
					log.Fatal(err)
				}
				new_update = append(new_update, as_int)
			}
			updates = append(updates, new_update)
		}
	}

	return page_pairs, updates
}

func create_rules(pairs [][]int) map[int]rule {
	all_rules := map[int]rule{}
	for _, pair := range pairs {
		val, ok := all_rules[pair[0]]
		if !ok {
			all_rules[pair[0]] = rule{
				page_number: pair[0],
				before:      map[int]bool{},
				after: map[int]bool{
					pair[1]: true,
				},
			}
		} else {
			val.after[pair[1]] = true
		}

		val2, ok := all_rules[pair[1]]
		if !ok {
			all_rules[pair[1]] = rule{
				page_number: pair[1],
				before: map[int]bool{
					pair[0]: true,
				},
				after: map[int]bool{},
			}
		} else {
			val2.before[pair[0]] = true
		}
	}

	return all_rules
}

// returns valid orders
func validate_order(orders [][]int, rules map[int]rule) [][]int {
	valid_orders := [][]int{}

Valid_order_loop:
	for _, order := range orders {
		for index, page_number := range order {
			// Validate before
			for _, before_number := range order[0:index] {
				_, ok := rules[page_number].before[before_number]
				if !ok {
					continue Valid_order_loop
				}
			}
			// Validate After
			if index+1 < len(order) {
				for _, after_number := range order[index+1:] {
					_, ok := rules[page_number].after[after_number]
					if !ok {
						continue Valid_order_loop
					}
				}
			}
		}
		valid_orders = append(valid_orders, order)
	}
	return valid_orders
}

func sum_middle_page(valid_orders [][]int) int {
	total := 0
	for _, order := range valid_orders {
		total += order[int(len(order)/2)]
	}
	return total
}
