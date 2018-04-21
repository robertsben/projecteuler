package main

import "testing"

func TestProblem6ExampleSolution(t *testing.T) {
	input := 10
	expected := 2640
	solution := sol.Problem6(input)
	if solution != int64(expected) {
		t.Errorf("For %v expected %v, got %v", input, expected, solution)
	}
}

func TestProblem6ActualSolution(t *testing.T) {
	input := 100
	expected := 25164150
	solution := sol.Problem6(input)
	if solution != int64(expected) {
		t.Errorf("For %v expected %v, got %v", input, expected, solution)
	}
}