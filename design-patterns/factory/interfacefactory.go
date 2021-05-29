package main

import "fmt"

/* we can return interface from a factory function */

type Person interface {
	SayHello()
}

//this will be private
type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, I am %d years old", p.name, p.age)
}

func (p *tiredPerson) SayHello() {
	fmt.Printf("My name is %s, I am %d years old. Sorry, I'm too tired to talk to you", p.name, p.age)
}

//NewPerson is a constructor function that
//returns a Person interface
func NewPerson(name string, age int) Person {
	//example of using different object based on conditions
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	p := NewPerson("Stan", 37)
	p.SayHello()
}
