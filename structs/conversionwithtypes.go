package main

import "fmt"

type foo int

func main() {
	var myAge foo = 44
	fmt.Printf("%T %v \n", myAge, myAge)

	var yourAge int = 29
	fmt.Printf("%T %v \n", yourAge, yourAge)

	//will not work
	//fmt.Println(myAge + yourAge)

	//this works
	fmt.Println(int(myAge) + yourAge)

	//this could work:
	fmt.Println(myAge + foo(yourAge))
}