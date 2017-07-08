package main

import "testing"

func testProblem1(t *testing.T) {
    actual := sol.Problem1(3, 5, 1000)
    expected := 233168
    if actual != expected {
        t.Error("For", 3, 5, 1000, "expected", expected, "got", actual)
    }
}
