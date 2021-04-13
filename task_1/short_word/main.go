package main

import (
	"fmt"
	"strings"
)

/*
	given a string of words, return the length of the shortest word(s).
	String will never be empty and you do not need to account for different data types.
*/

func findShort(s string) int {
	words := strings.Split(s, " ")
	short := len(words[0])
	for _, word := range words {
		if short > len(word) {
			short = len(word)
		}
	}
	return short
}

func main() {
	fmt.Println(findShort("String will never be empty"))
}
