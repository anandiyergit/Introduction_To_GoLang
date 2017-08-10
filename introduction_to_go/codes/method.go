package main

import (
	"fmt"
	"math"
)

// show Vertex OMIT
type Vertex struct {
	X, Y float64
}
// end Vertex OMIT

// start Abs OMIT
func (v Vertex) Abs() float64 {	// HL
	return math.Sqrt(v.X*v.X + v.Y*v.Y)	// HL
}
// end Abs OMIT

// start main OMIT
func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
// end main OMIT
