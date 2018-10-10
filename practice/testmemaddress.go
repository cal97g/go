package main

import "fmt"

/*
& in front of variable name is used to retrieve the address of where this variable’s 
value is stored. That address is what the pointer is going to store.

* in front of a type name, means that the declared variable will store an address 
of another variable of that type (not a value of that type).

* in front of a variable of pointer type is used to retrieve a value stored at given
address. In Go speak this is called dereferencing.
*/

/*
	Double pointer
	x = 8
	y = &x
	z = y ** = &x
*/

/*
	Single Pointer
	x = 42

*/

func main() {
	x := 42
	fmt.Println(x) //the value of x
	fmt.Println(&x) //the address in memory where that value is stored
	foo(&x)
	fmt.Println(x) //this should now be 43?
}

//y *int says we take a pointer to an integer - if you like, the TYPE is POINTER[int]
func foo(y *int){
	fmt.Println(y, "is passed to foo") //an address
	fmt.Println(*y ) // *y says give me the value at this address
	*y = 43 //set the value at this address
	bar(&y)
	fmt.Println(&y)
}

func bar(z **int){
	**z = 44 
	fmt.Println(z)

}
