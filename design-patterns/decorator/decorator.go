package main

import "fmt"

type Shape interface {
	Render() string
}

// ---------------
// Problem we are trying
// to solve: have a struct
// extending Circle and Square
//  ---------------
type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f",
		c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side %f", s.Side)
}

// ---------------
// This could work
// but it is not
// generic enough
//  ---------------
type ColoredSquare struct {
	Square
	Color string
}

// ---------------
// This is much better:
// this is our decorator
//  ---------------
type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s",
		c.Shape.Render(), c.Color)
}

type TransparentShape struct {
	Shape        Shape
	Transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has %f%% transparency",
		t.Shape.Render(), t.Transparency*100.0)
}

func main() {
	circle := Circle{2}
	//render comes from the shape interface
	fmt.Println(circle.Render())

	//render comes from the shape interface
	//but with the decorated part we can now add colour
	//limitations: circle.Resize() cannot be called
	//with redCircle
	redCircle := ColoredShape{&circle, "Red"}
	fmt.Println(redCircle.Render())

	rhsCircle := TransparentShape{&redCircle, 0.5}
	fmt.Println(rhsCircle.Render())
}
