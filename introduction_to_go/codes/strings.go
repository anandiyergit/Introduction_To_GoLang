package main

 import "fmt"
 import "strings"

 func main() {

 	str1 := "Anand, you have started, with strings!"
 	str2 := "Refer to the documentation for more information on this!"
 	fmt.Println("Equal ?", (str1 == str2))	// HL
 	strings.Compare(str1,str2)	// HL
 	fmt.Println(strings.SplitAfter(str1, ","))	// HL
 }
