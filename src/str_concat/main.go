package main

import (
	"fmt"
)

func main() {
	// Using + operator
	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)

	// Using fmt.Sprintf
	result = fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(result)

}
