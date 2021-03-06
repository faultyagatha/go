package main

import (
	"fmt"
)

//type assertion syntax: t := i.(T)
//t, ok := i.(T)
//if i does not hold a T, the statement will trigger a panic.

type Shape2D interface {
	Area() float64
	Perimeter() float64
}

type Triangle struct {
	a float64
	b float64
	c float64
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.b * t.c)
}

func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

type Rectangle struct {
	l float64
	w float64
}

//satisfies Shape2D interface
func (r Rectangle) Area() float64 {
	return r.l * r.w
}

//satisfies Shape2D interface
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.l + r.w)
}

func DrawShapeWithCustomTypeAssertion(s Shape2D) {
	//if interface has a concrete type
	//ok == true
	//if doesn't
	//ok == false
	rect, ok := s.(Rectangle)
	if ok {
		//do something
		fmt.Println(rect)
	}
	tri, ok := s.(Triangle)
	if ok {
		//do something
		fmt.Println(tri)
	}
}

func DrawShapeWithTypeSwich(s Shape2D) {
	//sh will be a concrete type
	//that s represents
	switch sh := s.(type) {
	case Rectangle:
		//do something
		fmt.Println(sh)
	case Triangle:
		//do something
		fmt.Println(sh)
	}
}

//Stringer is defined by fmt package
// type Stringer interface {
// 	String() string
// }

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	//reading from json
	var p map[string]interface{}
	// err = json.Unmarshal(data, &p)

	//using type switch
	for k, v := range p {
		switch c := v.(type) {
		case string:
			fmt.Printf("Item %q is a string, containing %q\n", k, c)
		case float64:
			fmt.Printf("Looks like item %q is a number, specifically %f\n", k, c)
		default:
			fmt.Printf("Not sure what type item %q is, but I think it might be %T\n", k, c)
		}
	}
}
