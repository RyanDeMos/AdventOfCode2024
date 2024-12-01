package main

import "testing"

func Test_part2(t *testing.T) {
	expectedSimularityScore := 31
	actualTotalDistance := part2("inputs/example_input.txt")
	if expectedSimularityScore != actualTotalDistance {
		t.Fatalf("Expected 31 but got: %v\n", actualTotalDistance)
	}
}
