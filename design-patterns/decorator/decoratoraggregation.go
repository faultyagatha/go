package main

import "fmt"

// ---------------
// Aged defines a contract
// for having the Age
// getter and setter
//  ---------------
type Aged interface {
	Age() int
	SetAge(age int)
}

// ---------------
// Satisfies Aged interface
//  ---------------
type Bird struct {
	age int
}

func (b *Bird) Age() int       { return b.age }
func (b *Bird) SetAge(age int) { b.age = age }

func (b *Bird) Fly() {
	if b.age >= 10 {
		fmt.Println("Flying")
	}
}

// ---------------
// Satisfies Aged interface
//  ---------------
type Lizard struct {
	age int
}

func (l *Lizard) Age() int       { return l.age }
func (l *Lizard) SetAge(age int) { l.age = age }

func (l *Lizard) Crawl() {
	if l.age < 10 {
		fmt.Println("Crawling")
	}
}

// ---------------
// Extend Bird and Lizard
// but keep them private
// to prevent a direct access
// and communicate via
// separate methods
// this is a Decorator
//  ---------------
type Dragon struct {
	bird   Bird
	lizard Lizard
}

func NewDragon() *Dragon {
	return &Dragon{Bird{}, Lizard{}}
}

// ---------------
// consistently set Age
// from two points and
// protect against errors
//  ---------------
func (d *Dragon) Age() int {
	return d.bird.age
}

func (d *Dragon) SetAge(age int) {
	d.bird.SetAge(age)
	d.lizard.SetAge(age)
}

func (d *Dragon) Fly() {
	d.bird.Fly()
}

func (d *Dragon) Crowl() {
	d.lizard.Crawl()
}

func main() {
	d := Dragon{}
	//how to set Age?
	d.SetAge(10)
	d.Fly()
	d.Crowl()
}
