// Program takes user integer input and adds it to a slice, sorts, and prints it.
// If input is X(x) end. If input is non integer other than X(x), reprompt
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	sli := make([]int, 3)

	// Declare user input variable and exit check variable
	var myString, exitCheck string

	// Start loop with exit not being true
	for exitCheck != "X" {
		fmt.Printf("Please enter in a number (X to quit): ")

		// Read in variable as string
		fmt.Scan(&myString)

		// Convert string to upper case for exit check on x as well as X
		myString = strings.ToUpper(myString)

		// Test for X to break out of loop
		if myString == "X" {
			fmt.Println("X has been encountered, Goodbye!")
			exitCheck = myString
		} else {

			// If X is not encountered, attempt to convert input to integer
			myInt, err := strconv.Atoi(myString)

			// Add Int to slice if no error is encountered
			if err == nil {
				// fmt.Println(myInt)
				sli = append(sli, myInt)

				// Sort the slice
				sort.Ints(sli)

				// Print the slice
				fmt.Println(sli)

			} else { // if error is encountered, no X nor an integer has been found
				fmt.Println("Integer or X not found")
			}
		}
	}
}
