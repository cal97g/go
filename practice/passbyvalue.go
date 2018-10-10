package main

import "fmt"

//this function accepts the addres of an integer (*int)
func changeMeRef(z *int){
	*z = 20 //derefrence the memory address at z and replace with this(?)
}

func changeMeValue(z int){
	z = 34
}

func main() {
	age := 44
	changeMe(&age) //send the address of age to changeMe (even though its a reference we still pass it as a value)
	fmt.Println(age)

	//now lets see what happens if we try to call me changeMeValue
	changeMeValue(age) //should set to 34
	fmt.Println(age)

	//It doesn't work because it's just a value, we aren't modifying the same place in memory.
	//we are essentially copying the location


}