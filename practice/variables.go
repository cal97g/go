package main
import "fmt"

func main(){
	var defineandassign string = "wassup"
	shorthand := "hello!" //type inferred
	fmt.Println(defineandassign)
	fmt.Println(shorthand)
}


/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
