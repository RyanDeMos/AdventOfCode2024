package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	example_input := read_input("./day09/part1/inputs/example_input.txt")
	example_total := part1(example_input)
	log.Printf("Example total: %v\n", example_total)

	input := read_input("./day09/part1/inputs/input.txt")
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

func part1(lines []string) int64 {
	line := lines[0] //Input is always one line
	blocks := input_to_blocks(line)
	after_move_blocks := move_blocks(blocks)
	checksum := calculate_checksum(after_move_blocks)
	return checksum
}

func input_to_blocks(line string) []int {
	block := []int{}
	for idx, char := range line {
		append_int := -1
		if idx%2 == 0 {
			append_int = int(idx / 2)
		}
		file_length := int(char - '0')
		for i := 0; i < file_length; i++ {
			block = append(block, append_int)
		}
	}
	return block
}

func move_blocks(blocks []int) []int {
	after_move := []int{}
	idx_to_move := len(blocks) - 1
	for blocks[idx_to_move] == -1 {
		idx_to_move--
	}

	number_of_digits := 0
	for _, number := range blocks {
		if number != -1 {
			number_of_digits++
		}
	}

	for _, number := range blocks {
		if number == -1 {
			after_move = append(after_move, blocks[idx_to_move])
			idx_to_move--
			// find next number that can be moved
			for blocks[idx_to_move] == -1 {
				idx_to_move--
			}
		} else {
			after_move = append(after_move, number)
		}
		if len(after_move) == number_of_digits {
			break
		}
	}
	return after_move
}

func calculate_checksum(blocks []int) int64 {
	total := int64(0)
	for idx, file_id := range blocks {
		total += int64(idx) * int64(file_id)
	}
	return total
}
