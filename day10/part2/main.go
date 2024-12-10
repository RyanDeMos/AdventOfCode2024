package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type position struct {
	x, y int
}

func main() {
	example_input := read_input("./day10/part1/inputs/example_input.txt")
	example_total := part1(example_input)
	log.Printf("Example total: %v\n", example_total)

	input := read_input("./day10/part1/inputs/input.txt")
	total := part1(input)
	log.Printf("Total: %v\n", total)
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
		lines = append(lines, string(scanner.Bytes()))
	}
	return lines
}

func part1(lines []string) int {
	lines_as_ints := transform_input_to_int_array(lines)
	trailheads := find_trailheads(lines_as_ints)
	orders := find_orders(lines_as_ints, trailheads)
	return find_total(orders)
}

func transform_input_to_int_array(lines []string) [][]int {
	int_arr := [][]int{}
	for _, line := range lines {
		sub_arr := []int{}
		for _, char := range line {
			as_int, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal(err)
			}
			sub_arr = append(sub_arr, as_int)
		}
		int_arr = append(int_arr, sub_arr)
	}
	return int_arr
}

func find_trailheads(lines [][]int) []position {
	trailheads := []position{}
	for row, line := range lines {
		for column, char := range line {
			if char == 0 {
				trailheads = append(trailheads, position{column, row})
			}
		}
	}
	return trailheads
}

func find_orders(lines [][]int, trailheads []position) []int {
	orders := []int{}
	for _, trailhead := range trailheads {
		all_surrounding := map[position]int{
			trailhead: 1,
		}
		for i := 1; i <= 9; i++ {
			new_surrounding := map[position]int{}
			for location, rating := range all_surrounding {
				next_location := find_surrounding(lines, location, i, rating)
				for new_location, new_rating := range next_location {
					new_surrounding[new_location] += new_rating
				}
			}
			all_surrounding = new_surrounding
		}
		order := 0
		for _, val := range all_surrounding {
			order += val
		}
		orders = append(orders, order)
	}
	return orders
}

func find_surrounding(lines [][]int, location position, next_number int, rating int) map[position]int {
	surrounding_locations := map[position]int{}
	if location.y-1 >= 0 && lines[location.y-1][location.x] == next_number {
		surrounding_locations[position{location.x, location.y - 1}] = rating
	}
	if location.y+1 < len(lines) && lines[location.y+1][location.x] == next_number {
		surrounding_locations[position{location.x, location.y + 1}] = rating
	}
	if location.x-1 >= 0 && lines[location.y][location.x-1] == next_number {
		surrounding_locations[position{location.x - 1, location.y}] = rating
	}
	if location.x+1 < len(lines[0]) && lines[location.y][location.x+1] == next_number {
		surrounding_locations[position{location.x + 1, location.y}] = rating
	}

	return surrounding_locations
}

func find_total(orders []int) int {
	total := 0
	for _, order := range orders {
		total += order
	}
	return total
}
