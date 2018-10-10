package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string `json:"-"` //field tags
	Last string
	Age int
	notExported int
}

func main() {
	p1 := Person{"James", "Bond", 20, 007}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs ))

	var p2 Person
	json.Unmarshal(bs, &p2) //byte slice + pointer to p2
	fmt.Println(p2.Last)
	fmt.Println(p2.Age)
	fmt.Println(p2)


}