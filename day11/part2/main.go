package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	example_input := read_input("./day11/part2/inputs/example_input.txt")
	example_total := part2(example_input)
	log.Printf("Example total: %v\n", example_total)

	input := read_input("./day11/part2/inputs/input.txt")
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
	stones := parse_input(lines)
	log.Printf("Stones: %v\n", stones)
	after_all_blinks := follow_rules(stones)
	return after_all_blinks
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

func follow_rules(stones []int) int {
	total_stone_map := map[int]int{}
	for _, stone := range stones {
		total_stone_map[stone] += 1
	}
	for i := 0; i < 75; i++ {
		new_stone_map := map[int]int{}
		for stone, count := range total_stone_map {
			after_changes := []int{}
			if stone == 0 {
				after_changes = append(after_changes, int(1))
			} else if digits := strconv.Itoa(int(stone)); len(digits)%2 == 0 {
				first_digits, _ := strconv.Atoi(digits[:len(digits)/2])
				second_digits, _ := strconv.Atoi(digits[len(digits)/2:])
				after_changes = append(after_changes, first_digits, second_digits)
			} else {
				after_changes = append(after_changes, stone*2024)
			}
			for _, new_stones := range after_changes {
				new_stone_map[new_stones] += count
			}
		}
		total_stone_map = new_stone_map
	}
	total := 0
	for _, sum := range total_stone_map {
		total += sum
	}
	return total
}
