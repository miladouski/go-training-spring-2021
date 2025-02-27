package main

import (
	"fmt"
	"sort"
)

/*
 You are given an array of n+1 integers 1 through n. In addition there is a single duplicate integer.

 The array is unsorted.

 An example valid array would be [3, 2, 5, 1, 3, 4]. It has the integers 1 through 5 and 3 is duplicated. [1, 2, 4, 5, 5] would not be valid as it is missing 3.

 You should return the duplicate value as a single integer.
*/

func getDuplicate(numbers []int) int {
	sort.Ints(numbers)
	for i, num := range numbers {
		if num == numbers[i+1] {
			return num
		}
	}
	return 0
}

func main() {
	fmt.Println(getDuplicate([]int{3, 2, 5, 1, 3, 4}))
}
