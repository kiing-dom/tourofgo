package main

import (
	"fmt"
	"runtime"
	"time"
)

// switch with no condition (same as saying switch true)
func greeting() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}
}

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		//return the os
		fmt.Printf("%s.\n", os)
	}
	greeting()
}