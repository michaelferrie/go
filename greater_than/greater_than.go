// Program to calculate if x is greater than y
// May 2022
package main

import (
	"fmt"
)

func main() {
	x := 10

	if x < 5 {
		fmt.Println("x is bigger")
	}

	if x > 100 {
		fmt.Println("x is bigger than 100")
	} else {
		fmt.Println("x is not that big")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is the right size")
	}

}
