package main

import "fmt"

//reserve the use of keyword var to indicate that a variable is being initialized to its
//default value - if it will be initialized to something other than a zero value, use 
//the shorthand definition

type person struct {
	first string
	last string
	age int
}

//(p person) is the reciever (not the params)
func (p person) fullName() string {
	return p.first + " " + p.last
}

func main() {
	p1 := person{"James", "Bond", 30}
	//fmt.Println(p1.fullName) -> gives pointer to function location
	fmt.Println(p1.fullName())
}