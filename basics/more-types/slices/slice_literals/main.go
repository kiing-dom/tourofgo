package main

import "fmt"

// A slice literal is like an array without the length

// This is an array literal:
// [3]bool{true, true, false}

// creating the same array as above,
// This is a slice literal, then building a slice that references it
// []bool{true, true, false}


func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	} {
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}