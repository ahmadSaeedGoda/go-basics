package main

import "fmt"

var message string

func init() {
	message = "Hello, World!"
}

func main() {
	fmt.Println(message)
}
