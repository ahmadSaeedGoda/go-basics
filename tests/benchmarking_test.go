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

/*
	Run the following command, if your IDE doesn't help running in GUI, for testing with benchmarking in CLI:
	bash
	```
	$ go test -bench=. ./tests/
	```

	Expected output result should be similar to the following:
	```
	$ go test -bench=. ./tests/
		goos: linux
		goarch: amd64
		pkg: examples/tests
		cpu: 12th Gen Intel(R) Core(TM) i5-12400
		BenchmarkFibRecursively-12         37012             31629 ns/op
		BenchmarkFibIteratively-12      160066993                7.380 ns/op
		PASS
		ok      examples/tests  3.442s
	```
	As you could see, the difference between iterative vs recursive approach, performance wise
	and time complexity wise, is way remarkable!
	Let alone the space complexity.
*/
