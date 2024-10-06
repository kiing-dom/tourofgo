package main

import "fmt"
/*
The type [n]T is an array of n values of type T.

The expression

var a [10]int
declares a variable a as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized. 
This seems limiting, but don't worry; Go provides a convenient way of working with arrays.
*/

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1]) // prints the contents of the array
	fmt.Println(a) // prints the array itself enclosed in []

	primes := [6]int{2, 3, 5, 7, 11, 13} // filling an array
	fmt.Println(primes)
}
