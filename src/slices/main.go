package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Creating a slice
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", s)

	// Slicing the slice
	s = s[1:4]
	fmt.Println("Sliced:", s)

	represent_slice_as_json()
}

func represent_slice_as_json() {
	var s1 []int
	s2 := []int{}
	s3 := make([]int, 5)
	s4 := make([]int, 0, 5)

	j1, err1 := json.Marshal(s1)
	j2, err2 := json.Marshal(s2)

	if err1 != nil || err2 != nil {
		fmt.Fprintf(os.Stderr, "Couldn't Marshal the slice")
	}

	// fmt.Printf("Slice as json: %10s \n", j)
	fmt.Println(string(j1))
	fmt.Println(string(j2))
	fmt.Printf("%T \n", j1)
	fmt.Printf("%T \n", j2)
	fmt.Printf("Cap of s1 = %d \n", cap(s1))
	fmt.Printf("Cap of s2 = %d \n", cap(s2))
	fmt.Printf("len of s1 = %d \n", len(s1))
	fmt.Printf("len of s2 = %d \n", len(s2))
	fmt.Printf("Cap of s3 = %d \n", cap(s3))
	fmt.Printf("len of s3 = %d \n", len(s3))
	fmt.Printf("Cap of s4 = %d \n", cap(s4))
	fmt.Printf("len of s4 = %d \n", len(s4))
}
