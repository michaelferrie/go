// Fizz buzz in Go
// May 2022
package main

import (
	"fmt"
)

func main() {

	// define x as an int
	x := 1

	// loop over 1 -20 and print
	for x < 21 {

		// first check if the number divides by 3 and 5 print fizbuzz
		if x%3 == 0 && x%5 == 0 {
			fmt.Println(x, "FizzBuzz")

			// then check if it divides by 3, print fizz
		} else if x%3 == 0 {
			fmt.Println(x, "Fizz")

			// then check if it divides by 5, print buzz
		} else {
			if x%5 == 0 {
				fmt.Println(x, "Buzz")
			}
		}
		x++
	}
}
