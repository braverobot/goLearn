package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Simple line split example
	data := strings.Split("Sam Adams", " ")
	first, last := data[0], data[1]
	fmt.Println("First Name: ", first, "\nLast Name: ", last)

	// open a file
	file, err := os.Open("namefile.txt")
	if err != nil {
		fmt.Println("An error ocurred opening the file")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineArray := strings.Split(line, " ")
		fmt.Println("First Name:", lineArray[0], "\nLast Name:", lineArray[1])
	}

}
