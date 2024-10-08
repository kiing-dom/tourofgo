package main

/*
The var statement declares a list of variables;
 as in function argument lists, the type is last.

A var statement can be at package or function level. 
We see both in this example.
*/
import "fmt"

// var statement
var c, python, java bool

// variables with initializers
var j, k int = 1, 2

func main() {
	var i int
	fmt.Println(i, c, python, java)

	// short variable declarations
	link := "https://www.youtube.com/@1KIINGDOM"

	var apple, banana, orange = true, false, "no!"
	fmt.Println(i, j, k, apple, banana, orange)
	fmt.Println(link)
}
