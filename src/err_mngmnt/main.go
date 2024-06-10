package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(divide(1, 0))
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
