package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Operation completed")
	case <-ctx.Done():
		fmt.Println("Operation timed out")
	}
}

/*
The context package in Go is used to manage deadlines, cancellation signals,
and other request-scoped values across Goroutines.
It is essential for handling timeouts and cancellations in a clean way.
*/
