package main

import (
	"fmt"
	"sync"
)

func getMultiples(mul, limit int, wg *sync.WaitGroup) <-chan int {
	c := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < limit; i += mul {
			c <- i
		}
	}()
	return c
}

func fanInMultipliers(multia, multib <-chan int) <-chan int {
	c := make(chan int)
	go func() { for { c <- <-multia } }()
	go func() { for { c <- <-multib } }()
	return c
}

func sumMultiples(mula, mulb, limit int) int {
	var total int
	var wg sync.WaitGroup
	c := fanInMultipliers(
		getMultiples(mula, limit, &wg),
		getMultiples(mulb, limit, &wg),
	)
	d := getMultiples(mula * mulb, limit, &wg)

	go func() {
		for {
			select {
			case x := <-c:
				total += x
			case y := <-d:
				total -= y
			}
		}
	}()

	wg.Wait()

	fmt.Println(total)

	return total
}

/*
	If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
	Find the sum of all the multiples of 3 or 5 below 1000.
*/
func (s Solution) Problem1(mula, mulb, limit int) int {
	val := sumMultiples(mula, mulb, limit)
	fmt.Println(val)
	return val
}