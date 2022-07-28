package main

import "fmt"

/*
	size.channel = len(function)
	number of function
	penghitungan summary dilakukan secara paralel
*/
func fillChannel(functions ...func() int) chan int {
	funcCh := make(chan int, len(functions))
	for _, fu := range functions {
		go fu()
		funcCh <- fu()
	}
	return funcCh
}

func summary(counter int) int {
	sum := 0
	for i := 0; i < counter; i++ {
		sum += 1
	}
	return sum
}

func main() {
	expensiveFunction := func() int { return summary(200) }
	cheapFunction := func() int { return summary(100) }

	ch := fillChannel(expensiveFunction, cheapFunction)

	if ch != nil {
		for i := 0; i < 2; i++ {
			fmt.Printf("Result: %d\n", <-ch)
		}
	}
}
