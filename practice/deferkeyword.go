package main

import "fmt"

func hello(){
	fmt.Print("Hello ")
}

func world(){
	fmt.Print(" World\n")
}

func main() {
	defer world() //defers the EXECUTION of world until just before main() exits
	hello()
}