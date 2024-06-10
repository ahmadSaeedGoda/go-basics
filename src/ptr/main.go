package main

import (
	"fmt"
)

func main() {
	var x int = 10
	var p *int = &x // Pointer to x
	fmt.Println(*p) // Dereference pointer to get the value of x

	*p = 20
	fmt.Println(x) // x is now 20

}
