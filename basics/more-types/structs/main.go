package main

import "fmt"

// A struct is a collection of fields
// struct fields are accessed using a dot
// struct fields can be accessed through struct pointers

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	// accessing struct fields
	v.X = 4
	//pointers to structs
	p := &v
	p.X = 1e9
	fmt.Println(v)
}