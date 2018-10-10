package main

import (
	"fmt"
	"net/http"
	"bufio"
	"os"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		fmt.Println(err)
	}

	words := make(map[string]string)

	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = ""
	}
	//initiate variable, then check it in an if statement :o
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	i := 0
	for k, v := range words {
		fmt.Println(k, v)
		if i == 100 {
			break
		}
		i++
	}

	fmt.Println(len(words))
}