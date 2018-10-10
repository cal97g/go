package main

import (
	"fmt"
	"math"	
)

//A concrete type, properties, fields and methods
type Square struct{
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

/*
	^^^^
	Square implements the area method according to the shape definition
	Therefore Square implements SHAPE
*/


//shape says that a type must implement the area function
//in this case the function must take no args and return float64.
type Shape interface {
	area() float64 
}

//another shape

type Circle struct {
	radius float64
}


//implement shape
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//this function specifies an interface as a parameter type
//It can only be passed types which implement this interface
func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	//in golang interfaces are satisfied IMPLICITLY. We don't have to explicitly state that a metheod IMPLEMENTS a type as in some other languages.
	//Circle and Square both implicitly implement the Shape interface.
	s := Square{10}
	c := Circle{10}

	info(s)
	info(c)
	
}