// Strings and things
// May 2022
package main

import (
	"fmt"
)

func main() {
	// create a string
	words := "People fail to succeed because they major in minor things"

	// print out
	fmt.Println(words)

	// print out length in bytes
	fmt.Println(len(words))

	// print out type and
	fmt.Printf("words[0] = %v (type %T)\n", words[0], words[0])

	// print slice
	fmt.Println(words[4:11])
	fmt.Println(words[:11])
	fmt.Println(words[:4])

	// concatenate
	fmt.Println(words[4:11])

	// multi line
	big_string := `
this is a very 
long 
multi
line
string`
	fmt.Println(big_string)

}
