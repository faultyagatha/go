package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// ---------------
// Binary serialisation:
// saves all states of an object
// including deep nested structures
//  ---------------

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	//make an encoder
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)
	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	//prepare the memory for the person object
	result := Person{}
	_ = d.Decode(&result)
	return &result
}

type Address struct {
	StreetAddress, City, Country string
}

func main() {
	john := Person{"John",
		&Address{"123 London Rd", "London", "UK"},
		[]string{"Chris", "Matt"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
