package main

import "math"

// Function for generating all the prime numbers below max
func AtkinSieve(max uint32) []uint32 {
	var x, y, n uint32
	maxSqrt := math.Sqrt(float64(max))

	isPrime := make([]bool, max+1, max+1)

	for x = 1; float64(x) <= maxSqrt; x++ {
		for y = 1; float64(y) <= maxSqrt; y++ {
			n = (4 * x * x) + (y * y)
			if n <= max && (n%12 == 1 || n%12 == 5) {
				isPrime[n] = !isPrime[n]
			}
			n = (3 * x * x) + (y * y)
			if n <= max && n%12 == 7 {
				isPrime[n] = !isPrime[n]
			}
			n = (3 * x * x) - (y * y)
			if x > y && n <= max && n%12 == 11 {
				isPrime[n] = !isPrime[n]
			}
		}
	}

	for n = 5; float64(n) <= maxSqrt; n++ {
		if isPrime[n] {
			for y = n * n; y < max; y += n * n {
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

