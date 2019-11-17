package main

import "fmt"

// function that takes a slice as input, makes changes to the slice.
// NOTE: its making changes to the actual slice, not a copy of it,
// therefore the changes are global, not local to the function
func mySliceFunc(slice []int) {
	slice[0] = slice[0] + 1
}

func main() {

	// Declare the slice
	a := []int{1, 2, 3, 4}
	// Print the slice that was declared
	fmt.Println(a)
	// Make changes to the slice in the function
	// NOTE: No need to return the slice! The Slice change is global
	mySliceFunc(a)
	// Print the slice after it was changed
	fmt.Println(a)
}
