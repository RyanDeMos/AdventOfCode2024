package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	example_input := read_input("./day14/part1/inputs/example_input.txt")
	example_total := part1(example_input)
	log.Printf("Example total: %v\n", example_total)

	input := read_input("./day14/part1/inputs/input.txt")
	total := part1(input)
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

func part1(lines [][]string) int {
	return 0
}
