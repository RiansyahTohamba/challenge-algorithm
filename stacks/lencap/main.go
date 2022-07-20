package main

import "fmt"

func main() {
	// len
	// cap
	//	Slice: The size specifies the length. The capacity of the slice is
	//	equal to its length.
	// A second integer argument may be provided to
	//	specify a different capacity; it must be no smaller than the
	//	length.
	// For example, make([]int, 0, 10) allocates an underlying array
	//	of size 10 and returns a slice of length 0 and capacity 10 that is
	//	backed by this underlying array.

	stack := make([]int, 0)
	// push
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 4)
	fmt.Println(stack)

	fmt.Println(stack[0:4])
	// pop
	// stack = stack[1:len(stack)]

}

func isBalanced() bool {
	isBal := true
	return isBal
}
