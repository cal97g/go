package main

import(
	"encoding/json"
	"os"
)

type Person struct {
	FirstName string
	LastName string
	Age int
}

func main() {
	p1 := Person{"James", "Bond", 20}
	json.NewEncoder(os.Stdout).Encode(p1)
}