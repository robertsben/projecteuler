package main

import "testing"

func TestProblem7ExampleSolution(t *testing.T) {
	input := 6
	expected := 13
	solution := sol.Problem7(input)
	if solution != int(expected) {
		t.Errorf("For %v expected %v, got %v", input, expected, solution)
	}
}

func TestProblem7ActualSolution(t *testing.T) {
	input := 10001
	expected := 104743
	solution := sol.Problem7(input)
	if solution != int(expected) {
		t.Errorf("For %v expected %v, got %v", input, expected, solution)
	}
}