package	main

import "fmt"

//visit takes a slice of ints, and a function which accepts an int.
//callback could be named anything
func visit(numbers []int, callback func(int)){
	for _, n := range numbers {
		callback(n)
	}
}

//we can do callbacks in go, but this is more of a functional feature
func main(){
	visit([]int{1,2,3,4}, func(n int){
		fmt.Println(n)
	})
}

//prefer readability over complexity and 'cleverness' IE.. PYTHONIC code.