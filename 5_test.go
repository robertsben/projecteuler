package main

import "testing"

func TestProblem5ExampleSolution(t *testing.T) {
	input := 10
	expected := 2520
	solution := sol.Problem5(input)
	if solution != expected {
		t.Errorf("For %v expected %v, got %v", input, expected, solution)
	}
}

func TestProblem5ActualSolution(t *testing.T) {
	input := 20
	expected := 232792560
	solution := sol.Problem5(input)
	if solution != expected {
		t.Errorf("For %v expected %v, got %v", input, expected, solution)
	}
}