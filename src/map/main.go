package main

import (
	"fmt"
)

func main() {
	// Using make function
	var myMap = make(map[string]int)
	myMap["age"] = 25

	// Using map literal
	anotherMap := map[string]int{
		"age":  30,
		"year": 2022,
	}
	fmt.Println(myMap, anotherMap)

}
