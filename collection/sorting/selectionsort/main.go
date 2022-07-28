package main

import "fmt"

func selectionSort(numbers []int) []int {
	iterator := 1
	// remove duplicate here
	// selama iterator dibawah len(numbers)

	// kenapa harus iterator dibawah len(numbers)-1?
	stop := len(numbers) - 1

	for iterator < stop {
		var iterator2 = iterator + 1
		var minIndex = iterator

		// jika iterator2 lebih kecil dari len(numbers)
		if iterator2 < len(numbers) {
			// compare min index with
			if numbers[iterator2] < numbers[minIndex] {
				minIndex = iterator2
			}
			iterator2++
		}
		// disini nilai iterator2 terus bertambah
		// iterator 2 berfungsi sebagai apa?
		// iterator 2 mencari pasangan yang akan dicompare

		// jika minIndex tidak
		// temukan dulu min_index nya baru apa?
		if minIndex != iterator {
			// swap
			var temp = numbers[iterator]
			numbers[iterator] = numbers[minIndex]
			numbers[minIndex] = temp
			// temp diambil dari number iterator
			// iterator = 1
		}

		iterator++
	}
	return numbers
}

func main() {
	list := []int{20, 9, 7, 2, 10}

	fmt.Println(selectionSort(list))
}
