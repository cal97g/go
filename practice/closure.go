import "fmt"

//this function returns a function

func wrapper() func() int{ //outer scope
	x := 0 
	return func() int { //inner scope
		x++ 
		return x
	} //close inner scope
} //close outer scope

func main(){
	increment := wrapper() 		//assign our anonymous function to increment
	fmt.Println(increment()) 	//prints 1
	fmt.Println(increment()) 	//prints 2 because the wrapper 
								//scope will remember the value of x
}


/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
