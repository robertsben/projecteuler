package main

import (
    "math"
    "sync"
)

func fillRange(min, max int) []int {
    rangeContents := make([]int, max-min)
    for i := min; i < max; i++ {
        rangeContents[i-min] = i
    }
    return rangeContents
}

func generateRangeChunks(chunkRangeLimits []int) []int {
    var rangeChunks []int

    for i := 0; i < len(chunkRangeLimits)-1; i++ {
        rangeChunks = addValesToSlice(rangeChunks, fillRange(chunkRangeLimits[i], chunkRangeLimits[i+1]))
    }

    return rangeChunks
}

func splitRange(min, max int) []int {
    lowerHalfMax := int(math.Floor(float64((max+min)/2)))
    return []int{min, lowerHalfMax, max}
}

func addValesToSlice(slice []int, values []int) []int {
    for _, val := range values {
        slice = append(slice, val)
    }
    return slice
}

/*
    Given the min and max number in a range, fire sets of 
 */
func rangeChunker(rangeChannel chan []int, min, max int, wg *sync.WaitGroup) {
    defer wg.Done()

    chunkRangeLimits := splitRange(min, max)

    if max - min <= 25 {
        rangeChannel <- generateRangeChunks(chunkRangeLimits)
        return
    } else {
        for i := 0; i < 2; i++ {
            wg.Add(1)
            go rangeChunker(rangeChannel, chunkRangeLimits[i], chunkRangeLimits[i+1], wg)
        }
    }
}

/*
    Split the range between min and max into chunks
 */
func splitRangeToChunks(min, max int) [][]int {
    rangeChannel := make(chan []int)
    var wg sync.WaitGroup
    var rangeChunks [][]int
    go func() {
        for {
            select {
            case rangeChunk := <-rangeChannel:
                rangeChunks = append(rangeChunks, rangeChunk)
            }
        }
    }()
    wg.Add(1)
    go rangeChunker(rangeChannel, min, max+1, &wg)
    wg.Wait()
    return rangeChunks
}

/*
    Get max and min of n-digit numbers, where n = length
 */
func getNumberRangeOfDigitLength(length int) (min, max int) {
    return int(math.Pow10(length-1)), int(math.Pow10(length) - 1)
}

/* e.g. 123456 -> 654321 */
func reverseNumber(number int) int {
    reversed := 0
    for number > 0 {
        remainder := number % 10
        reversed *= 10
        reversed += remainder
        number /= 10
    }
    return reversed
}

func isPalindrome(number int) bool {
    return number == reverseNumber(number)
}

func calculateAllMultiples(multipleChannel chan int, rangeA, rangeB []int, completeWg *sync.WaitGroup) {
    defer completeWg.Done()
    for _, factorA := range rangeA {
        for _, factorB := range rangeB {
            multipleChannel <- factorA * factorB
        }
    }
}

func multiplyRangesBy(multipleRange []int, ranges [][]int, multipleChannel chan int, completeWg *sync.WaitGroup) {
    defer completeWg.Done()
    var wg sync.WaitGroup
    for _, numRange := range ranges {
        wg.Add(1)
        go calculateAllMultiples(multipleChannel, multipleRange, numRange, &wg)
    }
    wg.Wait()
}

func findPalindromicMultiples(rangesA, rangesB [][]int) int {
    multipleChannel := make(chan int)
    var wg sync.WaitGroup
    maximumPalindrome := 0
    go func() {
        for {
            select {
            case multiple := <-multipleChannel:
                if multiple > maximumPalindrome && isPalindrome(multiple) {
                    maximumPalindrome = multiple
                }
            }
        }
    }()
    for _, numRangeA := range rangesA {
        wg.Add(1)
        go multiplyRangesBy(numRangeA, rangesB, multipleChannel, &wg)
    }
    wg.Wait()
    return maximumPalindrome
}

/*
	A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

    Find the largest palindrome made from the product of two 3-digit numbers.
*/
func (s Solution) Problem4(digitCountA, digitCountB int) int {
    minValA, maxValA := getNumberRangeOfDigitLength(digitCountA)
    minValB, maxValB := getNumberRangeOfDigitLength(digitCountB)
    aDigitNumberRangeChunks := splitRangeToChunks(minValA, maxValA)
    bDigitNumberRangeChunks := splitRangeToChunks(minValB, maxValB)
    maxPalindrome := findPalindromicMultiples(aDigitNumberRangeChunks, bDigitNumberRangeChunks)

    return maxPalindrome
}

