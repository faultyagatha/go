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

// ---------------
// Observer that knows
// how to send notification
//  ---------------
type Observer interface {
	Notify(data interface{})
}

// ---------------
// PropertyChange can
// communicate with struct's properties
// access it's old value and new value.
// The setter and getter are below
//  ---------------
type PropertyChange struct {
	Name     string      //property string
	NewValue interface{} //new value for the property
}

type Person struct {
	Observable
	age int //private property we want to get notified if changed
}

func NewPerson(age int) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		age:        age,
	}
}

//Age is a getter for age member
func (p *Person) Age() int { return p.age }

//SetAge is a setter for age member
func (p *Person) SetAge(newAge int) {
	if newAge == p.age {
		return
	}
	p.age = newAge
	//notify about changing in age
	p.Fire(PropertyChange{"Age", p.age})
}

// ---------------
// Real-life Example:
// TrafficManagement notification
//  ---------------
type TrafficManagement struct {
	o Observable
}

func (t *TrafficManagement) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.NewValue.(int) >= 16 {
			fmt.Println("Congrats, you can drive now!")
			// we no longer care
			t.o.Unsubscribe(t)
		}
	}
}

func main() {
	p := NewPerson(15)
	t := &TrafficManagement{p.Observable}
	p.Subscribe(t)

	for i := 16; i <= 20; i++ {
		fmt.Println("Setting age to", i)
		p.SetAge(i)
	}
}
