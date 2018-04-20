package main

import "testing"

func TestProblem1(t *testing.T) {
    actual := sol.Problem1(3, 5, 1000)
    expected := 233168
    if actual != expected {
        t.Error("For", 3, 5, 1000, "expected", expected, "got", actual)
    }
}

func TestProblem1ExampleSolution(t *testing.T) {
    expected := 23
    actual := sol.Problem1(3, 5, 10)
    if actual != expected {
        t.Errorf("For [%v, %v] expected %v, got %v", 3, 5, expected, actual)
    }
}
