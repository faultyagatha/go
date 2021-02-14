package main

import "fmt"

func foo() *int {
	x := 1
	return &x //returns a reference, address
}

func main() {
	var y *int
	y = foo()            //de-references the address and gets the val from it
	fmt.Printf("%d", *y) //garbage collector will dealloc the memory for this function in Go
}
