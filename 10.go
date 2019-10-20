package main

func getSumOfPrimesUnder(n int) uint64 {
	primes := AtkinSieve(uint32(n))
	var total uint64

	for _, prime := range primes {
		total += uint64(prime)
	}
	return total
}
/**
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
 */
func (s Solution) Problem10(n int) int {
	return int(getSumOfPrimesUnder(n))
}