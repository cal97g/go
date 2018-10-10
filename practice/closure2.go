package main

import "fmt"

func wrapper() func() int {
	x := 0 // or var x int >> 0
	return func() int {
		x++
		return x
	}
}

func main() {
	//x is stored inside wrapper
	increment := wrapper()
	//each time incrememnt is called it increments its internal x
	fmt.Println(increment())
	fmt.Println(increment())
}