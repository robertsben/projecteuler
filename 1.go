package main

import (
    "sync"
)

func getMultiples(multiple, limit int) <-chan int {
    channel := make(chan int)
    go func() {
        for i := 0; i < limit; i += multiple {
            channel <- i
        }
        close(channel)
    }()
    return channel
}

func fanInChannels(channels ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    for _, c := range channels {
        go func(x <-chan int) {
            for incoming := range x {
                out <- incoming
            }
            wg.Done()
        }(c)
    }

    wg.Add(len(channels))

    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}


func sumMultiples(multipleA, multipleB, limit int) int {
    multiplesChannel := fanInChannels(
        getMultiples(multipleA, limit),
        getMultiples(multipleB, limit),
    )
    duplicatesChannel := getMultiples(multipleA*multipleB, limit)
    total := 0

    for multiple := range multiplesChannel {
        total += multiple
    }
    for duplicate := range duplicatesChannel {
        total -= duplicate
    }

    return total
}

/*
	If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
	Find the sum of all the multiples of 3 or 5 below 1000.
*/
func (s Solution) Problem1(multipleA, multipleB, limit int) int {
    return sumMultiples(multipleA, multipleB, limit)
}
