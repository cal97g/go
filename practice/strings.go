package main

import "fmt"

//raw string literal - any character except another back quote, spacing maintained.
const html = `
	<!DOCTYPE HTML>
	<HTML>
		<H1>HI!</H1>
	</HTML>`

//interpreted string literal - any character except newline and unescaped ". \ is escape.
const intrprtdstr = "Hello I'm Dave"

 
//Strings are immutable, they cannot be changed once they are defined but can be re-assigned

func main() {
	fmt.Println(html)	
	fmt.Println(intrprtdstr)
}
