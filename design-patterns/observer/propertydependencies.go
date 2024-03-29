package main

import (
	"container/list"
	"fmt"
)

type Observable struct {
	subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Observer interface {
	Notify(data interface{})
}

type Person struct {
	Observable
	age int
}

func NewPerson(age int) *Person {
	return &Person{Observable{new(list.List)}, age}
}

type PropertyChanged struct {
	Name  string
	Value interface{}
}

func (p *Person) Age() int { return p.age }

// ---------------
// SetAge is a single source of truth
// for sending notification
// about property changes.
// It may quickly get convoluted if we
// have many property-dependent cases.
// Instead, a good practice is to build
// a map with all dependencies.
//  ---------------
func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	} // no change

	//get the old value
	oldCanVote := p.CanVote()

	p.age = age
	p.Fire(PropertyChanged{"Age", p.age})

	//fire only if the property is new
	if oldCanVote != p.CanVote() {
		p.Fire(PropertyChanged{"CanVote", p.CanVote()})
	}
}

// ---------------
// CanVote is a method
// that depends on the age field
// But we can set notifications
// only in SetAge()
//  ---------------
func (p *Person) CanVote() bool {
	return p.age >= 18
}

type ElectoralRole struct {
}

func (e *ElectoralRole) Notify(data interface{}) {
	if pc, ok := data.(PropertyChanged); ok {
		if pc.Name == "CanVote" && pc.Value.(bool) {
			fmt.Println("Congratulations, you can vote!")
		}
	}
}

func main() {
	p := NewPerson(0)
	er := &ElectoralRole{}
	p.Subscribe(er)

	for i := 10; i < 20; i++ {
		fmt.Println("Setting age to", i)
		p.SetAge(i)
	}
}
