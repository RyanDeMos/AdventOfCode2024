package main

import "testing"

func Test_part1(t *testing.T) {
	expectedTotalDistance := 11
	actualTotalDistance := part1("inputs/example_input.txt")
	if expectedTotalDistance != actualTotalDistance {
		t.Fatalf("Expected 11 but got: %v\n", actualTotalDistance)
	}
}
