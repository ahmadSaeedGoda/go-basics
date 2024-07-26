package main

import "time"

/*
Example of done pattern with for-select
A solution for goroutines leak
*/

func main() {
	done := make(chan bool)

	go func(d <-chan bool) {
		for {
			select {
			case <-d:
				return
			default:
				println("Working...")
			}
		}
	}(done)

	time.Sleep(time.Second * 2)

	close(done)

	println("Routines are done but server is running")

	time.Sleep(time.Hour)
}
