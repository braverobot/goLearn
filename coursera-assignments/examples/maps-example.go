package main

import "fmt"

func main() {

	// Declare the map: var idMap map[KeyType]valueType
	var idMap map[string]int
	idMap = make(map[string]int)

	// You can create a map literal as well, which adds values to the map
	// The following declares a map with a key of Joe and data of 123
	idMap2 := map[string]int{"Joe": 123, "Jim": 234, "John": 824, "Brian": 1337}

	// To print both the empty declared map and the populated map:
	fmt.Println(idMap, idMap2)

	// To access a value in a map you can use the key:
	fmt.Println(idMap2["Joe"])

	// To add data to a map
	idMap["Jane"] = 456
	fmt.Println(idMap)

	// To delete info from a map, call the delete function on it
	delete(idMap2, "Jim")
	fmt.Println("Deleted Jim from Map2->", idMap2)

	id, p := idMap["Bob"]
	fmt.Println("Key Bob Present?:", p, "...  Value:", id)

	// To iterate through a Map, use a for loop with the range keyword
	for key, val := range idMap2 {
		fmt.Println("-->", key, val)
	}
}
