package main

import "fmt"

/*

	Structs are aggregate types
	encapsulation apparently
	Structs have fields

*/


//reserve the use of keyword var to indicate that a variable is being initialized to its
//default value - if it will be initialized to something other than a zero value, use 
//the shorthand definition

type manufacturer struct {
	name string
	founded int
	country string
}

type car struct {
	name string
	brand manufacturer
	released int
	wheels int
	hoursepower int
}

func main(){
	volkswagen := manufacturer{name: "Volkswagen", founded: 1895, country: "Germany"}
	ford := manufacturer{name: "Ford", founded: 1912, country: "United States"}

	mondeo := car{name: "Mondeo", brand: ford, released: 1990, wheels: 4, hoursepower: 100}
	polo := car{name: "Polo", brand: volkswagen, released: 1995, wheels: 4, hoursepower: 200}

	fmt.Println(volkswagen)
	fmt.Println(ford)

	fmt.Println(mondeo)
	fmt.Println(polo)
	fmt.Println(polo.name, " ", polo.brand.name)
}