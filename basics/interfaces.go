package main

import "fmt"

//Interface is a set of method signatures
//used to expess conceptual similarity between types

//Type satisfies an interface if type defines
//all method signatures specified in the interface

type error interface {
	Error() string
}

type Shape2D interface {
	Area() float64
	Perimeter() float64
}

type Triangle struct {
	a float64
	b float64
	c float64
}

//satisfies Shape2D interface
//we don't need to state it anywhere
//compiler does it implicitely
func (t Triangle) Area() float64 {
	return 0.5 * (t.b * t.c)
}

//satisfies Shape2D interface
func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

//In GO, interface has dynamic type and dynamic value
//we can call the methods is inteface has dynamic type
//but no dynamic value
//BUT IF INTERFACE HAS NO DYNAMIC TYPE
//WE CANNOT CALL INTERFACE'S METHODS (we don't know on what type to call it)
type Speaker interface {
	Speak()
}

type Dog struct {
	name string
}

//Dog satisfies Speaker interface
func (d Dog) Speak() {
	fmt.Println(d.name)
}

func main() {
	var s1 Speaker
	var d1 Dog = Dog{"Brian"}
	//it's legal because the Dog satisfies interface Speaker
	//now s1 has both dynamic type and dynamic value
	s1 = d1
	s1.Speak()
	var d2 *Dog
	//now s2 has dynamic type but no dynamic value (nil)
	//it's legal in GO
	s1 = d2
	//we can do this
	s1.Speak()
}
