package main

import "fmt"

/*
	An array has fixed size, but slices are dynamically sized, flexible views into the elements of an array
	In practical use, slices are more common than arrays

	The type []T is a slice with elements of type T

	A slice is formed by specifying two indices, a low and high bound, separated by a colon
	e.g. a[low(inclusive) : high(exclusive)]
*/

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}

	var s[]int = primes[1:4] // expected 3, 5, 7
	fmt.Println(s)
}