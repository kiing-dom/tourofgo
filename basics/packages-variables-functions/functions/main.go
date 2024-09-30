package main

import "fmt"

// integers
func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

// strings
func swap(x, y string) (string, string) {
	return y, x
}

// named return values
func split (sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add2(add(42, 13), 13)) // expected: 68
	fmt.Println(swap("hello", "world")) // expected: world hello

	fmt.Println(split(20)) // expected: 8 12
}