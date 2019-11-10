package main

import "fmt"

func main() {

	// Define a struct type of type Person
	type Person struct {
		name  string
		addr  string
		phone string
	}

	// Now you can declare a type Person
	var p1 Person

	// You can initialize a  struct with 0 values by using the new keyword
	p2 := new(Person)

	// You can also initialize a new struct using a struct literal
	p3 := Person{name: "Brian", addr: "1000 Thompson Rd atp 000", phone: "(512)000-0000"}

	// Use dot notation to assign values to the struct
	p1.name = "Joe"
	p1.addr = "1717 Toomey Rd Apt 5"
	p1.phone = "(512)555-5115"

	fmt.Println(p1, p2, p3)

}
