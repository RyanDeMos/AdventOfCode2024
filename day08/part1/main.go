package main

import (
	"bufio"
	"log"
	"os"
)

type position struct {
	x, y int
}

func main() {
	example_input := read_input("./day08/part1/inputs/example_input.txt")
	example_total := part1(example_input)
	log.Printf("Example total: %v\n", example_total)

	input := read_input("./day08/part1/inputs/input.txt")
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
	antennas := parse_input(lines)
	// log.Printf("Antennas: %v\n", antennas)
	antinodes := find_all_antinodes(antennas)
	// log.Printf("antinodes: %v\n", antinodes)
	total_valid := find_valid_antinodes(antinodes, lines)
	return total_valid
}

func parse_input(lines []string) map[string][]position {
	antennas := map[string][]position{}
	for row, line := range lines {
		for column, node := range line {
			if string(node) != "." {
				positions, ok := antennas[string(node)]
				if ok {
					positions = append(positions, position{column, row})
					antennas[string(node)] = positions
				} else {
					antennas[string(node)] = []position{{column, row}}
				}
			}
		}
	}
	return antennas
}

func find_all_antinodes(antennas map[string][]position) map[position]bool {
	antinodes := map[position]bool{}
	for _, positions := range antennas {
		for idx1 := range positions {
			for idx2 := range idx1 {
				anti1, anti2 := find_antinode_pair(positions[idx1], positions[idx2])
				antinodes[anti1] = true
				antinodes[anti2] = true
			}
		}
	}
	return antinodes
}

func find_antinode_pair(pos1, pos2 position) (position, position) {
	xDif := pos1.x - pos2.x
	yDif := pos1.y - pos2.y
	return position{pos1.x + xDif, pos1.y + yDif}, position{pos2.x - xDif, pos2.y - yDif}
}

func find_valid_antinodes(antinodes map[position]bool, lines []string) int {
	total := 0
	for antinode := range antinodes {
		if antinode.y >= 0 && antinode.x >= 0 && antinode.y < len(lines) && antinode.x < len(lines[0]) {
			total += 1
		}
	}
	return total
}
