package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x * x + y * y))
	var z uint = uint(f)
	var f2 float64 = float64(z)

	fmt.Printf("%v, %v, %v, %.1f", x, y, z, f2)
}