package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[:4])

	// delete 1 element
	nums = nums[:len(nums)-1]
	fmt.Println(nums)
	fmt.Println("================")
	fmt.Println(nums[0:1])
	fmt.Println(nums[1:2])
	fmt.Println(nums[:3])
	fmt.Println(nums[:2])
	fmt.Println(nums[:1])

}
