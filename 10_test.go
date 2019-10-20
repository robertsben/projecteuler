package main

import "testing"

func TestProblem10ExampleSolution(t *testing.T) {
	input := 10
	expected := 17
	solution := sol.Problem10(input)
	if solution != expected {
		t.Errorf("For %v expected %v, got %v", input, expected, solution)
	}
}

func TestProblem10ActualSolution(t *testing.T) {
	input := 2000000
	expected := 142913828922
	solution := sol.Problem10(input)
	if solution != expected {
		t.Errorf("For %v expected %v, got %v", input, expected, solution)
	}
}