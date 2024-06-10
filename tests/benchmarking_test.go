package main

import "testing"

func Fib(n int, recursive bool) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		if recursive {
			return Fib(n-1, recursive) + Fib(n-2, recursive)
		}
		a, b := 0, 1
		for i := 1; i < n; i++ {
			a, b = b, a+b
		}
		return b
	}
}

func BenchmarkFibRecursively(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Fib(20, true)
	}
}

func BenchmarkFibIteratively(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Fib(20, false)
	}
}
