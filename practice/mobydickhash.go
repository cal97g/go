package main

import (
	"fmt"
	"net/http"
	"bufio"
)

//code creates a slice with 200 values
//it converts letters to their ascii code
//everytime it finds a word, the ascii value for that word is iterated


func main() {

	HashBucket := func(word string) int {
		return int(word[0]) //return the ascii code for the first letter in the word
	}

	res, err := http.Get("https://gist.githubusercontent.com/StevenClontz/4445774/raw/1722a289b665d940495645a5eaaad4da8e3ad4c7/mobydick.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(res.Body)
	scanner.Split(bufio.ScanWords)

	buckets := make([]int, 200)

	for scanner.Scan() {
		n := HashBucket(scanner.Text()) //find the ascii code for the first character in word
		buckets[n]++ //iterate that code to show another word starts with that letter
		//I would honestly use a map: make(map[string]integer) where string is word/char and int is count
	}

	fmt.Println(buckets[65:123])


}