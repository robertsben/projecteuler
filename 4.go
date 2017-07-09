package main

import (
    "math"
    "fmt"
)

func fillRange(min, max int) []int {
    rangeContents := make([]int, max-min)
    for i := min; i < max; i++ {
        rangeContents[i-min] = i
    }
    return rangeContents
}

func generateRangeChunks(chunkRangeLimits []int) []interface{} {
    var rangeChunks []interface{}

    for i := 0; i < len(chunkRangeLimits)-1; i++ {
        rangeChunks = append(rangeChunks, fillRange(chunkRangeLimits[i], chunkRangeLimits[i+1]))
    }

    return rangeChunks
}

func splitRange(min, max int) []int {
    lowerHalfMax := int(math.Floor(float64((max+min)/2)))
    return []int{min, lowerHalfMax, max}
}

func addValesToSlice(slice []interface{}, values []interface{}) []interface{} {
    for _, val := range values {
        slice = append(slice, val)
    }
    return slice
}

func rangeChunker(min, max int) []interface{} {
    var ranges []interface{}
    chunkRangeLimits := splitRange(min, max)

    if max - min <= 25 {
        ranges = addValesToSlice(ranges, generateRangeChunks(chunkRangeLimits))
    } else {
        for i := 0; i < 2; i++ {
            ranges = addValesToSlice(ranges, rangeChunker(chunkRangeLimits[i], chunkRangeLimits[i+1]))
        }
    }
    return ranges
}

/*
	A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

    Find the largest palindrome made from the product of two 3-digit numbers.
*/
func (s Solution) Problem4(digitCountA, digitCountB int) int {
    minValA := int(math.Pow10(digitCountA-1))
    maxValA := int(math.Pow10(digitCountA) - 1)
    minValB := int(math.Pow10(digitCountB-1))
    //maxValB := 10^digitCountB - 1
    fmt.Println(rangeChunker(minValA, maxValA+1))
    return minValB
}

