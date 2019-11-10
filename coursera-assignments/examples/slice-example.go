package main

import "fmt"

func main() {

	// Declare the array
	arr := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	s1 := arr[2:4]

	// s1 is basically another array assigned with the elements stated in the declaration
	// Note how it reinitialized s1 at 0, so basically s1[0] is the same as arr[2]
	fmt.Println(s1[0], s1[1])
	fmt.Println(arr[2], arr[3])

	// Length of the slice- The number of elements in the slice
	// Capacity of the slice The maximum number of elements that can be in a slice,
	// UP TO the size of the array that it is a slice of
	// If you would like to return lenght and capacity, you can use len and cap
	fmt.Println("Length of slice s1: ", len(s1))
	fmt.Println("Capacity of slice s1: ", cap(s1))
	fmt.Println("Capacity of full array arr: ", cap(arr))

	// Slice Literals can be used to initialize a slice (and create an underlying array)
	sli := []int{1, 2, 3}
	fmt.Println(sli)

	// You can append to the slice using the append function:  append(slice-name, data-to-append)
	sli = append(sli, 4)
	fmt.Println(sli)

}
