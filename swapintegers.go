package main

import "fmt"

func swap(x *int, y *int) {
	// swaps the int values in x and y

	temp := new(int) // allocates memory for temp
	*temp = *x       // deferenece temp to access the value pointer points to
	*x = *y          // deference x to "put" the derefenced value of y into it
	*y = *temp       // same as above
}

func main() {
	x := 1
	y := 2

	swap(&x, &y)

	fmt.Println(x, "This is x")
	fmt.Println(y, "This is y")
}
