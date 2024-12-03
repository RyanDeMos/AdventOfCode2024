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
	example_input := read_input("./day03/part2/inputs/example_input.txt")
	example_total := part2(example_input)
	log.Printf("Example total: %v\n", example_total)

	input := read_input("./day03/part2/inputs/input.txt")
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
	single_line := convert_input_to_single_line(lines)
	lines_after_removal := remove_donts(single_line)
	mults := mults(lines_after_removal)
	total := find_total(mults)
	return total
}

func mults(line string) []int {
	r, err := regexp.Compile(`mul\((\d\d?\d?,\d\d?\d?)\)`)
	if err != nil {
		log.Fatal(err)
	}

	mults := []int{}

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
	return mults
}

func find_total(mults []int) int {
	total := 0
	for _, mult := range mults {
		total += mult
	}
	return total
}

func convert_input_to_single_line(lines []string) string {
	single_line := ""
	for _, line := range lines {
		single_line += line
	}
	return single_line
}

func remove_donts(line string) string {
	donts, err := regexp.Compile(`don't\(\)`)
	if err != nil {
		log.Fatal(err)
	}

	dos, err := regexp.Compile(`do\(\)`)
	if err != nil {
		log.Fatal(err)
	}

	match_all_donts := donts.FindAllStringSubmatchIndex(line, -1)
	match_all_dos := dos.FindAllStringSubmatchIndex(line, -1)

	substrings_to_remove := []string{}
	for dont_index, dont := range match_all_donts {
		remove_start := dont[0]
		remove_end := len(line)
		for _, do := range match_all_dos {
			if do[0] > dont[0] {
				remove_end = do[0]
				if ((dont_index + 1) < len(match_all_donts)) && (remove_end > match_all_donts[dont_index+1][0]) {
					remove_end = match_all_donts[dont_index+1][0]
				}
				break
			}
		}
		substrings_to_remove = append(substrings_to_remove, line[remove_start:remove_end])
	}

	for _, substring := range substrings_to_remove {
		line = strings.ReplaceAll(line, substring, "")
	}

	return line
}
