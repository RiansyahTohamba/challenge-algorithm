package main

import (
	"fmt"
)

func old() {
	var password string = "#abcd$1AK"

	ascii := make([]int, 128)
	for _, ch := range password {
		ascii[ch]++
		// fmt.Println(ch)
	}
	for idx, ch := range ascii {
		if ch > 0 {
			// bagaimana cara cek special character?

			fmt.Println(string(idx))
		}
	}

}
func main() {
	contoh := 'r'
	fmt.Println(contoh)
	// var password string = "[abcdAK{"
	// // var special_characters string = "!@#$%^&*()-+"
	// // pswdmap := make(map[rune])
	// // var remain int
	// letter := func(r rune) bool {
	// 	// 65-90=UPPERCASE
	// 	// 97-122=lowercase
	// 	// r < 'A' = r < 65
	// 	// atau
	// 	// r > 'z' = r > 122
	// 	// diantara 91-96 = ini adalah apa?
	// 	// 91 = [
	// 	// maka pengecekan ini masih salah
	// 	return r < 'A' || r > 'z'
	// }
	// if strings.IndexFunc(password, letter) != -1 {
	// 	fmt.Println("Found special char")
	// }

	// # = 35
	// jika return indexFunc = -1 berarti special character ketemu
	// jika return indexFunc bukan -1, berarti special character tidak ketemu
}
