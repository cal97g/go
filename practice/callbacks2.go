package main

import "fmt"

/*
	This will filter a list according to the function you pass as a 'callback'

	We specify that the callback function must take an integer, and return a bool

	Our function will check if the number is above one, if it is filter will append
	it to the list
*/

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	for _, n := range numbers {
		if callback(n) {
			xs.append(xs, n)
		}
	}
	return xs
}

func main() {
		xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
			return n > 1 //return bool based on this operation
		})
}

/* slice/list/array append: 

var aslice []int
aslice = append(aslice, n)
*/