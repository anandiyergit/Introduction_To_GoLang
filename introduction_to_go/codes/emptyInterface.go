package main

import "fmt"

// start describe OMIT
func describe(i interface{}) {	// HL
	fmt.Printf("(%v, %T)\n", i, i)
}
// end describe OMIT

// start main OMIT
func main() {
	var i interface{}	// HL
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}
// end main OMIT


