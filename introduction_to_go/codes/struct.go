package main

 import "fmt"

// show struct OMIT
 type Person struct {	// HL
 	name   string	// HL
 	age    int	// HL
 	gender string	// HL
 }			// HL
// end struct OMIT

// start main OMIT
 func main() {
 	var person1 Person	// HL

 	person1.name = "Anand"
 	person1.age = 26
 	person1.gender = "M"

 	printPerson(person1)
 }
// end main OMIT

// start print OMIT
 func printPerson(person Person) {
 	fmt.Printf("Name : %s\n", person.name)
 	fmt.Printf("Gender : %s\n", person.gender)
 	fmt.Printf("age : %d\n", person.age)
 }
// end print OMIT

