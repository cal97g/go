package main

import "fmt"

func main() {
	for i := 50; i <= 140; i++{
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	runes()
}

func runes(){
	arune := 'a'
	fmt.Println(arune)
	fmt.Printf("%T \n", arune) //%T == Type?
}