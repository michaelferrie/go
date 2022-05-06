// Program demonstrate switch case
// May 2022
package main

import (
	"fmt"
)

func main() {

	n := 2

	switch n {

	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	}

	// naked switch with conditionals inside the statement

	switch {
	case n > 100:
		fmt.Println("Its big")
	case n < 100:
		fmt.Println("It's small")
	default:
		fmt.Println("Its none of these")
	}

}
