// Program to calculate the mean of two numbers
// May 2022
package main

import "fmt"

func main() {
	var x int
	var y int

	x = 1
	y = 2

	// or we can do
	// x := 1.0 assign and = at the same time
	// y := 2.0

	// or even faster, doing auto assign and type infer
	// x, y := 1.0, 2.0
	// only works if they are the same type

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("x=%v, type of %T\n", y, x)

	mean := (x + y) / 2.0
	fmt.Printf("result: %v type of %T\n", mean, mean)
}
