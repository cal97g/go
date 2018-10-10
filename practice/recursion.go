package main

import "fmt"

// 4, 3, 2, 1, 1
func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
	//4 * 3 * 2 * 1
}
func main() {
	fmt.Println(factorial(12))
}

//when you hit a return status in a function the function returns and is done. DUH.

