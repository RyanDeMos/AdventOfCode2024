package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	example_input := read_input("./day03/part1/inputs/example_input.txt")
	example_total := part1(example_input)
	log.Printf("Example total: %v\n", example_total)

	input := read_input("./day03/part1/inputs/input.txt")
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
	mults := mults(lines)
	total := find_total(mults)
	return total
}

func mults(lines []string) []int {
	r, err := regexp.Compile(`.*mul\((\d\d?\d?,\d\d?\d?)\).*`)
	if err != nil {
		log.Fatal(err)
	}

	mults := []int{}
	for _, line := range lines {
		for r.MatchString(line) {
			// Find an occurence and turn it into integers
			matching_numbers_string := r.FindStringSubmatch(line)[1]
			matching_numbers_list := strings.Split(matching_numbers_string, ",")
			number_1, err := strconv.Atoi(matching_numbers_list[0])
			if err != nil {
				log.Fatal(err)
			}
			number_2, err := strconv.Atoi(matching_numbers_list[1])
			if err != nil {
				log.Fatal(err)
			}
			mults = append(mults, number_1*number_2)

			// Remove this occurence from the string
			match_indexes := r.FindStringSubmatchIndex(line)
			line = line[0:match_indexes[2]] + line[match_indexes[3]:]
		}
	}
	return mults
}

func find_total(mults []int) int {
	total := 0
	for _, mult := range mults {
		total += mult
	}
	return total
}
