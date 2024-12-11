package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	example_input := read_input("./day11/part1/inputs/example_input.txt")
	example_total := part1(example_input)
	log.Printf("Example total: %v\n", example_total)

	input := read_input("./day11/part1/inputs/input.txt")
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
	stones := parse_input(lines)
	log.Printf("Stones: %v\n", stones)
	after_all_blinks := follow_rules(stones)
	// log.Printf("Stones after all blinks: %v\n", after_all_blinks)
	return len(after_all_blinks)
}

func parse_input(lines []string) []int {
	stones := []int{}
	for _, line := range lines {
		numbers := strings.Fields(line)
		for _, char := range numbers {
			stone, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal(err)
			}
			stones = append(stones, stone)
		}
	}
	return stones
}

func follow_rules(stones []int) []int {
	before_changes := make([]int, len(stones))
	copy(before_changes, stones)
	for i := 0; i < 25; i++ {
		after_changes := []int{}
		for _, stone := range before_changes {
			if stone == 0 {
				after_changes = append(after_changes, int(1))
			} else if digits := strconv.Itoa(int(stone)); len(digits)%2 == 0 {
				first_digits, err := strconv.Atoi(digits[:len(digits)/2])
				if err != nil {
					log.Fatal(err)
				}
				second_digits, err := strconv.Atoi(digits[len(digits)/2:])
				if err != nil {
					log.Fatal(err)
				}
				after_changes = append(after_changes, first_digits, second_digits)
			} else {
				after_changes = append(after_changes, stone*2024)
			}
		}
		before_changes = after_changes
	}
	return before_changes
}
