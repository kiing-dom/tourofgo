package main

import "fmt"

// A struct is a collection of fields
// struct fields are accessed using a dot
// struct fields can be accessed through struct pointers
// struct literals denote newly added struct values by listing the value of their fields

type Vertex struct {
	X int
	Y int
}

// struct literals
var (
	v1 = Vertex{1, 2} // has type Vertex
	v2 = Vertex{X : 1} // implicitly sets Y to 0
	v3 = Vertex{} // implicity sets both X and Y to 0
	p = &Vertex{1, 2} // has type *Vertex
)

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)

	// accessing struct fields
	v.X = 4

	//pointers to structs
	p := &v
	p.X = 1e9
	fmt.Println(v)
	fmt.Println(v1, p, v2, v3)
}
