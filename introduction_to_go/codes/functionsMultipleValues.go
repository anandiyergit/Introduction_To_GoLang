package main

 import "fmt"

// start swap OMIT
 func swap(x, y string) (string, string) {	
 	return y, x	// HL
 }
// end swap OMIT

// start main OMIT
 func main() {
 	a, b := swap("Anand", "Iyer")
 	fmt.Println(a, b)
 }
// end main OMIT
