package main

import (
	"fmt"
	"sort"
)


type people []string



func main() {
	ppl1 := []string{"Zeno", "John", "Al", "Jenny"}
	ppl2 := people{"Zeno", "John", "Al", "Jenny", "Gomez", "Dan"}
	n := sort.IntSlice([]int{7, 5, 3, 55, 33, 22, 19, 9, 12, 3})

	sort.Strings(ppl1)
	sort.Strings(ppl2)
	n.Sort()

	fmt.Println(n)
	fmt.Println(ppl1)
	fmt.Println(ppl2)

	fmt.Prinln(sort.Reverse(sort.StringSlice(ppl1)))
	fmt.Println(sort.Reverse(sort.StringSlice(ppl2)))
	fmt.Println(sort.Reverse(n))
}
