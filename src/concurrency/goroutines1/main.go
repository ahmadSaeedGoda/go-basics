package main

import (
	"fmt"
	"time"
)

func saySomething(a, b func()) {
	a()
	b()
}

func foo() {
	fmt.Print("foo")
}

func bar() {
	fmt.Print("bar")
}

func talkForAWhile() {
	for {
		saySomething(foo, bar)
	}
}

func main() {
	go talkForAWhile()
	time.Sleep(1 * time.Second)
}
