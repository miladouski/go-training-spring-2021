package main

import (
	"fmt"
	"regexp"
)

/*
	Write a simple regex to validate a username. Allowed characters are:
	lowercase letters,
	numbers,
	underscore
	Length should be between 4 and 16 characters (both included).
*/

func isUsername(username string) bool {
	regex := "^[a-z0-9_]{4,16}$"
	matched, _ := regexp.MatchString(regex, username)
	return matched
}

func main() {
	fmt.Println(isUsername("lowercase_"))
}
