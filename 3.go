package main

import (
    "math"
)

func findFactors(channel, quit chan int, number int) {
    for i := 2; i < int(math.Sqrt(float64(number))); i++ {
        if 0 == math.Mod(float64(number), float64(i)) {
            channel <- i
            channel <- number / i
        }
    }
    quit <- 0
}

func isPrime(number int) bool {
    factorChannel := make(chan int)
    quitChannel := make(chan int)
    isPrime := true

    go func() {
        for {
            select {
            case <-factorChannel:
                isPrime = false
            case <-quitChannel:
                return
            }
        }
    }()

    findFactors(factorChannel, quitChannel, number)

    return isPrime
}

func findMaxPrimeFactor(number int) int {
    channel := make(chan int)
    quit := make(chan int)
    max := 1
    go func() {
        for {
            select {
            case factor := <-channel:
                if factor > max && isPrime(factor) {
                    max = factor
                }
            case <-quit:
                return
            }
        }
    }()
    findFactors(channel, quit, number)
    return max
}

/*
	The prime factors of 13195 are 5, 7, 13 and 29.

	What is the largest prime factor of the number 600851475143 ?
*/
func (s Solution) Problem3(number int) int {
    return findMaxPrimeFactor(number)
}
