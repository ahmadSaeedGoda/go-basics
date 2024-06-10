package main

import "fmt"

func main() {
	/*
		var: Used for declaring variables with a specified type.
		:=: Short-hand syntax for declaring and initializing variables without specifying the type.
		const: Used for declaring constants whose value cannot change.
	*/

	var name string = "Alice" // using var
	age := 30                 // using :=
	const country = "Egypt"   // using const

	fmt.Println(name, age, country)
}
