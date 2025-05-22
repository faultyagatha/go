package main

import "fmt"

// ---------------
// Graphic programme for different graphic objects
// without Bridge, can quickly get convoluted:
// RasterCircle, VectorCircle, RasterSquare ...
//  ---------------

// ---------------
// Renderer interface solves the problem above
// using Bridge pattern
// all objects will need to implement
// this interface
// It can result in a cascading set of methods
// on this interface with the introduction
// of new different shapes
//  ---------------
type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius", radius)
}

type RasterRenderer struct {
	Dpi int
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for circle of radius", radius)
}

// ---------------
// Circle has a Bridge
// to the Renderer interface
//  ---------------
type Circle struct {
	renderer Renderer
	radius   float32
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

func main() {
	//raster := RasterRenderer{}
	vector := VectorRenderer{}
	circle := NewCircle(&vector, 5)
	circle.Draw()
}
