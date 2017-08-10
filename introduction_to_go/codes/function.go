package main

import "fmt"

// start add OMIT
func add(x, y int) int {	// HL	
	return x + y
}
// end add OMIT


// start main OMIT
func main() {
	fmt.Println(add(100, 200))
}
// end main OMIT
