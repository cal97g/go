//If you want to change a value by passing it to a function, you must pass the address
// (if it is not a reference type already)
package main

import "fmt"

//slices and maps are reference types

func main(){
	m := make([]string, 1, 25)
	fmt.Println(m)
	changeMe(m)
	fmt.Println(m)
}

func changeMe(z []string) {
	z[0] = "Callam"
	fmt.Println(z)
}