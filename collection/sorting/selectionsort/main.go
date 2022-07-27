package main

import "fmt"

func findMaxSum(numbers []int) int {
	sorted := selectionSort(numbers)
	size := len(numbers)

	return sorted[size-1] + sorted[size-2]
}

func selectionSort(n []int) []int {
	i := 1
	// remove duplicate here
	//
	for i < len(n)-1 {
		var j = i + 1
		var minIndex = i

		if j < len(n) {
			if n[j] < n[minIndex] {
				minIndex = j
			}
			j++
		}

		if minIndex != i {
			var temp = n[i]
			n[i] = n[minIndex]
			n[minIndex] = temp
		}

		i++
	}
	return n
}

func main() {
	list := []int{5, 9, 7, 11, 10, 20, 30}
	fmt.Println(findMaxSum(list))
}
