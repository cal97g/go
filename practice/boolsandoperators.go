package main

import "fmt"

func main(){
	cake := true
	icecream := true
	cookies := false

	//or
	if true || false {
		fmt.Println("This runs because one is true")
	}

	x := 14

	if x == 14 || x % 2 == 1 {
		fmt.Println("This should run because x == 14")
	}

	if cake && icecream {
		fmt.Println("We have both cake and icecream!")
	}

	if cake && icecream && cookies {
		fmt.Println("Will never run because we don't have cookies!! :((")
	}

}