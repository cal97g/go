package main

import "fmt"

func greatestnum(nums ...int) (int){
	biggest := 0

	for _, n := range nums{
		if n > biggest{
			biggest = n
		}
	}
	return biggest
}

func main(){

	afunc := func(x int) (int, bool){
	return x / 2, x % 2 == 0
	}

	var rtrnint int
	var rtrnbool bool

	rtrnint, rtrnbool = afunc(5)
	
	fmt.Println(rtrnint)
	fmt.Println(rtrnbool)

	biggest := greatestnum(1, 2, 3, 22, 88, 33)
	fmt.Println(biggest)
}

//4: the value of the expression is True because the last check evaluates to true
// !(false && false) -> True -- verified

//to work with all the different methods of calling, with slices, lists etc..
//simple variadic
func foo(numbers ...int) {
	fmt.Println(numbers)
}
