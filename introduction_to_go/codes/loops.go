package main

 import (
 	"fmt"
 )

// start main OMIT
 func main() {
 	count := 10
 	sum := 0
 	for index := 0; index < count; index++ {	// HL
 		fmt.Println("The value of the index is:", index)
 		sum += index
 		if sum > 30 {
 			break
 		}
 		fmt.Println("Current value of sum is:", sum)
 	}
 	fmt.Println("The total sum is:", sum)
 }
// end main OMIT
