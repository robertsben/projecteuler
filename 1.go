package main

import (
    "sync"
)

func getMultiples(multiple, limit int, wg *sync.WaitGroup) <-chan int {
    channel := make(chan int)
    wg.Add(1)
    go func() {
        defer wg.Done()
        for i := 0; i < limit; i += multiple {
            channel <- i
        }
    }()
    return channel
}

func fanInMultipliers(multiplierA, multiplierB <-chan int) <-chan int {
    channel := make(chan int)
    go func() { for { channel <- <-multiplierA } }()
    go func() { for { channel <- <-multiplierB } }()
    return channel
}

func sumMultiples(multipleA, multipleB, limit int) int {
    var total int
    var wg sync.WaitGroup
    muliplesChannel := fanInMultipliers(
        getMultiples(multipleA, limit, &wg),
        getMultiples(multipleB, limit, &wg),
    )
    duplicatesChannel := getMultiples(multipleA*multipleB, limit, &wg)

    go func() {
        for {
            select {
            case x := <-muliplesChannel:
                total += x
            case y := <-duplicatesChannel:
                total -= y
            }
        }
    }()

    wg.Wait()

    return total
}

/*
	If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
	Find the sum of all the multiples of 3 or 5 below 1000.
*/
func (s Solution) Problem1(multipleA, multipleB, limit int) int {
    return sumMultiples(multipleA, multipleB, limit)
}
