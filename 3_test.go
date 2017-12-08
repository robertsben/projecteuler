package main

import "testing"

func TestProblem3(t *testing.T) {
    input := 600851475143
    actual := sol.Problem3(input)
    expected := 6857
    if actual != expected {
        t.Error("For", input, "expected", expected, "got", actual)
    }
}
