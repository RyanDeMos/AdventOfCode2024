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
	locations_visited   map[position]bool
}

func main() {
	example_input := read_input("./day06/part1/inputs/example_input.txt")
	example_total := part1(example_input)
	log.Printf("Example total: %v\n", example_total)

	input := read_input("./day06/part1/inputs/input.txt")
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
		line_string := string(scanner.Bytes())
		lines = append(lines, line_string)
	}
	return lines
}

func part1(lines []string) int {
	guard := find_guard(lines)
	// log.Printf("Guard: %v\n", guard)
	guard.move_guard(lines)
	return len(guard.locations_visited)
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
					locations_visited:   map[position]bool{{column, row}: true},
				}
			}
		}
	}
	log.Fatal("Failed to find guard")
	return guard{}
}

func (g *guard) move_guard(lines []string) {
	//map[current direction]new direction 90 deg clockwise
	next_direction_map := map[position]position{
		position{0, -1}: position{1, 0},
		position{1, 0}:  position{0, 1},
		position{0, 1}:  position{-1, 0},
		position{-1, 0}: position{0, -1},
	}
	for g.current_pos.y+g.direction_of_travel.y < len(lines) && g.current_pos.x+g.direction_of_travel.x < len(lines[g.current_pos.y+g.direction_of_travel.y]) && g.current_pos.y+g.direction_of_travel.y >= 0 && g.current_pos.x+g.direction_of_travel.x >= 0 {
		next_position := lines[g.current_pos.y+g.direction_of_travel.y][g.current_pos.x+g.direction_of_travel.x]
		// log.Printf("Next Position: %v\n", string(next_position))
		if string(next_position) == "#" {
			// Change direction
			g.direction_of_travel = next_direction_map[g.direction_of_travel]
		} else {
			g.current_pos = position{
				y: g.current_pos.y + g.direction_of_travel.y,
				x: g.current_pos.x + g.direction_of_travel.x,
			}
			g.locations_visited[g.current_pos] = true
		}
	}
}
