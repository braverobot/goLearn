package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Write a program which reads information from a file and represents it in
// a slice of structs. Assume that there is a text file which contains a
// series of names. Each line of the text file has a first name and a last
// name, in that order, separated by a single space on the line. Your program
// will define a name struct which has two fields, fname for the first name,
// and lname for the last name. Each field will be a string of size 20
// (characters). Your program should prompt the user for the name of the
// text file. Your program will successively read each line of the text file
// and create a struct which contains the first and last names found in the
// file. Each struct created will be added to a slice, and after all lines
// have been read from the file, your program will have a slice containing
// one struct for each line in the file. After reading all lines from the
// file, your program should iterate through your slice of structs and print
// the first and last names found in each struct.  Submit your source code
// for the program, “read.go”.

func main() {
	var fileName string
	fmt.Printf("Please enter a filename to read: ")
	fmt.Scan(&fileName)
	// Open a file for processing names
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("An error ocurred opening the file")
	}
	defer file.Close()
	//Define a struct type of type Person
	type Person struct {
		fname string
		lname string
	}
	// Define a slice to keep names in. The sliece will be of len 0 to start
	nameSlice := make([]Person, 1)
	// Read the file into an array
	scanner := bufio.NewScanner(file)
	// Create new person struct and assign the lineArray data to its values
	p1 := new(Person)
	for scanner.Scan() {
		// Assign an array with 2 strings, assigned from a split of a space
		line := scanner.Text()
		lineArray := strings.Split(line, " ")

		// Assign first and last name of each line to elements fname and lname in struct
		p1.fname = lineArray[0]
		p1.lname = lineArray[1]
		nameSlice = append(nameSlice, *p1)
	}
	// Iterate through the slice and print the 2 elements of the struct of each slice
	for n := 0; n < len(nameSlice); n++ {
		fmt.Println(nameSlice[n].fname, nameSlice[n].lname)
	}
}
