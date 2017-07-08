package main

import (
    "math"
)

func fibonacci(channel, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case channel <- x:
            x, y = y, x+y
        case <-quit:
            return
        }
    }
}

func fibber(limit int) int {
    channel := make(chan int)
    quit := make(chan int)
    sum := 0
    go func() {
        for fibElem := range channel {
            if fibElem > limit {
                quit <- 0
            }
            if 0 == math.Mod(float64(fibElem), 2) {
                sum += fibElem
            }
        }
    }()
    fibonacci(channel, quit)
    return sum
}

/*
	Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

	1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

	By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/
func (s Solution) Problem2(limit int) int {
    return fibber(limit)
}
