package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func Timer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("took %v\n", time.Since(start))
	}
}

func main() {
	example_input := read_input("./day09/part2/inputs/example_input.txt")
	example_total := part2(example_input)
	log.Printf("Example total: %v\n", example_total)

	defer Timer()()
	input := read_input("./day09/part2/inputs/input.txt")
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

func part2(lines []string) int64 {
	line := lines[0] //Input is always one line
	// log.Printf("Line: %v\n", line)
	blocks := input_to_blocks(line)
	// log.Printf("Blocks:\n %v\n", blocks)
	after_move_blocks := move_blocks(blocks)
	// log.Printf("After move Blocks:\n %v\n", after_move_blocks)
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
	// after_move := []int{}
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

	after_move := make([]int, len(blocks))
	copy(after_move, blocks)

	current_block_int := -1
	current_block_length := 0

	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i] != current_block_int {
			if current_block_int != -1 {
				// Go through original blocks and find first place this block fits
				starting_fill_index := find_fill_block(after_move, current_block_length)
				// log.Printf("Starting index: %v\n", starting_fill_index)
				if starting_fill_index != -1 && starting_fill_index < i {
					// Replace fill block
					for j := starting_fill_index; j < starting_fill_index+current_block_length; j++ {
						after_move[j] = current_block_int
					}

					// Go through copy and remove numbers
					for j := i + 1; j < i+1+current_block_length; j++ {
						after_move[j] = -2
					}
				}
				// log.Printf("After single movement:\n %v\n", after_move)

			}

			// Switch to new value
			current_block_int = blocks[i]
			current_block_length = 1
		} else {
			current_block_length += 1
		}
		// log.Printf("Current Int: %v\n", current_block_int)
		// log.Printf("Current Length: %v\n", current_block_length)
	}
	return after_move
}

func find_fill_block(blocks []int, required_length int) int {
	fill_starting_index := -1
	fill_block_length := 0
	for j := 0; j < len(blocks); j++ {
		if blocks[j] == -1 && fill_block_length == 0 {
			fill_starting_index = j
			fill_block_length += 1
		} else if blocks[j] == -1 {
			fill_block_length += 1
		} else {
			fill_block_length = 0
		}

		if fill_block_length >= required_length {
			return fill_starting_index
		}
	}
	return -1
}

func calculate_checksum(blocks []int) int64 {
	total := int64(0)
	for idx, file_id := range blocks {
		if file_id != -1 && file_id != -2 {
			total += int64(idx) * int64(file_id)
		}
	}
	return total
}
