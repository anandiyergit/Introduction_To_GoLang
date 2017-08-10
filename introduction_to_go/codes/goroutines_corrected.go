package main

 import (
 	"fmt"
 	"runtime"
 	"time"
 )

// start main OMIT
 func main() {
 	fmt.Println("Start the goroutines program...")
 	go newMethod() // HL
 	time.Sleep(1 * time.Second) // HL
 	fmt.Println("Final No. of go-routines:", runtime.NumGoroutine())
 	fmt.Println("Terminating the program...")

 }
// end main OMIT

// start newMethod OMIT
func newMethod() {
 	for index := 0; index < 10; index++ {
 		fmt.Println(index)
 		fmt.Println("No. of go-routines:", runtime.NumGoroutine())
 	}
}
// end newMethod OMIT

