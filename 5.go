package main

func isDivisible(n, max int) bool {
	for i := max; i > 3; i-- {
		if n % i != 0 { return false }
	}
	return true
}

func searchForDivisible(max int) int {
	for i := max;; i+=max {
		if isDivisible(i, max) { return i }
	}
}

/**
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
 */
func (s Solution) Problem5(max int) int {
	return searchForDivisible(max)
}
