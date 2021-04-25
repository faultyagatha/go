package main

import (
	"fmt"
	"math"
)

func foo() *int {
	x := 1
	return &x //returns a reference, address
}

func addOneByCopy(x int) {
	x += 1
}

func addOne(x *int) {
	//dereference the pointer before adding
	*x += 1
}

//******** POINTER RECEIVERS
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
	//go implicitly dereferences the pointer
	//in C: v->v.X
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	var y *int           //pointer to an int
	y = foo()            //de-references the address and gets the val from it
	fmt.Printf("%d", *y) //garbage collector will dealloc the memory for this function in Go
	v := Vertex{3, 4}
	//Go interprets this as (&v).Scale(5) since the Scale method has a pointer receiver.
	v.Scale(10)
	fmt.Println((v.Abs())) //50
	Scale(&v, 10)
	fmt.Println((v.Abs())) //500

	x := 2
	var xPtr *int = &x

	addOneByCopy(x)
	fmt.Println("x after addOneByCopy: ", x) //2
	addOne(xPtr)
	fmt.Println("x after addOne: ", x) //3
}
