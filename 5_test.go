package main

import "testing"

func TestProblem5ExampleSolution(t *testing.T) {
	exampleRange := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 2520
	solution := sol.Problem5(exampleRange)
	if solution != expected {
		t.Errorf("For %v expected %v, got %v", exampleRange, expected, solution)
	}
}
