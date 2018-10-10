package main

import "fmt"


//this should really be a slice of maps or something
func main() {
	records := make([][]string, 0)

	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "75.00"
	//store??
	records = append(records, student1)
	//student 2
	student2 := make([]string, 4)
	student2[0] = "Dankiels"
	student2[1] = "David"
	student2[2] = "55.00"
	student2[3] = "64.00" 

	records = append(records, student2)
	fmt.Println(records)


}