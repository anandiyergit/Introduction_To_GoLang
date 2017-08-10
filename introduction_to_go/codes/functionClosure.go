package main

 import "fmt"

// start getSequence OMIT
 func getSequence() func() int {	// HL
 	i := 0
	fmt.Println("Starting new sequence with '0'.")
 	return func() int {	// HL
 		i += 1
 		return i
 	}
 }
// end getSequence OMIT

// start main OMIT
 func main() {
 	/* nextNumber is now a function with i as 0 */
 	nextNumber := getSequence()

 	/* invoke nextNumber to increase i by 1 and return the same */
 	fmt.Println(nextNumber())
 	fmt.Println(nextNumber())
 	fmt.Println(nextNumber())

 	/* create a new sequence and see the result, i is 0 again*/
 	newnextNumber := getSequence()
 	fmt.Println(newnextNumber())
 	fmt.Println(newnextNumber())
 }
// end main OMIT
