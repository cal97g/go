package main

import "fmt"

type agent struct{
	name string
	number int
}

type mission struct{
	agent
	alias string
}

func main() {
	//test promotion

	myagent := agent{"James Bond", 007}
	hismission := mission{myagent, "To kill the bad guys"}

	fmt.Println(hismission.name) //should print James Bond
	//this is because the mission struct has the agent type as an anonymous field..
	//so it's 'promted'
}