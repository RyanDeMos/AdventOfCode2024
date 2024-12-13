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
			current_matrix[0][2] = float64(prize_x) + float64(10000000000000)
			current_matrix[1][2] = float64(prize_y) + float64(10000000000000)
		}
	}
	matrices_to_solve = append(matrices_to_solve, current_matrix)
	return matrices_to_solve
}

func part1(matrices [][2][3]float64) float64 {
	total_req := float64(0)
	for _, matrix := range matrices {
		result := CramersRule(matrix)
		total_req += result
	}
	return total_req
}

func CramersRule(matrix [2][3]float64) float64 {
	A := [][]float64{{0, 0}, {0, 0}}
	for row_idx, row := range matrix {
		A[row_idx][0] = row[0]
		A[row_idx][1] = row[1]
	}
	det_A := A[0][0]*A[1][1] - A[0][1]*A[1][0]
	if det_A != 0 {
		A1 := [][]float64{{A[0][0], A[0][1]}, {A[1][0], A[1][1]}}
		A1[0][0] = matrix[0][2]
		A1[1][0] = matrix[1][2]
		det_A1 := A1[0][0]*A1[1][1] - A1[0][1]*A1[1][0]

		button_a := (det_A1 / det_A)

		A2 := [][]float64{{A[0][0], A[0][1]}, {A[1][0], A[1][1]}}
		A2[0][1] = matrix[0][2]
		A2[1][1] = matrix[1][2]
		det_A2 := A2[0][0]*A2[1][1] - A2[0][1]*A2[1][0]

		button_b := det_A2 / det_A

		if has_integer_solution(button_a, button_b) && button_b >= 0 && button_a >= 0 {
			return 3*button_a + button_b
		}
	}
	return 0
}

func has_integer_solution(a, b float64) bool {
	return math.Abs(math.Round(a)-a) < 0.001 && math.Abs(math.Round(b)-b) < 0.001
}
