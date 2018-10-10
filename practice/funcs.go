package main

import "fmt"



func main() { 
	var namea, nameb, age = greet("David", "Davis", 43)

	//call variadic function
	avg := average(34, 22, 53, 11, 59, 44, 22)
	//or
	data := []int(22, 44, 55, 34, 66, 33, 17)
	avg = average(data...)  //... -> expands to individual int types - not the slice type

	//anonymous function - function 'expression' 
	greeting := func() {
		fmt.Println("Hello World!") //only way to define a function within a function
	}

	greeting()
	fmt.Printf("%T", greeting) //type is func()
}

//variadic - takes as many numbers as you like
func average(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total / len(numbers)
}

func candrink(age int) bool{
	if age >= 18{
		return true
	} else {
		return false
	}
}

//also multiple arg definition

func canride(age, heightcm int) string {
	if (age >= 12) && (heightcm > 100){
		return "Can ride!"
	} else {
		return "You cannot ride!"
	}
} 

//multiple return

func greet(fname, lname string, age int) (string, string, int){
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname), fmt.Sprint(fname, lname, age)
}


//named return - less readable - complicates scope

func greet2(fname, lname string) (s string){
	s = fmt.Sprint(fname, lname)
	return
}

/*variadic functions

These functions accept as many parameters of one type as desired.

IE ...int //zero or more int arguments



*/

