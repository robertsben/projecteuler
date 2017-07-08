package main

import "testing"

func testProblem2(t *testing.T) {
    actual := sol.Problem2(4000000)
    expected := 4613732
    if actual != expected {
        t.Error("For", 4000000, "expected", expected, "got", actual)
    }
}
