package main

import "fmt"

func main() {
	list := []int{20, 9, 7, 2, 30}
	firstRank, secondRank := getFirstSecondRank(list)
	fmt.Println(firstRank)
	fmt.Println(secondRank)

}

// O(n)
func getFirstSecondRank(nums []int) (firstRank, secondRank int) {
	for _, num := range nums {
		// is new ranking 1?
		if num > firstRank {
			// ranking 2 diganti oleh ranking 1 sebelumnya
			secondRank = firstRank
			// ranking 1 digeser oleh elemen pada array
			firstRank = num
		} else if num != firstRank && num > secondRank {
			// jika num=firstRank, maka ranking 2 sama dengan ranking 1
			// ranking 2 diganti oleh bilangan num besar selain firstRank
			secondRank = num
		}
	}
	return
}

// pada algoritma ini, kita tidak perlu mengurutkan semua bilangan hingga terurut sempurna
// jadi tidak perlu sorting untuk mencari rank 1 dan rank 2
// sorting is overkill

// if - if lebih lambat dibanding if - elseif
