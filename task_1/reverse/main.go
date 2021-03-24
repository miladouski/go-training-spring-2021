package main

import "fmt"

/**
 * Description: Complete the solution so that it reverses the string passed into it.
 * Example: "world"  =>  "dlrow"
 */

func reverse(word string) string {
	var rev []byte
	for i := len(word) - 1; i >= 0; i-- {
		rev = append(rev, word[i])
	}
	return string(rev)
}

func main() {
	fmt.Println(reverse("world"))
}
