package main

import "fmt"

/*
	An array has fixed size, but slices are dynamically sized, flexible views into the elements of an array
	In practical use, slices are more common than arrays

	The type []T is a slice with elements of type T

	A slice is formed by specifying two indices, a low and high bound, separated by a colon
	e.g. a[low(inclusive) : high(exclusive)]
	---
	Slices are also like references to arrays
	A slice can describe a section of an underlying array
	Changing the elements of a slice, modifies the corresponding elements of its underlying array
	Slices that share the same underlying array will also see that change
*/

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}

	var s[]int = primes[1:4] // expected 3, 5, 7
	fmt.Println(s)

	// using a slice as a reference
	names:= [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}