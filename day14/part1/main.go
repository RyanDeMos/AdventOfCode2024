package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x, y int
}

type robot struct {
	location, direction position
}

func main() {
	example_input := read_input("./day14/part1/inputs/example_input.txt")
	example_total := part1(example_input, 11, 7)
	log.Printf("Example total: %v\n", example_total)

	input := read_input("./day14/part1/inputs/input.txt")
	total := part1(input, 101, 103)
	log.Printf("Total: %v\n", total)
}

func read_input(inputFile string) []robot {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	robots := []robot{}
	for scanner.Scan() {
		line := strings.Fields(string(scanner.Bytes()))
		start := strings.Split(line[0], ",")
		start_x, _ := strconv.Atoi(start[0][2:])
		start_y, _ := strconv.Atoi(start[1][:])
		starting_position := position{
			start_x,
			start_y,
		}

		dir := strings.Split(line[1], ",")
		dir_x, _ := strconv.Atoi(dir[0][2:])
		dir_y, _ := strconv.Atoi(dir[1][:])
		direction := position{
			dir_x,
			dir_y,
		}

		robots = append(robots, robot{
			starting_position,
			direction,
		})
	}
	return robots
}

func part1(robots []robot, grid_width, grid_height int) int {
	// log.Printf("Robots: %v\n", robots)
	quadrant_map := map[string]int{}
	for _, robot := range robots {
		robot.move(100, grid_width, grid_height)
		quadrant_map[calculate_quadrant(robot, grid_width, grid_height)] += 1
	}
	// log.Printf("Quadrant Map: %v\n", quadrant_map)

	safety_factor := 1
	for quad, val := range quadrant_map {
		if quad != "none" {
			safety_factor *= val
		}
	}
	return safety_factor
}

func (r *robot) move(seconds, grid_width, grid_height int) {
	r.location = position{
		(r.location.x + r.direction.x*seconds) % grid_width,
		(r.location.y + r.direction.y*seconds) % grid_height,
	}
	if r.location.x < 0 {
		r.location.x += grid_width
	}
	if r.location.y < 0 {
		r.location.y += grid_height
	}
}

func calculate_quadrant(robot robot, grid_width, grid_height int) string {
	horizontal_middle := int(grid_height / 2)
	verticle_middle := int(grid_width / 2)

	// Top Left
	if robot.location.x < verticle_middle && robot.location.y < horizontal_middle {
		return "top_left"
	}

	// Top Right
	if robot.location.x > verticle_middle && robot.location.y < horizontal_middle {
		return "top_right"
	}

	// Bottom Left
	if robot.location.x < verticle_middle && robot.location.y > horizontal_middle {
		return "bottom_left"
	}

	// Bottom Right
	if robot.location.x > verticle_middle && robot.location.y > horizontal_middle {
		return "bottom_right"
	}

	return "none"
}
