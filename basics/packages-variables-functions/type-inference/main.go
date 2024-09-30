package main

import (
	"fmt"
	"math"
)

func main() {
	v := 42
	v2 := math.Pi
	v3 := 0.867 + 0.5i
	v4 := "https://www.youtube.com/@1KIINGDOM"

	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("v2 is of type %T\n", v2)
	fmt.Printf("v3 is of type %T\n", v3)
	fmt.Printf("v4 is of type %T\n", v4)
}