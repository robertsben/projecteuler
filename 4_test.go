package main

import "testing"

func testProblem4(t *testing.T) {
    inputA, inputB := 3, 3
    actual := sol.Problem4(inputA, inputB)
    expected := 906609
    if actual != expected {
        t.Error("For", inputA, inputB, "expected", expected, "got", actual)
    }
}

func testProblem4anotherInput(t *testing.T) {
    inputA, inputB := 2, 2
    actual := sol.Problem4(inputA, inputB)
    expected := 9009
    if actual != expected {
        t.Error("For", inputA, inputB, "expected", expected, "got", actual)
    }
}
