package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			fmt.Println(i, " - ", j)
		}
	}
	loopbreakcontinue(4)
}

//probably bad naming convention
func loopbreakcontinue(i int){ 
	for {
		if i < 10 {
			i++
			fmt.Println(i)
			continue
		} else {
			break //done
		}
	}
}
