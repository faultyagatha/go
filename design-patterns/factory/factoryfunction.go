package main

type Person struct {
	Name    string
	Age     int
	NumLegs int //assume a person with 2 legs
}

func NewPerson(name string, age int) *Person {
	//we can add additional logic here
	return &Person{name, age, 2}
}

func main() {
	p := NewPerson("John", 33)
	//if something happened to John
	p.NumLegs = 1
}
