package main

import (
    "fmt"
    "strconv"
  )

func main() {
	var x int = 12
	var y = 12.12343
	//widening conversion
	fmt.Println(y + float64(x))

	var a rune = 'A'
	var b int32 = 'b'
	fmt.Println(string(a))
	fmt.Println(b)
  fmt.Println([]byte{'h', 'e', 'l', 'l', 'o'}) //[104 101 108 108 111] -- ascii not utf-8 > utf has two numbers per character
  fmt.Println(string([]byte{'h', 'e', 'l', 'l', 'o'}))

  i, _ := strconv.Atoi("-42")
  s := strconv.Itoa(-42)

  fmt.Println(i)

  fmt.Println(s)
  /* all of the above is CONVERSION.

  Below we try some assertions.

  Type assertions in Go
  Type assertions are used to check if value held by interface type variable either implements desired interface or is of a concrete type.
  */

  /*
    var val interface{} = 7
    fmt.Println(val + 6)
    //fmt.Printf("%T \n", val)  --> int
    //  mismatched types
  */

  var val interface{} = 7
  fmt.Println(val.(int) + 6) //assert that it is an int
  



}
