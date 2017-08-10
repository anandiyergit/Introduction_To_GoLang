package main

 import (
 	"fmt"
 	"strconv"
 	"time"
 )

 // 1. Make the cake.
 // 2. Send the cake.

// start i OMIT
 var i int
// end i OMIT

// start main OMIT
 func main() {
 	cake := make(chan string)	// HL
 	for index := 0; index < 5; index++ {
 		go makeCake(cake)	// HL
 		go sendCake(cake)	// HL
 	}
 	time.Sleep(1 * time.Second)
 }
// end main OMIT

// start makeCake OMIT
 func makeCake(cake chan string) {
 	i = i + 1
 	cakeName := "Cake " + strconv.Itoa(i)
 	fmt.Println(cakeName, " is made and ready to be sent!")
 	cake <- cakeName	// HL
 }
// end makeCake OMIT

// start sendCake OMIT
 func sendCake(cake chan string) {
 	msg := <-cake	// HL
 	fmt.Println(msg, " is sent for delivery!")
 }
// end sendCake OMIT
