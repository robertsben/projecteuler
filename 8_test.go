package main

import "testing"

func TestProblem8ExampleSolution(t *testing.T) {
	input := 4
	expected := 5832
	solution := sol.Problem8(input)
	if solution != expected {
		t.Errorf("For %v expected %v, got %v", input, expected, solution)
	}
}

func TestProblem8ActualSolution(t *testing.T) {
	input := 13
	expected := 23514624000
	solution := sol.Problem8(input)
	if solution != expected {
		t.Errorf("For %v expected %v, got %v", input, expected, solution)
	}
}