package main

import "fmt"

//Interface is a set of method signatures
//used to expess conceptual similarity between types

//Type satisfies an interface if type defines
//all method signatures specified in the interface

type error interface {
	Error() string
}

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
//we don't need to state it anywhere
//compiler does it implicitely
func (t Triangle) Area() float64 {
	return 0.5 * (t.b * t.c)
}

//satisfies Shape2D interface
func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

// ---------------
// In GO, interface has dynamic type and dynamic value
// we can call the methods is inteface has dynamic type
// but no dynamic value
// BUT IF INTERFACE HAS NO DYNAMIC TYPE
// WE CANNOT CALL INTERFACE'S METHODS (we don't know on what type to call it)
// ---------------
type Speaker interface {
	Speak()
}

type Dog struct {
	name string
}

// Dog satisfies Speaker interface
func (d Dog) Speak() {
	fmt.Println(d.name)
}

// use if we don't know what will be stored in the map
// think of Json
var p map[string]interface{}

// ----------
// type assertions on empty interface
// two values can only be compared if one value
// is of the same (or underlying) type with the other
// ---------------
type T interface{}

type hashMap struct {
	m map[T]T
	k []T
}

func (h *hashMap) Less(i, j int) bool {
	switch v := h.m[h.k[i]].(type) {
	case int:
		return v > h.m[h.k[j]].(int)
	case float32:
		return v > h.m[h.k[j]].(float32)
	case float64:
		return v > h.m[h.k[j]].(float64)
	case string:
		return v > h.m[h.k[j]].(string)
	default:
		return false
	}
}

// ----------
// example of how to use interfaces as arguments:
// if Go can successfully determine the type,
// it'll use a needed member function
// ---------------
func maybeLess(i interface{}, j, k int) {
	// runtime check that may fail if the interface if of type hashMap
	if hm, ok := i.(hashMap); ok {
		hm.Less(j, k)
	}

	// compile time check: requires no additional allocation
	// var _ hm = hashMap{}
}

// ---------------
// In GO, interface can be empty
// we can assign anything to a variable or parameter of this type
// but we lose type safety.
// EMPTY INTERFACE SHOULD BE AVOIDED
// ---------------
type empty interface{}

// example of safely using empty interfaces
type monster struct {
	damage int
}

func (m *monster) attack() int {
	return m.damage
}

type attacker interface {
	attack() int
}

type defender interface {
	defend() int
}

func attackOrDefend(attackerDefender interface{}) {
	// Inside this function, we don't know what we're getting, but we can check
	if attacker, ok := attackerDefender.(attacker); ok {
		fmt.Printf("Attacking with damage %d\n", attacker.attack())
	} else if defender, ok := attackerDefender.(defender); ok {
		fmt.Printf("Defending with damage %d\n", defender.defend())
	}
}

func main() {
	var s1 Speaker
	var d1 Dog = Dog{"Brian"}
	//it's legal because the Dog satisfies interface Speaker
	//now s1 has both dynamic type and dynamic value
	s1 = d1
	s1.Speak()
	var d2 *Dog
	//now s2 has dynamic type but no dynamic value (nil)
	//it's legal in GO
	s1 = d2
	//we can do this
	s1.Speak()

	var a attacker = &monster{200}
	attackOrDefend(a)       // Prints "Attacking with damage 200"
	attackOrDefend("Hello") // This is allowed, but does nothing. We can handle it with error.
}
