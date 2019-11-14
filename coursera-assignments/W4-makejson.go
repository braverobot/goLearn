// Write a program which prompts the user to first enter a name, and then
// enter an address. Your program should create a map and add the name
// and address to the map using the keys “name” and “address”,
// respectively. Your program should use Marshal() to create a JSON object
// from the map, and then your program should print the JSON object.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Prompt the user for their name and add it to variable string 'name'
	fmt.Printf("Please enter a first name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	// Prompt the user for their address and add it to variable string 'address'
	fmt.Printf("Please enter an address: ")
	address, _ := reader.ReadString('\n')

	// Strip the newline character from the variables before adding them to the map
	name = strings.TrimSuffix(name, "\n")
	address = strings.TrimSuffix(address, "\n")

	// Create a map (unordered collection of key / value pairs with unique keys)
	// and assign key name and value address to the map
	var myMap map[string]string
	myMap = map[string]string{name: address}

	// Create JSON out of the map by Marshalling it and then printing it as a string
	jsonData, _ := json.Marshal(myMap)
	fmt.Println(string(jsonData))

}
