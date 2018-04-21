package main

import (
	"math"
	"fmt"
)

func atkinSieve(max uint32) []uint32 {
	var x, y, n uint32
	maxSqrt := math.Sqrt(float64(max))

	isPrime := make([]bool, max+1, max+1)

	for x = 1; float64(x) <= maxSqrt; x++ {
		for y = 1; float64(y) <= maxSqrt; y++ {
			n = (4*x*x) + (y*y)
			if n <= max && (n % 12 == 1 || n %12 == 5) {
				isPrime[n] = !isPrime[n]
			}
			n = (3*x*x) + (y*y)
			if n <= max && n % 12 == 7 {
				isPrime[n] = !isPrime[n]
			}
			n = (3*x*x) - (y*y)
			if x > y && n <= max && n % 12 == 11 {
				isPrime[n] = !isPrime[n]
			}
		}
	}

	for n = 5; float64(n) <= maxSqrt; n++ {
		if isPrime[n] {
			for y = n*n; y < max; y += n*n {
				isPrime[y] = false
			}
		}
	}

	isPrime[2] = true
	isPrime[3] = true

	var primes []uint32

	for x = 0; int(x) < len(isPrime)-1; x++ {
		if isPrime[x] {
			primes = append(primes, x)
		}
	}

	return primes
}

/** https://primes.utm.edu/howmany.html **/
func findMax(n int) uint32 {
	x := float64(n)
	return uint32(x * (math.Log(x) + math.Log(math.Log(x - 1)) + 1.8 * math.Log(math.Log(x)) / math.Log(x)))
}

func findNthPrime(n int) int {
	max := findMax(n)
	primes := atkinSieve(max)

	return int(primes[n-1])
}

/**
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
 */
func (s Solution) Problem7(n int) int {
	return findNthPrime(n)
}