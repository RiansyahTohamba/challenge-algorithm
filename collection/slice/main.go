package main

import "fmt"

func main() {
	// beda antara nil vs 0
	// ketika len(things) = 2, cap(things) = 2 maka array sudah terisi dengan nilai 0 sebanyak 2 tempat
	// ketika len(things) = 0, cap(things) = 2 maka array sudah terisi dengan nilai 0 sebanyak 2 tempat
	// cap > len tidak mungkin len > cap
	things := make([]int, 3, 20)
	// cap = 20, len 3, maka hanya array 0 - 2 yang sudah terisi 0
	// 3 - 19 belum diisi
	fmt.Println(cap(things))
	fmt.Println(len(things))
	fmt.Println(things[0])
	fmt.Println(things[4])

}

func testlength() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[:len(nums)-1])
	fmt.Println(nums[:4])
	// num ke-5 nya hilang

	fmt.Println(nums[:len(nums)-2])
	fmt.Println(nums[:3])
	// num ke-4 dan ke-5 nya hilang

	fmt.Println(nums[:len(nums)-3])
	fmt.Println(nums[:len(nums)-4])

	// kurangi dari kanan
}
