package data

type Point struct {
	x float64
	y float64
}

//expose this function to be used outside (think public)
//to perform init (think constructor)
//but don't give a direct access to the func
//note the pointer in the receiver
//it'll copy a pointer to p, not the p
func (p *Point) Init(xn, yn float64) {
	p.x = xn
	p.y = yn
}

func (p *Point) Scale(v float64) {
	p.x = p.x * v
	p.y = p.y * v
}

func (p *Point) OffsetX(v float64) {
	p.x = p.x + v //more accurate: (*p.x) = p.x + v
}

/* HOW TO USE:

package main

func main() {
	var p data.Point
	p.Init(3, 4)
	p.Scale(2)
	p.OffsetX(5)
}
*/
