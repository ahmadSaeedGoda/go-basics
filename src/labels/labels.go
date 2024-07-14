package main

import "fmt"

// This code prints from 0 to 4, breaks the loop and go to the end label when i == 5
func foo() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto end
		}
		fmt.Println(i)
	}
end:
	fmt.Println("Exited Loop.")
}

// This code will leave both loops when the if condition is met.
func bar() {
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i*j > 6 {
				fmt.Println(fmt.Sprintf("Exiting because %d * %d is greater than 6", i, j))
				break outerLoop
			}
			fmt.Println(i, j)
		}
	}
}

// This code prints every combination of i and j from 0 to 2, except when equal.
func baz() {
outerLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				continue outerLoop
			}
			fmt.Println(i, j)
		}
	}
}

func main() {
	fmt.Println("Example that prints from 0 to 4, breaks the loop and go to the end label when i == 5")
	foo()
	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Println("Example that will leave both loops when the if condition is met.")
	bar()
	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Println("Example that prints every combination of i and j from 0 to 2, except when equal.")
	baz()
}
