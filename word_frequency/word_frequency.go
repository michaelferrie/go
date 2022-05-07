// count how many times a word appears in text
package main

import (
	"fmt"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	ss = make([][]byte, 0)

	for _, i := range text {
		fmt.Print(string(i))
		ss = append(ss, i)
		fmt.Println(string(ss[i]), &next[0])
	}
	fmt.Println()
}
