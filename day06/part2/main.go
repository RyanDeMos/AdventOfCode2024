package main

import (
	"bufio"
	"log"
	"os"
)

type position struct {
	x, y int
}

type guard struct {
	current_pos         position
	direction_of_travel position // {1,0}, {-1,0}, {0,1}, {0,-1}
	locations_visited   map[position]position
}

func main() {
	example_input := read_input("./day06/part2/inputs/example_input.txt")
	example_total := part2(example_input)
	log.Printf("Example total: %v\n", example_total)

	input := read_input("./day06/part2/inputs/input.txt")
	total := part2(input)
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
		line_string := string(scanner.Bytes())
		lines = append(lines, line_string)
	}
	return lines
}

func part2(lines []string) int {
	guard := find_guard(lines)
	guard.move_guard(lines)
	first_trip_positions := []position{}
	for key := range guard.locations_visited {
		first_trip_positions = append(first_trip_positions, key)
	}
	loops := find_all_loops(lines, first_trip_positions)
	return loops
}

func find_guard(lines []string) guard {
	for row, line := range lines {
		for column, character := range line {
			if string(character) == "^" {
				return guard{
					current_pos: position{
						x: column,
						y: row,
					},
					direction_of_travel: position{0, -1},
					locations_visited:   map[position]position{{column, row}: {0, -1}},
				}
			}
		}
	}
	log.Fatal("Failed to find guard")
	return guard{}
}

// returns 1 if we entered a loop, 0 otherwise
func (g *guard) move_guard(lines []string) int {
	//map[current direction]new direction 90 deg clockwise
	next_direction_map := map[position]position{
		{0, -1}: {1, 0},
		{1, 0}:  {0, 1},
		{0, 1}:  {-1, 0},
		{-1, 0}: {0, -1},
	}
	for g.current_pos.y+g.direction_of_travel.y >= 0 && g.current_pos.x+g.direction_of_travel.x >= 0 && g.current_pos.y+g.direction_of_travel.y < len(lines) && g.current_pos.x+g.direction_of_travel.x < len(lines[g.current_pos.y+g.direction_of_travel.y]) {
		next_position := lines[g.current_pos.y+g.direction_of_travel.y][g.current_pos.x+g.direction_of_travel.x]
		if string(next_position) == "#" {
			// Change direction
			g.direction_of_travel = next_direction_map[g.direction_of_travel]
		} else {
			g.current_pos = position{
				y: g.current_pos.y + g.direction_of_travel.y,
				x: g.current_pos.x + g.direction_of_travel.x,
			}
			val, ok := g.locations_visited[g.current_pos]
			if ok && val == g.direction_of_travel {
				// Same position, same direction ==> We are in a loop
				return 1
			}
			g.locations_visited[g.current_pos] = g.direction_of_travel
		}
	}
	return 0
}

func find_all_loops(lines []string, first_trip_positions []position) int {
	total_loops := 0
	for _, location_to_replace := range first_trip_positions {
		guard := find_guard(lines)
		// Cant place one on guards starting location skip that loop
		if location_to_replace != guard.current_pos {
			lines_copy := make([]string, len(lines))
			copy(lines_copy, lines)
			// Replace the location
			if location_to_replace.x+1 < len(lines_copy[location_to_replace.y]) {
				lines_copy[location_to_replace.y] = lines_copy[location_to_replace.y][0:location_to_replace.x] + "#" + lines_copy[location_to_replace.y][location_to_replace.x+1:]
			} else {
				lines_copy[location_to_replace.y] = lines_copy[location_to_replace.y][0:location_to_replace.x] + "#"
			}
			total_loops += guard.move_guard(lines_copy)
		}
	}
	return total_loops
}
