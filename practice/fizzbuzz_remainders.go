package main

import "fmt"

/*

*/


func main() {
	x := 13 % 3 //sets x to an int and its value to the remainder of 13/3
	fmt.Println(x) //1

	if x == 1 {
		fmt.Println("Odd!")
	} else {
		fmt.Println("Even!")
	}

	//for each i below 70 (iterate i)
	for i := 1; i < 70; i++ {
		fizzbuzz(i)
	}
}

func fizzbuzz(number int) {
	if (number % 3 == 0) && (number % 5 == 0) {
		fmt.Println("FizzBuzz")
	} else if number % 3 == 0 {
		fmt.Println("Fizz")
	} else if number % 5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(number)
	}
}

