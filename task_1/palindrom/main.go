package main

import (
	"bytes"
	"fmt"
)

/*
 Description: A palindrome is a word, phrase, number, or other
 sequence of characters which reads the same backward or forward.
  This includes capital letters, punctuation, and word dividers.

 Implement a function that checks if something is a palindrome.

 Examples:
 isPalindrome("anna")   ==> true
 isPalindrome("walter") ==> false
 isPalindrome("12321")    ==> true
 isPalindrome("123456")   ==> false
*/

func isPalindrome(str string) bool {
	l := len(str)
	half := l / 2
	byteStr := []byte(str)
	forw := byteStr[:half]
	if l%2 != 0 {
		half += 1
	}
	rev := byteStr[half:]
	var back []byte
	for i := l/2 - 1; i >= 0; i-- {
		back = append(back, rev[i])
	}
	return bytes.Equal(forw, back)
}

func main() {
	fmt.Println(isPalindrome("anna"))
}
