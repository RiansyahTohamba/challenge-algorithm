package main

import (
	"fmt"
	"math"
)

func plusMinus(arr []int32) {
	// Write your code here
	length := len(arr)
	numPositive := 0
	numNegative := 0
	numZero := 0

	for _, v := range arr {
		if v > 0 {
			numPositive++
		} else if v < 0 {
			numNegative++
		} else {
			numZero++
		}
	}

	ratioPs := float32(numPositive) / float32(length)
	fmt.Println(ratioPs)

	ratioNg := float32(numNegative) / float32(length)
	fmt.Println(ratioNg)

	ratioZero := float32(numZero) / float32(length)
	fmt.Println(ratioZero)
}

func main() {
	// arr := []int32{0, -1, 2, 3, -4, 4}
	// plusMinus(arr)
	fmt.Println(math.Floor(1 / 3))
}
