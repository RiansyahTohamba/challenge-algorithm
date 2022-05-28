package main

import (
	"fmt"
	"reflect"
)

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	return 1
}

func lcm(a, b int32) int32 {
	// a double-slash for “floor” division (rounds down to nearest whole number).
	// syntax       // ini apa?
	x1 := a * b
	x2 := gcd(a, b)
	res := x1 / x2
	return res
}

func gcd(a, b int32) int32 {
	for {
		a, b = b, (a % b)
		if a%b != 0 {
			break
		}
	}
	return b
	// while a % b != 0
	// 	a, b = b, (a % b)
	// return b
}

func evenly_divides(number, divisor int) bool {
	return (number % divisor) == 0
}

// Reduce computes the reduction of the pair function across the elements of
// the slice. (If the types of the slice and function do not correspond, Reduce
// panics.) For instance, if the slice contains successive integers starting at
// 1 and the function is multiply, the result will be the factorial function.
// If the slice is empty, Reduce returns zero; if it has only one element, it
// returns that element. The return value must be type-asserted by the caller
// back to the element type of the slice. Example:
//	func multiply(a, b int) int { return a*b }
//	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	factorial := Reduce(a, multiply, 1).(int)
func Reduce(slice, pairFunction, zero interface{}) interface{} {
	in := reflect.ValueOf(slice)
	if in.Kind() != reflect.Slice {
		panic("reduce: not slice")
	}
	n := in.Len()
	if n == 0 {
		return zero
	}
	elemType := in.Type().Elem()
	fn := reflect.ValueOf(pairFunction)
	if !goodFunc(fn, elemType, elemType, elemType) {
		str := elemType.String()
		panic("apply: function must be of type func(" + str + ", " + str + ") " + str)
	}
	var ins [2]reflect.Value
	out := in.Index(0) // By convention, fn(zero, in[0]) = in[0].
	// Run from index 1 to the end.
	for i := 1; i < n; i++ {
		ins[0] = out
		ins[1] = in.Index(i)
		out = fn.Call(ins[:])[0]
	}
	return out.Interface()
}

func main() {
	arrA := []int32{2, 4}
	arrB := []int32{16, 32, 96}
	filter()
	// # Find LCM of all integers of a
	lcm_value = reduce(gcd, arrA)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := reduce(numbers, func(acc, current int) int {
		return acc + current
	}, 0)


	fmt.Println(sum)

	// # Find GCD of all integers of b
	gcd_value = reduce(gcd, arrB)

	// # Count the number of multiples of LCM that evenly divides the GCD.
	counter := 0
	multiple_of_lcm = lcm_value

	for {
		if evenly_divides(gcd_value, multiple_of_lcm) {
			counter++
		}
		multiple_of_lcm += lcm_value
		if multiple_of_lcm <= gcd_value {
			break
		}
	}
	// while multiple_of_lcm <= gcd_value:
	//     if evenly_divides(gcd_value, multiple_of_lcm):
	//         counter += 1
	//     multiple_of_lcm += lcm_value

	fmt.Println(counter)
		
	int[] a = new int[n];
        for(int a_i=0; a_i < n; a_i++){
            a[a_i] = in.nextInt();
        }
        int[] b = new int[m];
        for(int b_i=0; b_i < m; b_i++){
            b[b_i] = in.nextInt();
        }
        
        int f = lcm(a);
        int l = gcd(b);
        int count = 0;
        
		for(int i = f, j =2; i<=l; i=f*j,j++){
            if(l%i==0){ 
				count++;
			}
        }

        System.out.println(count);
}

private static int gcd(int a, int b) {
	while (b > 0) {
		int temp = b;
		b = a % b; // % is remainder
		a = temp;
	}
	return a;
}

private static int gcd(int[] input) {
	int result = input[0];
	for (int i = 1; i < input.length; i++) {
		result = gcd(result, input[i]);
	}
	return result;
}

private static int lcm(int a, int b) {
	return a * (b / gcd(a, b));
}

private static int lcm(int[] input) {
	int result = input[0];
	for (int i = 1; i < input.length; i++) {
		result = lcm(result, input[i]);
	}
	return result;
}