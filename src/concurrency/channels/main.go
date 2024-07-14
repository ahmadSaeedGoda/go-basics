package main

import (
	"fmt"
)

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i // Send value to the channel
	}
	close(ch) // Indicate producer is done
}

func consumer(ch chan int) {
	for num := range ch { // Receive from the channel
		fmt.Println(num)
	}
}

func main() {
	ch := make(chan int) // Create an unbuffered channel
	go producer(ch)      // Launch producer goroutine
	consumer(ch)         // Call consumer directly
}
