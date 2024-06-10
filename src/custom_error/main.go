package main

import (
	"fmt"
)

type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}

func doSomething() error {
	return &MyError{Message: "something went wrong"}
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
