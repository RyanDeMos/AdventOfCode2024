package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	example_input := read_input("./day04/part2/inputs/example_input.txt")
	log.Printf("Example Input: %v\n", example_input)
	example_total := part2(example_input)
	log.Printf("Example total: %v\n", example_total)

	input := read_input("./day04/part2/inputs/input.txt")
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
	total := search_x_mas(lines)
	return total
}

func search_x_mas(lines []string) int {
	total := 0
	for row_number, line := range lines {
		for column_number, character := range line {
			MAS_count := 0
			if string(character) == "A" {
				// If Top left is an "M"
				if row_number-1 >= 0 && column_number-1 >= 0 && string(lines[row_number-1][column_number-1]) == "M" {
					if row_number+1 < len(lines) && column_number+1 < len(line) && string(lines[row_number+1][column_number+1]) == "S" {
						MAS_count += 1
					}
				}
				// If Top right is an "M"
				if row_number-1 >= 0 && column_number+1 < len(line) && string(lines[row_number-1][column_number+1]) == "M" {
					if row_number+1 < len(lines) && column_number-1 >= 0 && string(lines[row_number+1][column_number-1]) == "S" {
						MAS_count += 1
					}
				}
				// If Bottom left is an "M"
				if row_number+1 < len(lines) && column_number-1 >= 0 && string(lines[row_number+1][column_number-1]) == "M" {
					if row_number-1 >= 0 && column_number+1 < len(line) && string(lines[row_number-1][column_number+1]) == "S" {
						MAS_count += 1
					}
				}
				// If Bottom right is an "M"
				if row_number+1 < len(lines) && column_number+1 < len(line) && string(lines[row_number+1][column_number+1]) == "M" {
					if row_number-1 >= 0 && column_number-1 >= 0 && string(lines[row_number-1][column_number-1]) == "S" {
						MAS_count += 1
					}
				}

				if MAS_count == 2 {
					total += 1
				}
			}
		}
	}
	return total
}
