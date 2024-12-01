package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1_example_solution := part1("./day01/part1/inputs/example_input.txt")
	log.Printf("Part 1 example solution: %v\n", part1_example_solution)
	part1_solution := part1("./day01/part1/inputs/input.txt")
	log.Printf("Part 1 solution: %v\n", part1_solution)
}

func part1(inputFile string) int {
	left, right := create_paired_lists(inputFile)
	sort.Ints(left)
	sort.Ints(right)
	diffs := find_paired_diffs(left, right)
	total := sum_diffs(diffs)
	return total
}

func create_paired_lists(inputFile string) ([]int, []int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	left := []int{}
	right := []int{}
	for scanner.Scan() {
		pairs := strings.Fields(string(scanner.Bytes()))

		left_int, err := strconv.Atoi(pairs[0])
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, left_int)

		right_int, err := strconv.Atoi(pairs[1])
		if err != nil {
			log.Fatal(err)
		}
		right = append(right, right_int)
	}

	return left, right
}

func find_paired_diffs(left []int, right []int) []int {
	if len(left) != len(right) {
		log.Fatal("Uh oh, uneven list lengths")
	}
	diffs := []int{}
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			diff = diff * -1
		}
		diffs = append(diffs, diff)
	}
	return diffs
}

func sum_diffs(diffs []int) int {
	total := 0
	for i := 0; i < len(diffs); i++ {
		total += diffs[i]
	}
	return total
}
