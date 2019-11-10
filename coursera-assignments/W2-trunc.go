// Week 2 assignment
package main

import "fmt"

func main() {
	var floaterNum float32
	var truncInt int
	fmt.Printf("Please enter a floating point decimal number: ")
	fmt.Scan(&floaterNum)		// Read in num and place in memory location for floaterNum
	truncInt = int(floaterNum)	// Convert floating point to int
	fmt.Println(truncInt)
}
