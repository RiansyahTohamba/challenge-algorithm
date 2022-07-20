package main

import "fmt"

var db map[rune]rune

func main() {
	seedDB()
	str1 := "(())("
	fmt.Println(isBalanced(str1))
}

func isBalanced(str string) bool {
	isBal := true

	stack := make([]rune, 0, len(str))

	for _, val := range str {
		if _, ok := db[val]; ok {
			stack = append(stack, val)
		} else if isNotPair(stack, db, val) {
			isBal = false
			break
		} else if len(stack) == 0 {
			isBal = false
			break
		} else {
			// pop tiap kali dapat pair, i.e ) } ]
			popLength := len(stack) - 1
			stack = stack[:popLength]
		}
	}
	// tidak punya pair {
	if len(stack) != 0 {
		isBal = false
	}
	return isBal
}

// LIFO: yang paling akhir yg dicek pasangannya
func isNotPair(stack []rune, db map[rune]rune, key rune) bool {
	top := len(stack) - 1
	valFromStack := stack[top]
	keyFromDB := db[valFromStack]
	return keyFromDB != key
}
func seedDB() {
	db = map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}
}
