package main

func isPalindrome(word string) bool {
	// check length
	isPalindrome := true

	stack := make([]rune, 0, len(word))
	var newWord []rune

	// push
	for _, ch := range word {
		stack = append(stack, ch)
	}
	// pop
	for len(stack) > 0 {
		top := len(stack) - 1
		newWord = append(newWord, stack[top])
		stack = stack[:top]
	}
	if word != string(newWord) {
		isPalindrome = false
	}
	return isPalindrome
}

/*
bagaimana menyelesaikan palindrome?
1. push ch ke stack
2. pop sampai len(stack) == 0
3.
*/
