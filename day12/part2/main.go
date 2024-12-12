package main

import (
	"bufio"
	"log"
	"maps"
	"os"
	"slices"
	"strings"
)

type position struct {
	x, y int
}

func main() {
	example_input := read_input("./day12/part2/inputs/example_input.txt")
	example_total := part2(example_input)
	log.Printf("Example total: %v\n", example_total)

	input := read_input("./day12/part2/inputs/input.txt")
	total := part2(input)
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

func part2(lines [][]string) int {
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
	sides := 0
	top_sides, bottom_sides, left_sides, right_sides := []position{}, []position{}, []position{}, []position{}
	for current_position := range region {
		area += 1
		if !region[position{current_position.x - 1, current_position.y}] {
			left_sides = append(left_sides, current_position)
		}
		if !region[position{current_position.x + 1, current_position.y}] {
			right_sides = append(right_sides, current_position)
		}
		if !region[position{current_position.x, current_position.y - 1}] {
			top_sides = append(top_sides, current_position)
		}
		if !region[position{current_position.x, current_position.y + 1}] {
			bottom_sides = append(bottom_sides, current_position)
		}
	}

	// Get top Sides
	slices.SortFunc(top_sides, func(pos1, pos2 position) int {
		if pos1.y < pos2.y {
			return -1
		} else if pos1.y > pos2.y {
			return 1
		} else {
			if pos1.x < pos2.x {
				return -1
			} else if pos1.x > pos2.x {
				return 1
			}
		}
		return 0
	})
	current_pos := position{-1, -1}
	top_side_count := 0
	for _, top_pos := range top_sides {
		if current_pos.y != top_pos.y || current_pos.x-top_pos.x < -1 {
			top_side_count += 1
		}
		current_pos = top_pos
	}

	// Get Bottom Sides
	slices.SortFunc(bottom_sides, func(pos1, pos2 position) int {
		if pos1.y < pos2.y {
			return -1
		} else if pos1.y > pos2.y {
			return 1
		} else {
			if pos1.x < pos2.x {
				return -1
			} else if pos1.x > pos2.x {
				return 1
			}
		}
		return 0
	})
	current_pos = position{-1, -1}
	bottom_side_count := 0
	for _, top_pos := range bottom_sides {
		if current_pos.y != top_pos.y || current_pos.x-top_pos.x < -1 {
			bottom_side_count += 1
		}
		current_pos = top_pos
	}

	// Get Left Sides
	slices.SortFunc(left_sides, func(pos1, pos2 position) int {
		if pos1.x < pos2.x {
			return -1
		} else if pos1.x > pos2.x {
			return 1
		} else {
			if pos1.y < pos2.y {
				return -1
			} else if pos1.y > pos2.y {
				return 1
			}
		}
		return 0
	})
	current_pos = position{-1, -1}
	left_side_count := 0
	for _, top_pos := range left_sides {
		if current_pos.x != top_pos.x || current_pos.y-top_pos.y < -1 {
			left_side_count += 1
		}
		current_pos = top_pos
	}

	// Get Right Sides
	slices.SortFunc(right_sides, func(pos1, pos2 position) int {
		if pos1.x < pos2.x {
			return -1
		} else if pos1.x > pos2.x {
			return 1
		} else {
			if pos1.y < pos2.y {
				return -1
			} else if pos1.y > pos2.y {
				return 1
			}
		}
		return 0
	})
	current_pos = position{-1, -1}
	right_side_count := 0
	for _, top_pos := range right_sides {
		if current_pos.x != top_pos.x || current_pos.y-top_pos.y < -1 {
			right_side_count += 1
		}
		current_pos = top_pos
	}

	sides = left_side_count + right_side_count + top_side_count + bottom_side_count
	return area * sides
}
