package main

import "fmt"

func main(){
	//anonymous - it has no name
	func() {
		fmt.Println("I'm driving!")
	}()
	//self executing due to trailing ()
}