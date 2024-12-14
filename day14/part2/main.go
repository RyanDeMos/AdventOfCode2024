package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	example_input := read_input("./day13/part2/inputs/example_input.txt")
	example_total := part2(example_input)
	log.Printf("Example total: %v\n", example_total)

	input := read_input("./day13/part2/inputs/input.txt")
	total := part2(input)
	log.Printf("Total: %v\n", total)
}

func read_input(inputFile string) [][]string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := [][]string{}
	for scanner.Scan() {
		lines = append(lines, strings.Split(string(scanner.Bytes()), ""))
	}
	return lines
}

func part2(lines [][]string) int {
	return 0
}
