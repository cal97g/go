package main

import "fmt"

func main() {
	testcase("Facebook")
}

func testcase(acase string){
	switch acase{
		case "Google":
			fmt.Println("WOMEN FOR THE SAKE OF WOMEN.")
		case "Interoute":
			fmt.Println("I WORK HERE")
		case "Facebook":
			fmt.Println("WE HIRE WOMEN WHO OBJECTGIVELY ARENT AS GOOD AT THEIR JOBS AS OTHER MALE CANDIDATES JUST BECAUSE THEY ARE WOMEN.")
		default:
			fmt.Println("Tech giants suck")
	}
}

func testcasefallthrough(bcase string){
	//if we send adam then the next case will evaluate as true (because fallthrough statement)
	//thus if bcase == adam we will print 2 lines: 1: adam, 2: dave
	switch bcase{
		case "adam":
			fmt.Println("adam")
			fallthrough
		case "dave":
			fmt.Println("dave")
		case "james":
			fmt.Println("james")
	}
}

func othercases(ccase int){
	switch ccase{
		case "Tim", "Jenny":
			fmt.Println("Wassup Tim or Jenny")
		case "Marcus":
			fmt.Println("Marcus")
		case "Julian", "Cedric":
			fmt.Println("Soup Julian / Cedric")
		case len(ccase) == 3:
			fmt.Println("Sup 3 char guy")
		default:
			fmt.Println("BEGONE THOT")
	}
}

func casetype(x interface()){
	switch x.(type){
		case int:
			fmt.Println("Int")
		case string:
			fmt.Println("Is str")

	}
}