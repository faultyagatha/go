package main

import (
	"fmt"
	"math"
)

func foo() *int {
	x := 1
	return &x //returns a reference, address
}

//POINTER RECEIVERS
//we can declare methods with pointer receivers
//these methods can modify the value to which the receiver points
type Vertex struct {
	X, Y float64
}

//regular method
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//method with pointer receiver
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//below we use a pointer in a usual manner
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	var y *int
	y = foo()            //de-references the address and gets the val from it
	fmt.Printf("%d", *y) //garbage collector will dealloc the memory for this function in Go
	v := Vertex{3, 4}
	//Go interprets this as (&v).Scale(5) since the Scale method has a pointer receiver.
	v.Scale(10)
	fmt.Println((v.Abs())) //50
	Scale(&v, 10)
	fmt.Println((v.Abs())) //500
}
