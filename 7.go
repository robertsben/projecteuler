package main

import (
	"math"
)

/** https://primes.utm.edu/howmany.html **/
func findMax(n int) uint32 {
	x := float64(n)
	return uint32(x * (math.Log(x) + math.Log(math.Log(x - 1)) + 1.8 * math.Log(math.Log(x)) / math.Log(x)))
}

func findNthPrime(n int) int {
	max := findMax(n)
	primes := AtkinSieve(max)

	return int(primes[n-1])
}

/**
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
 */
func (s Solution) Problem7(n int) int {
	return findNthPrime(n)
}