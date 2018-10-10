package main

import "fmt"

func main() {
	//ok definition
	var myGreeting = make(map[string]string)
	myGreeting["Max"] = "Hello, Max!"
	myGreeting["Zara"] = "Hello, Zara!"
	delete(myGreeting["Max"])

	if val, exists := myGreeting["Zara"]; exists {
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value does not exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists:", exists)
	}
}