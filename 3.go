package main

import (
	"math"
)

func findFactors(c, quit chan int, number int) {
	for i := 2; i < int(math.Sqrt(float64(number))); i++ {
		if 0 == math.Mod(float64(number), float64(i)) {
			c <- i
			c <- number/i
		}
	}
	quit <- 0
}

func isPrime(number int) bool {
	c := make(chan int)
	quit := make(chan int)
	isPrime := true

	go func() {
		for {
			select {
			case <-c:
				isPrime = false
			case <-quit:
				return
			}
		}
	}()

	findFactors(c, quit, number)

	return isPrime
}

func findMaxPrimeFactor(number int) int {
	c := make(chan int)
	quit := make(chan int)
	max := 1
	go func() {
		for {
			select {
			case x := <-c:
				if x > max && isPrime(x) {
					max = x
				}
			case <-quit:
				return
			}
		}
	}()
	findFactors(c, quit, number)
	return max
}


/*
	The prime factors of 13195 are 5, 7, 13 and 29.

	What is the largest prime factor of the number 600851475143 ?
*/
func (s Solution) Problem3(number int) int {
	return findMaxPrimeFactor(number)
}