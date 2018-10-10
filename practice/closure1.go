package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	//I can define a new scope with {}
	{
		fmt.Println(x)
		y := 56
		fmt.Println(y)
	}
	//I can't call y here
}