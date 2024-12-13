package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	example_input := read_input("./day13/part1/inputs/example_input.txt")
	example_total := part1(example_input)
	log.Printf("Example total: %v\n", example_total)

	input := read_input("./day13/part1/inputs/input.txt")
	total := part1(input)
	log.Printf("Total: %v\n", total)
}

func read_input(inputFile string) [][2][3]float64 {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	matrices_to_solve := [][2][3]float64{}
	current_matrix := [2][3]float64{}
	for scanner.Scan() {
		line := string(scanner.Bytes())
		if line == "" {
			matrices_to_solve = append(matrices_to_solve, current_matrix)
			current_matrix = [2][3]float64{}
		} else if strings.Contains(line, "Button A") {
			split_string := strings.Split(line, ",")
			a_x, _ := strconv.Atoi(strings.Split(split_string[0], "+")[1])
			a_y, _ := strconv.Atoi(strings.Split(split_string[1], "+")[1])
			current_matrix[0][0] = float64(a_x)
			current_matrix[1][0] = float64(a_y)
		} else if strings.Contains(line, "Button B") {
			split_string := strings.Split(line, ",")
			b_x, _ := strconv.Atoi(strings.Split(split_string[0], "+")[1])
			b_y, _ := strconv.Atoi(strings.Split(split_string[1], "+")[1])
			current_matrix[0][1] = float64(b_x)
			current_matrix[1][1] = float64(b_y)
		} else if strings.Contains(line, "Prize") {
			split_string := strings.Split(line, ",")
			prize_x, _ := strconv.Atoi(strings.Split(split_string[0], "=")[1])
			prize_y, _ := strconv.Atoi(strings.Split(split_string[1], "=")[1])
			current_matrix[0][2] = float64(prize_x)
			current_matrix[1][2] = float64(prize_y)
		}
	}
	matrices_to_solve = append(matrices_to_solve, current_matrix)
	// log.Printf("matrices_to_solve: %v\n", matrices_to_solve)
	return matrices_to_solve
}

func part1(matrices [][2][3]float64) int {
	total_req := 0
	for _, matrix := range matrices {
		result := Gauss_Jordian_elim(matrix)
		// log.Printf("Result: %v\n", result)
		tokens_for_this_matrix := count_tokens(result)
		// log.Printf("Tokens Spent: %v\n", tokens_for_this_matrix)
		total_req += tokens_for_this_matrix
	}
	return total_req
}

func Gauss_Jordian_elim(matrix [2][3]float64) [2][3]float64 {
	mult_factor := matrix[0][0] / matrix[1][0]
	for i := 0; i < 3; i++ {
		matrix[1][i] *= mult_factor
		matrix[1][i] -= matrix[0][i]
	}

	divide_factor := matrix[1][1]
	for i := 0; i < 3; i++ {
		matrix[1][i] /= divide_factor
		if math.Abs(matrix[1][i]-math.Round(matrix[1][i])) < math.Pow10(-5) {
			matrix[1][i] = math.Round(matrix[1][i])
		}
	}

	mult_factor = matrix[0][1]
	for i := 0; i < 3; i++ {
		matrix[0][i] -= mult_factor * matrix[1][i]
	}

	divide_factor = matrix[0][0]
	for i := 0; i < 3; i++ {
		matrix[0][i] /= divide_factor
		if math.Abs(matrix[0][i]-math.Round(matrix[0][i])) < math.Pow10(-5) {
			matrix[0][i] = math.Round(matrix[0][i])
		}
	}

	return matrix
}

func count_tokens(matrix [2][3]float64) int {
	tokens_req := 0
	if matrix[0][2] == math.Round(matrix[0][2]) && matrix[1][2] == math.Round(matrix[1][2]) {
		tokens_req += 3*int(matrix[0][2]) + int(matrix[1][2])
	}
	return tokens_req
}
