package main

import "fmt"

func main() {

}

func getDuplicateVal() {
	arr := []int{1, 2, 2, 4, 5, 1, 8}

	nums := make(map[int]int)

	for _, v := range arr {
		nums[v]++
	}
	fmt.Println(nums)
}

func groupByVal() {
	arr := []int{1, 2, 2, 4, 5, 1, 8}

	nums := make(map[int]int)

	for _, v := range arr {
		nums[v]++
	}
	fmt.Println(nums)
}
