package main

import "fmt"

// when making slices we can omit the upper and lower bounds and use defaults
// the defaults are:
// lower bound: 0
// upper bound: length of the slice

// for the array var a[10] int
// these slice expressions are equivalent
// a[0:10] = a[:10] = a[0:] = a[:]

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}