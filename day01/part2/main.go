package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2_example_solution := part2("./day01/part2/inputs/example_input.txt")
	log.Printf("Part 2 example solution: %v\n", part2_example_solution)
	part2_solution := part2("./day01/part2/inputs/input.txt")
	log.Printf("Part 2 solution: %v\n", part2_solution)
}

func part2(inputFile string) int {
	left, right := create_paired_lists(inputFile)
	// log.Printf("\n%v\n%v\n", left, right)
	right_map := create_right_map(right)
	// log.Printf("%v\n", right_map)
	similarity_scores := find_similarity_scores(left, right_map)
	// log.Printf("%v\n", similarity_scores)
	total_sim_score := sum_list(similarity_scores)
	return total_sim_score
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

func create_right_map(right []int) map[int]int {
	right_map := map[int]int{}
	for i := 0; i < len(right); i++ {
		_, ok := right_map[right[i]]
		if ok {
			right_map[right[i]] += 1
		} else {
			right_map[right[i]] = 1
		}
	}
	return right_map
}

func find_similarity_scores(left []int, right_map map[int]int) []int {
	similarity_scores := []int{}
	for i := 0; i < len(left); i++ {
		val, ok := right_map[left[i]]
		if ok {
			similarity_scores = append(similarity_scores, left[i]*val)
		} else {
			similarity_scores = append(similarity_scores, 0)
		}
	}
	return similarity_scores
}

func sum_list(list []int) int {
	total := 0
	for i := 0; i < len(list); i++ {
		total += list[i]
	}
	return total
}
