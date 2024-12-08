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
	example_input := read_input("./day08/part2/inputs/example_input.txt")
	example_total := part2(example_input)
	log.Printf("Example total: %v\n", example_total)

	input := read_input("./day08/part2/inputs/input.txt")
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
		lines = append(lines, string(scanner.Bytes()))
	}
	return lines
}

func part2(lines []string) int {
	antennas := parse_input(lines)
	antinodes := find_all_antinodes(antennas, len(lines[0]), len(lines))
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

func find_all_antinodes(antennas map[string][]position, maxX, maxY int) map[position]bool {
	antinodes := map[position]bool{}
	for _, positions := range antennas {
		for idx1 := range positions {
			for idx2 := range idx1 {
				new_antis := find_antinode_pair(positions[idx1], positions[idx2], maxX, maxY)
				for _, anti := range new_antis {
					antinodes[anti] = true
				}
			}
		}
	}
	return antinodes
}

func find_antinode_pair(pos1, pos2 position, maxX, maxY int) []position {
	xDif := pos1.x - pos2.x
	yDif := pos1.y - pos2.y
	antinodes := []position{}

	// direction 1
	i := 0
	for pos1.x+i*xDif >= 0 && pos1.y+i*yDif >= 0 && pos1.x+i*xDif < maxX && pos1.y+i*yDif < maxY {
		antinodes = append(antinodes, position{pos1.x + i*xDif, pos1.y + i*yDif})
		i++
	}

	// direction 2
	i = 0
	for pos2.x-i*xDif >= 0 && pos2.y-i*yDif >= 0 && pos2.x-i*xDif < maxX && pos2.y-i*yDif < maxY {
		antinodes = append(antinodes, position{pos2.x - i*xDif, pos2.y - i*yDif})
		i++
	}
	return antinodes
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
