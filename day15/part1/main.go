package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type position struct {
	x, y int
}

func main() {
	small_robot, small_walls, small_boxes, small_movements := read_input("./day15/part1/inputs/small_example.txt")
	small_final_box_locations := part1(small_robot, small_walls, small_boxes, small_movements)
	log.Printf("Small example total: %v\n", small_final_box_locations)

	example_robot, example_walls, example_boxes, example_movements := read_input("./day15/part1/inputs/example_input.txt")
	example_final_box_locations := part1(example_robot, example_walls, example_boxes, example_movements)
	log.Printf("Larger example total: %v\n", example_final_box_locations)

	robot, walls, boxes, movements := read_input("./day15/part1/inputs/input.txt")
	final_box_locations := part1(robot, walls, boxes, movements)
	log.Printf("Total: %v\n", final_box_locations)
}

func read_input(inputFile string) (position, map[position]bool, map[position]bool, []string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	walls := map[position]bool{}
	boxes := map[position]bool{}
	robot_pos := position{}
	grid_flag := true
	grid_row := 0

	movements := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := string(scanner.Bytes())
		if line == "" {
			grid_flag = false
		} else if grid_flag {
			chars := strings.Split(line, "")
			for column, char := range chars {
				pos := position{column, grid_row}
				if char == "#" {
					walls[pos] = true
				} else if char == "O" {
					boxes[pos] = true
				} else if char == "@" {
					robot_pos = position{column, grid_row}
				}
			}
			grid_row += 1
		} else {
			line_movements := strings.Split(line, "")
			movements = append(movements, line_movements...)
		}
	}
	return robot_pos, walls, boxes, movements
}

func part1(robot position, walls, boxes map[position]bool, movements []string) int {
	movement_map := map[string]position{
		"<": {-1, 0},
		">": {1, 0},
		"^": {0, -1},
		"v": {0, 1},
	}
	for _, movement := range movements {
		direction := movement_map[movement]
		next_pos := position{robot.x + direction.x, robot.y + direction.y}
		if boxes[next_pos] {
			first_pos := next_pos
			for boxes[next_pos] {
				next_pos = position{next_pos.x + direction.x, next_pos.y + direction.y}
			}
			if !walls[next_pos] {
				robot = first_pos
				delete(boxes, first_pos)
				boxes[next_pos] = true
			}
		} else if !walls[next_pos] {
			robot = next_pos
		}
	}
	total := 0
	for pos := range boxes {
		total += 100*pos.y + pos.x
	}
	return total
}
