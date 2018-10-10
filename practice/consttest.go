package main

import "fmt"

const p = "death and taxes" //const type infer works

const x, y = "hello x", "goodbye y"

const (
    Pi = 3.14
    Language = "Go"
)

func main() {
    fmt.Println(p)
    fmt.Println("" + Language, Pi)
    fmt.Println(x + y)
}


const typedConst string = "typed String!"
const untypedConst = "hmm"


/*

https://blog.golang.org/constants

*/