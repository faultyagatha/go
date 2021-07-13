package main

import "fmt"

// ---------------
// Problem we are trying
// to solve: have a Dragon
// extending two structs
//  ---------------
type Bird struct {
	Age int
}

func (b *Bird) Fly() {
	if b.Age >= 10 {
		fmt.Println("Flying")
	}
}

type Lizard struct {
	Age int
}

func (l *Lizard) Crawl() {
	if l.Age < 10 {
		fmt.Println("Crawling")
	}
}

// ---------------
// Canonical way to extend
// multiple structs in Go
// but this can cause problems:
// for example, methods with the
// same name: Age
//  ---------------
type Dragon struct {
	Bird
	Lizard
}

// ---------------
// Workaround methods for
// common Age methods
// but this introduces
// inconsistencies: we can
// set Age from various points
// and do not protect against
// error described in main
//  ---------------
func (d *Dragon) Age() int {
	return d.Bird.Age
}

func (d *Dragon) SetAge(age int) {
	d.Bird.Age = age
	d.Lizard.Age = age
}

func main() {
	d := Dragon{}
	//how to set Age?
	d.Bird.Age = 10
	//what if you set the second to 15?
	//see one workaround SetAge() above
	d.Lizard.Age = 10
}
