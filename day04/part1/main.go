package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	example_input := read_input("./day04/part1/inputs/example_input.txt")
	log.Printf("Example Input: %v\n", example_input)
	example_total := part1(example_input)
	log.Printf("Example total: %v\n", example_total)

	input := read_input("./day04/part1/inputs/input.txt")
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
	total := search_xmas(lines)
	return total
}

func search_xmas(lines []string) int {
	total := 0
	for row_number, line := range lines {
		for column_number, character := range line {
			if string(character) == "X" {
				// Search to the right
				if column_number+1 < len(line) && string(line[column_number+1]) == "M" {
					if column_number+2 < len(line) && string(line[column_number+2]) == "A" {
						if column_number+3 < len(line) && string(line[column_number+3]) == "S" {
							total += 1
						}
					}
				}
				// Search to the Left
				if column_number-1 >= 0 && string(line[column_number-1]) == "M" {
					if column_number-2 >= 0 && string(line[column_number-2]) == "A" {
						if column_number-3 >= 0 && string(line[column_number-3]) == "S" {
							total += 1
						}
					}
				}
				// Search Down
				if row_number+1 < len(lines) && string(lines[row_number+1][column_number]) == "M" {
					if row_number+2 < len(lines) && string(lines[row_number+2][column_number]) == "A" {
						if row_number+3 < len(lines) && string(lines[row_number+3][column_number]) == "S" {
							total += 1
						}
					}
				}
				// Search Up
				if row_number-1 >= 0 && string(lines[row_number-1][column_number]) == "M" {
					if row_number-2 >= 0 && string(lines[row_number-2][column_number]) == "A" {
						if row_number-3 >= 0 && string(lines[row_number-3][column_number]) == "S" {
							total += 1
						}
					}
				}
				// Search Up & Right
				if row_number-1 >= 0 && column_number+1 < len(line) && string(lines[row_number-1][column_number+1]) == "M" {
					if row_number-2 >= 0 && column_number+2 < len(line) && string(lines[row_number-2][column_number+2]) == "A" {
						if row_number-3 >= 0 && column_number+3 < len(line) && string(lines[row_number-3][column_number+3]) == "S" {
							total += 1
						}
					}
				}
				// Search Up & Left
				if row_number-1 >= 0 && column_number-1 >= 0 && string(lines[row_number-1][column_number-1]) == "M" {
					if row_number-2 >= 0 && column_number-2 >= 0 && string(lines[row_number-2][column_number-2]) == "A" {
						if row_number-3 >= 0 && column_number-3 >= 0 && string(lines[row_number-3][column_number-3]) == "S" {
							total += 1
						}
					}
				}
				// Search Down & Right
				if row_number+1 < len(lines) && column_number+1 < len(line) && string(lines[row_number+1][column_number+1]) == "M" {
					if row_number+2 < len(lines) && column_number+2 < len(line) && string(lines[row_number+2][column_number+2]) == "A" {
						if row_number+3 < len(lines) && column_number+3 < len(line) && string(lines[row_number+3][column_number+3]) == "S" {
							total += 1
						}
					}
				}
				// Search Down & Left
				if row_number+1 < len(lines) && column_number-1 >= 0 && string(lines[row_number+1][column_number-1]) == "M" {
					if row_number+2 < len(lines) && column_number-2 >= 0 && string(lines[row_number+2][column_number-2]) == "A" {
						if row_number+3 < len(lines) && column_number-3 >= 0 && string(lines[row_number+3][column_number-3]) == "S" {
							total += 1
						}
					}
				}
			}
		}
	}
	return total
}
