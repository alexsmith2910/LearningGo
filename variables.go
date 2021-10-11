package main

import "fmt"

// Declare a single variable.
var a int

// Declare a group of variables.
var (
	b bool
	c float32
	d string
)

// Declare a constant.
const e string = "this is a constant"

func main() {
	a = 42                     // Assign single value
	b, c = true, 32.0          // Assign multiple values
	d = "string"               // Strings must contain double quotes
	fmt.Println(a, b, c, d, e) // 42 true 32 string this is a constant
}
