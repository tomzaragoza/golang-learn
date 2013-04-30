package main

import "fmt"


func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func square(x *float64){
	*x = *x * *x
}

func main() {
	x := 5
	fmt.Println(x, "This is the value of x")
	zero(&x)
	fmt.Println(x, "This is the value of x after calling zero")
	// -----

	// new takes a type as an argument, allocates enough memory to 
	// fit a value of that type and returns a pointer to it.
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1

	// ------
	j := 1.5
	square(&j)
	fmt.Println(j)
}