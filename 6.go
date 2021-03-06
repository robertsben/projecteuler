package main

func sumOfSquares(max int64) int64 {
	return (max * (max + 1) * ((2*max) + 1))/6
}

func sum(max int64) int64 {
	return (max * (max + 1))/2
}

/**
The sum of the squares of the first ten natural numbers is,

1^2 + 2^2 + ... + 10^2 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)^2 = 55^2 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
 */
func (s Solution) Problem6(max int) int64 {
	bigmax := int64(max)
	bigsum := sum(bigmax)
	return (bigsum*bigsum) - sumOfSquares(bigmax)
}
