package main

import (
	"fmt"
)

// start main OMIT
func main() {
	const NUMB = 1024 // just a number
	const str = “this is a 日本語 string\n”
	var x, y *float
	var ch = '\u1234'
	fmt.Printf("str is of type %T\n", str)
	fmt.Printf("NUMB is of type %T\n", NUMB)
	fmt.Printf("ch is of type %T", ch)
}
// end main OMIT

