package main

 import "fmt"

// show Mutable OMIT
 type Mutable struct {
 	a int
 	b int
 }
// end Mutable OMIT

// start staytheSame OMIT
 func (m Mutable) staytheSame() {
 	m.a = 10
 	m.b = 20
 }
// end staytheSame OMIT

// start mutate OMIT
 func (m *Mutable) mutate() {	// HL
 	m.a = 10
 	m.b = 20
 }
// end mutate OMIT

// start main OMIT
 func main() {
 	m := Mutable{5, 7}
 	fmt.Println(m)
 	m.staytheSame()
 	fmt.Println(m)
 	m.mutate()	// HL
 	fmt.Println(m)
 }
// end main OMIT
