package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex = &sync.Mutex{} // Create a mutex

func increment() {
	mutex.Lock()         // Acquire the lock
	defer mutex.Unlock() // Release the lock when done
	counter++
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10) // Add 10 goroutines to the wait group
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Final Counter:", counter)
}
