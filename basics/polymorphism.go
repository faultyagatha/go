package main

//GO DOESN'T SUPPORT INHERITANCE
//GO DOESN'T SUPPORT OVERRIDING

//GO USES INTERFACES TO ACCOMPLISH POLYMORPHISM
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
func (t Triangle) Area() float64 {
	return 0.5 * (t.b * t.c)
}

//satisfies Shape2D interface
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

//function that makes use of polymorphism
//does the pool's shape fit the yard
func FitInYard(s Shape2D) bool {
	if s.Area() > 100 &&
		s.Perimeter() > 100 {
		return true
	}
	return false
}
