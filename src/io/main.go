package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Write to a file
	err := ioutil.WriteFile("example.txt", []byte("Hello, Go!"), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	// Read from a file
	data, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:", string(data))
}
