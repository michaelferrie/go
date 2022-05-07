// check for highest number in a slice
// bubble sort method, good on a small slice
package main

import (
	"fmt"
)

// my solution first
func main() {
	// create slice
	nums := []int{16, 8, 42, 55, 23, 15}

	// set up test int at zero
	j := 0

	// iterate over the slice with i
	for _, i := range nums {

		// set j to i if i > j
		if i > j {
			j = i
		}
		// fmt.Println(j) - just to test
	}
	// return result
	fmt.Println("The highest number in the slice is", j)

	// instructors solution - he starts with first two values
	max := nums[0]               // set max to first value in slice
	for _, i := range nums[1:] { // create iterator start on 2nd value
		if i > max {
			max = i
		}
	}
	fmt.Println("other way", max)
}
