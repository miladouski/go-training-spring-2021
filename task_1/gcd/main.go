package main

import "fmt"

/*
	Find the greatest common divisor of two positive integers. The integers can be large, so you need to find a clever solution.

	The inputs x and y are always greater or equal to 1, so the greatest common divisor will always be an integer that is also greater or equal to 1.
*/

func compute(x, y int) int {
	if y == 0 {
		return x
	}
	return compute(y, x%y)
}

func main() {
	fmt.Println(compute(30, 36))
}
