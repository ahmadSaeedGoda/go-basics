package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

/*
	Generics in Go provide a powerful way to write flexible, reusable, and type-safe code.
	They open up new possibilities for abstraction & code organization.
	With generics, Go becomes an even more versatile language, capable of handling a wider range of
	programming paradigms efficiently.
*/

// 1. Function Generics
// This is a generic function that works with any type T
func Print[T any](value T) {
	fmt.Println(value)
}

// 2. Type Constraints
// Numeric is a constraint that allows only numeric types
type Numeric interface {
	int | float64
}

func Add[T Numeric](a, b T) T {
	return a + b
}

// 3. Generic struct
type Pair[T any] struct {
	First, Second T
}

// 4. Type Sets
// Type sets are used to define the set of types that a type parameter can take.
// This can be useful for more complex constraints.

// Here's a Constraint that allows only types with a String method
type Stringer interface {
	String() string
}

// PrintString is a generic function that works with any type that implements the Stringer interface
func PrintString[T Stringer](value T) {
	fmt.Println(value.String())
}

type MyType struct {
	Name string
}

func (m MyType) String() string {
	return m.Name
}

// 5. Type Inference
// allowing you to omit type arguments when they can be inferred from the context.

// Double is a function that multiplies numbers by 2 using type inference.
func Double[T int | float64](value T) T {
	return value * 2
}

// 6. Using Generics in the Standard Library
// The standard library has been updated to make use of generics in several packages.
// For example, the constraints package defines various type constraints that can be used with generics.

// Min function that works with any ordered type
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	Print(42)      // Prints: 42
	Print("hello") // Prints: hello
	Print(3.14)    // Prints: 3.14

	fmt.Println(Add(1, 2))     // Prints: 3
	fmt.Println(Add(1.1, 2.2)) // Prints: 3.3
	// fmt.Println(Add(1, "2"))  // Compilation error: "string does not implement Numeric"

	p1 := Pair[int]{1, 2}
	p2 := Pair[string]{"hello", "world"}
	// p3 := Pair[int]{1, "yes"} // cannot use "yes" (untyped string constant) as int value in struct literal. compiler IncompatibleAssign
	fmt.Println(p1) // Prints: {1 2}
	fmt.Println(p2) // Prints: {hello world}

	myValue := MyType{"Generics"}
	PrintString(myValue) // Prints: Generics

	fmt.Println(Double(2))   // Prints: 4
	fmt.Println(Double(2.5)) // Prints: 5
	// fmt.Println(Double("aa")) // Error: string does not satisfy int | float64 (string missing in int | float64) compiler InvalidTypeArg

	fmt.Println(Min(2, 3))       // Prints: 2
	fmt.Println(Min(3.14, 1.59)) // Prints: 1.59
	fmt.Println(Min("a", "b"))   // Prints: a
}

// 7. Benefits of Generics
// Code Reusability: Write functions and types that work with any data type.
// Type Safety: Catch type errors at compile time.
// Performance: Generics in Go are designed to be as performant as non-generic code.
