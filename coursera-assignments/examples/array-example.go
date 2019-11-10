package main

import "fmt"

func main() {

	// Declare the array
	nums := []int{1, 2, 3, 42}

	// For loop to loop through that array using range.
	// takes the form for index, element := range someSlice
	for i, num := range nums {
		fmt.Println(i, num)
	}

	// If you dont care about the index, you can use _ in its place
	// Heres the same example with skipping keeping the index
	for _, element := range nums {
		fmt.Println("Element:", element)
	}

	// Here is another way to do this
	// Set a counter 'n' to zero, and iterate through
	for n := 0; n < len(nums); n++ {
		fmt.Println(n, nums[n])
	}
}
