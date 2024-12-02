package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reports_example := create_reports("./day02/part2/inputs/example_input.txt")
	total_example := find_total_safe(reports_example)
	log.Printf("Total Example Safe: %v\n", total_example)

	reports := create_reports("./day02/part2/inputs/input.txt")
	total := find_total_safe(reports)
	log.Printf("Total Safe: %v\n", total)
}

func create_reports(inputFile string) [][]int {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	reports := [][]int{}
	for scanner.Scan() {
		report_string := string(scanner.Bytes())
		reports_string := strings.Fields(report_string)
		report := []int{}
		for i := 0; i < len(reports_string); i++ {
			as_int, err := strconv.Atoi(reports_string[i])
			if err != nil {
				log.Fatal(err)
			}
			report = append(report, as_int)
		}
		reports = append(reports, report)
	}
	return reports
}

func find_total_safe(reports [][]int) int {
	total := 0
	for i := 0; i < len(reports); i++ {
		if is_safe_part2(reports[i]) {
			total += 1
		}
	}
	return total
}

func is_safe(report []int) bool {
	ascending_report := slices.Clone(report)
	descending_report := slices.Clone(report)
	is_ascending := is_safe_ascending(ascending_report)
	is_descending := is_safe_descending(descending_report)
	return is_ascending || is_descending
}

func is_safe_descending(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		if (report[i] < report[i+1]) || (math.Abs(float64(report[i]-report[i+1])) < 1) || (math.Abs(float64(report[i]-report[i+1])) > 3) {
			return false
		}
	}
	return true
}

func is_safe_ascending(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		if (report[i] > report[i+1]) || (math.Abs(float64(report[i]-report[i+1])) < 1) || (math.Abs(float64(report[i]-report[i+1])) > 3) {
			return false
		}
	}
	return true
}

// Dumbest way possible
// Remove each element one at a time to see if any are valid
func is_safe_part2(report []int) bool {
	if is_safe(report) {
		return true
	}
	for i := 0; i < len(report); i++ {
		new_slice := slices.Clone(report)
		new_slice = append(new_slice[:i], new_slice[i+1:]...)
		if is_safe(new_slice) {
			return true
		}
	}
	return false
}
