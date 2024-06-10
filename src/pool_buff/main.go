package main

import (
	"fmt"
	"sync"
)

func main() {
	var pool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024) // Create a 1KB buffer
		},
	}

	// Get a buffer from the pool
	buffer := pool.Get().([]byte)

	// Use the buffer
	buffer[0] = 1

	// Put the buffer back into the pool
	pool.Put(buffer)

	fmt.Println("Buffer reused from pool")
}

/*
	Minimizing memory allocations and using structures like sync.Pool can help optimize memory usage.
*/
