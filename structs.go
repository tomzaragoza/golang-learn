package main

import ("fmt"; "math")

type Circle struct {
	x, y, r float64
}


func main() {
	var c Circle
	// or...use the new() function
	//circ := new(Circle)
	// or.. give fields (can leave out fields, too)
	x := Circle{0, 0, 5}
}