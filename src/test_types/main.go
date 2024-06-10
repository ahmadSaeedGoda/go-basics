package main

import (
	"fmt"
	"testing"
)

func main() {

}

// Unit Test
func TestMultiply(t *testing.T) {
	result := Multiply(2, 3)
	if result != 6 {
		t.Errorf("Multiply(2, 3) = %d; want 6", result)
	}
}

// Benchmark Test
func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(2, 3)
	}
}

// Example Test
func ExampleMultiply() {
	fmt.Println(Multiply(2, 3))
	// Output: 6
}

func Multiply(i1, i2 int) int {
	return i1 * i2
}
