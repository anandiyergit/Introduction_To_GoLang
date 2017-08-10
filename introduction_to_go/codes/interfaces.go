// _Interfaces_ are named collections of method signatures.

package main

import "fmt"
import "math"

// Here's a basic interface for geometric shape.

// start geometry OMIT
type geometry interface {	// HL
	area() float64		// HL
}		

// end geometry OMIT

// show circle OMIT
type circle struct {
	radius float64
}
// end circle OMIT

// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// The implementation for `circle`s.

// start area OMIT
func (c circle) area() float64 {	// HL
	return math.Pi * c.radius * c.radius	// HL
}

// end area OMIT
// If a variable has an interface type, then we can call
// methods that are in the named interface. Here's a
// generic `measure` function taking advantage of this
// to work on any `geometry`.

// start measure OMIT
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
}

// end measure OMIT

// The `circle` and `rect` struct types both
	// implement the `geometry` interface so we can use
	// instances of
	// these structs as arguments to `measure`.
	
// start main OMIT
func main() {
	c := circle{radius: 5}
	measure(c)
}

// end main OMIT

