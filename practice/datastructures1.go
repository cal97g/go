package main

import "fmt"

func main(){
	//array -> Tuple of SAME TYPE - must define length of array. Does not change size.
	//Not dynamic.

	//var will initialize our ints to 0s
	var arrayA [60]int
	//if it was a string it would just be a list of empty strings
	fmt.Println(arrayA)
	fmt.Println(arrayA[4:10])

	//ascii character example with strings
	var arrayB [60]string

	for i := 0; i < 59; i ++ {
		arrayB[i] = string(i)
	}

	fmt.Println(arrayA)
	fmt.Println(arrayB)


	//slice -> Python List uninitialized value is nil -> reference type. Dynamic.
	//make is used for slices because underlying setup is required (array init, etc)
	//you could define a total capacity for your slice.. make([]T, length, capacity)
	//if you exceed the capacity a new array must be created, twice the size of the current array
	//the old array will then be disgarded, and reference updated to the new array
	//see the example of this in sliceexample.go 

	var slice = []int{1,4,6,22,44,2,1,} //slice defined - unable to specify capacity here
	fmt.Println(slice)
	fmt.Println(slice[2:4]) //slicing a slice - like python only read 2 elements, stops at 4
	fmt.Println(slice[2]) //specific index 

	//you can increment/change an integer in a slice as normal
	fmt.Println(slice[2])
	slice[2]++
	//or slice[2] += 1
	//or slice[2] = slice[2] + 1
	fmt.Println(slice[2])

	slicefromarray := new([100]int)[0:50] //magic - defined capacity and length of slice
	fmt.Println(slicefromarray)

	//preferred slice defenition method
	sliceb := make([]int, 10, 20) //slice of int, length of 10, capacity of 20
	fmt.Println(sliceb)

	//attempting to access an index that doesn't exist in go
	student := []string{} //define a slice of string - but no items or length
	//student[0] = "hello" //will not work!
	student = append(student, "hello") //append must be used to create the index


	//map -> python dict uninitialized value is nil -> reference type
	//a map is unordered -> built on a hash table
	mymap := make(map[string]int)
	mymap["k1"] = 7
	mymap["k2"] = 44

	delete(mymap, "k2")
	fmt.Println("map: ", mymap)

	//retrieve value from map
	value, ok := mymap["k2"]
	fmt.Println(value, " ", ok) //0 false -> because we deleted k2 above

	value, ok := mymap["k1"] //7 true

	//mymap["key"] (value, bool) -> 
	//first return will be value, second is bool of whether it exists

	//struct -> object -> 'composite type' presumably made up of other types as defined 

}