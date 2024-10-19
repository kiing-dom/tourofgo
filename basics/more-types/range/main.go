package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	innerPow := make([]int, 10)
	for i := range innerPow {
		innerPow[i] =  1 << uint(i) // == 2** i
	}

	for _, value := range innerPow {
		fmt.Printf("%d\n", value)
	}
}