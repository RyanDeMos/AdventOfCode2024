package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	example_input := read_input("./day07/part2/inputs/example_input.txt")
	example_total := part1(example_input)
	log.Printf("Example total: %v\n", example_total)

	input := read_input("./day07/part2/inputs/input.txt")
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
	total_valid := int64(0)
	for _, line := range lines {
		target, numbers := parse_input(line)
		if is_valid_equation(target, numbers, int64(0), 0) {
			total_valid += target
		}
	}
	return total_valid
}

func parse_input(line string) (int64, []int64) {
	key_val := strings.Split(line, ":")
	key, err := strconv.ParseInt(key_val[0], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	values := []int64{}
	for _, str := range strings.Fields(key_val[1]) {
		val, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, val)
	}
	return key, values
}

func is_valid_equation(desired_result int64, numbers []int64, current_result int64, index int) bool {
	if index == len(numbers) || current_result > desired_result {
		return current_result == desired_result
	}

	return is_valid_equation(desired_result, numbers, current_result+numbers[index], index+1) || is_valid_equation(desired_result, numbers, current_result*numbers[index], index+1) || is_valid_equation(desired_result, numbers, concatenate_ints(current_result, numbers[index]), index+1)
}

func concatenate_ints(x, y int64) int64 {
	concatenation_as_string := strconv.FormatInt(x, 10) + strconv.FormatInt(y, 10)
	concatenation, err := strconv.ParseInt(concatenation_as_string, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return concatenation
}
