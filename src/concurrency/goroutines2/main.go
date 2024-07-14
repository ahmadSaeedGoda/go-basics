package main

import (
	"fmt"
	"time"
)

func f(a int) {
	x := f1(a)
	f2(x)
}
func f1(a int) int {
	r := step("f1a", a)
	r = step("f1b", r)
	return 1000 * a
}
func f2(a int) {
	r := step("f2a", a)
	r = step("f2b", r)
}
func step(a string, b int) int {
	fmt.Printf("%s %d\n", a, b)
	time.Sleep(1000 * time.Second)
	return 10 * b
}

func main() {
	for i := 0; i < 100000; i++ {
		go f(i)
	}
	//go f(20);
	time.Sleep(1000 * time.Second)
}
