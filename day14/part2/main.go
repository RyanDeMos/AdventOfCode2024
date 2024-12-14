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
	// example_input := read_input("./day14/part2/inputs/example_input.txt")
	// example_total := part2(example_input, 11, 7)
	// log.Printf("Example total: %v\n", example_total)

	input := read_input("./day14/part2/inputs/input.txt")
	total := part2(input, 101, 103)
	log.Printf("Total: %v\n", total)
}

func read_input(inputFile string) []*robot {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	robots := []*robot{}
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

		robots = append(robots, &robot{
			starting_position,
			direction,
		})
	}
	return robots
}

func part2(robots []*robot, grid_width, grid_height int) int {
	i := 0
	for {
		pos_map := map[position]bool{}
		for _, r := range robots {
			r.move(1, grid_width, grid_height)
		}
		for _, r := range robots {
			pos_map[r.location] = true
		}
		i += 1
		// Dumb solution to see if all the robots are on different spots then maybe it makes a christmas tree
		if len(pos_map) == len(robots) {
			break
		}
	}
	return i
}

func (r *robot) move(seconds, grid_width, grid_height int) {
	r.location = position{
		(r.location.x + r.direction.x*seconds) % grid_width,
		(r.location.y + r.direction.y*seconds) % grid_height,
	}
	if r.location.x < 0 {
		r.location.x = grid_width + r.location.x
	}
	if r.location.y < 0 {
		r.location.y = grid_height + r.location.y
	}
}
