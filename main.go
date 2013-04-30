package main

import "fmt"

// this is a comment
var lol string = "lol"

func main() {
	var x string = "This is a string!"
	fmt.Println(x)
	//for_loop()
	//array_ex()
	//slice_ex()
	maps_ex()
}






func maps_ex() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	delete(x,1)
}


func slice_ex() {
	arr := []float64{1,2,3,4,5}
	x := arr[1:4]
	fmt.Println(x)
}


func array_ex() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for _, value := range x{
		total += value
	}

	fmt.Println(total / float64(len(x)))
}

func for_loop(){

	for i := 1; i <= 10; i++ {
		switch i {
			case 0: fmt.Println("Zero")
			case 1: fmt.Println("One")
			case 2: fmt.Println("Two")
			case 3: fmt.Println("Three")
			case 4: fmt.Println("Four")
			case 5: fmt.Println("Five")
			default: fmt.Println("Unknown Number")
		}
	}
}
