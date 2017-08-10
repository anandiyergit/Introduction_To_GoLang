package main

 import (
 	"fmt"
 	"math"
 )

// start main OMIT
 func main() {
 	/* declare a function variable */
 	getSquareRoot := func(x float64) float64 {	// HL
 		return math.Sqrt(x)
 	}

 	/* use the function */
 	fmt.Println(getSquareRoot(9604))

 }
// end main OMIT
