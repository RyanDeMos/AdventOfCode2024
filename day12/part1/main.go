package main

import (
	"bufio"
	"log"
	"maps"
	"os"
	"strings"
)

type position struct {
	x, y int
}

func main() {
	example_input := read_input("./day12/part1/inputs/example_input.txt")
	example_total := part1(example_input)
	log.Printf("Example total: %v\n", example_total)

	input := read_input("./day12/part1/inputs/input.txt")
	total := part1(input)
	log.Printf("Total: %v\n", total)
}

func read_input(inputFile string) [][]string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := [][]string{}
	for scanner.Scan() {
		lines = append(lines, strings.Split(string(scanner.Bytes()), ""))
	}
	return lines
}

func part1(lines [][]string) int {
	total_price := 0
	positions_seen_in_a_region := map[position]bool{}
	for row, line := range lines {
		for column, plot := range line {
			if !positions_seen_in_a_region[position{column, row}] {
				new_region := find_region(positions_seen_in_a_region, position{column, row}, plot, lines)
				price := find_price(new_region)
				total_price += price
			}
		}
	}
	return total_price
}

func find_region(positions_seen_in_a_region map[position]bool, current_position position, plot string, lines [][]string) map[position]bool {
	positions_in_this_region := map[position]bool{
		current_position: true,
	}
	if current_position.x-1 >= 0 && lines[current_position.y][current_position.x-1] == plot && !positions_seen_in_a_region[position{current_position.x - 1, current_position.y}] {
		new_pos := position{current_position.x - 1, current_position.y}
		positions_in_this_region[new_pos] = true
		positions_seen_in_a_region[new_pos] = true

		adjacent_matching := find_region(positions_seen_in_a_region, new_pos, plot, lines)
		maps.Copy(positions_in_this_region, adjacent_matching)
	}
	if current_position.x+1 < len(lines[0]) && lines[current_position.y][current_position.x+1] == plot && !positions_seen_in_a_region[position{current_position.x + 1, current_position.y}] {
		new_pos := position{current_position.x + 1, current_position.y}
		positions_in_this_region[new_pos] = true
		positions_seen_in_a_region[new_pos] = true

		adjacent_matching := find_region(positions_seen_in_a_region, new_pos, plot, lines)
		maps.Copy(positions_in_this_region, adjacent_matching)
	}
	if current_position.y-1 >= 0 && lines[current_position.y-1][current_position.x] == plot && !positions_seen_in_a_region[position{current_position.x, current_position.y - 1}] {
		new_pos := position{current_position.x, current_position.y - 1}
		positions_in_this_region[new_pos] = true
		positions_seen_in_a_region[new_pos] = true

		adjacent_matching := find_region(positions_seen_in_a_region, new_pos, plot, lines)
		maps.Copy(positions_in_this_region, adjacent_matching)
	}
	if current_position.y+1 < len(lines) && lines[current_position.y+1][current_position.x] == plot && !positions_seen_in_a_region[position{current_position.x, current_position.y + 1}] {
		new_pos := position{current_position.x, current_position.y + 1}
		positions_in_this_region[new_pos] = true
		positions_seen_in_a_region[new_pos] = true

		adjacent_matching := find_region(positions_seen_in_a_region, new_pos, plot, lines)
		maps.Copy(positions_in_this_region, adjacent_matching)
	}
	return positions_in_this_region
}

func find_price(region map[position]bool) int {
	area := 0
	perimeter := 0
	for current_position := range region {
		area += 1
		perimeter += 4
		if region[position{current_position.x - 1, current_position.y}] {
			perimeter -= 1
		}
		if region[position{current_position.x + 1, current_position.y}] {
			perimeter -= 1
		}
		if region[position{current_position.x, current_position.y - 1}] {
			perimeter -= 1
		}
		if region[position{current_position.x, current_position.y + 1}] {
			perimeter -= 1
		}
	}
	return area * perimeter
}
