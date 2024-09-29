package main

import "fmt"

func add(x int, y int) int {
	return x + y;
}

func main() {
	fmt.Println(add(add(42, 13), 13)) // expected 68
}