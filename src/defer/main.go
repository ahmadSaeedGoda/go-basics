package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

/*
The defer statement in Go is used to ensure that a function call is performed later in a program’s execution,
usually for purposes of cleanup.

Deferred functions are executed in LIFO (last in, first out) order just before the surrounding function returns.
*/
