package main

import "fmt"

func greeting() {
	fmt.Println("Hello World")
}

func makeGreeter() func() string {
	return func() string {
		return "Hello Wolrd function returned"
	}
}

func main(){
	//anonymous function - := < func expression
	afunction := func(){ 
		fmt.Println("Only way to have a function in a function")
	}
	fmt.Printf("%T \n", afunction)//type func()
	
	greeting()

	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Printf("%T \n", greet) //type func() string
}

//vs

