package main

import "fmt"

func main() {
	list := []int{20, 9, 7, 2, 30}
	high1, high2 := getFirstSecondRank(list)
	fmt.Println(high1)
	fmt.Println(high2)
}

// O(n)
func getFirstSecondRank(nums []int) (high1, high2 int) {
	for _, num := range nums {
		if num > high1 {
			high2 = high1
			high1 = num
		} else if num != high1 && num > high2 {
			high2 = num
		}
	}
	return
}
