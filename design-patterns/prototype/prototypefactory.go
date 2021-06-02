package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// ---------------
// Same deep copy with binary serialisation:
// saves all states of an object
// including deep nested structures
//  ---------------

type Address struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office Address
}

func (p *Employee) DeepCopy() *Employee {
	// note: no error handling below
	b := bytes.Buffer{}
	//make an encoder
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	// peek into structure
	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	//prepare the memory for the person object
	result := Employee{}
	_ = d.Decode(&result)
	return &result
}

// ---------------
// Employee factory:
// either a struct or some functions
//  ---------------
var mainOffice = Employee{
	"", Address{0, "123 East Dr", "London"}}
var auxOffice = Employee{
	"", Address{0, "66 West Dr", "London"}}

// ---------------
// Utility method for configuring employee
// based on a particular prototype
//  ---------------
func newEmployee(proto *Employee,
	name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

// ---------------
// Factory functions
// based on a particular prototype
//  ---------------
func NewMainOfficeEmployee(
	name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(
	name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}

func main() {
	//john := Employee{"John",
	//  Address{100, "123 East Dr", "London"}}
	//
	//jane := john.DeepCopy()
	//jane.Name = "Jane"
	//jane.Office.Suite = 200
	//jane.Office.StreetAddress = "66 West Dr"

	john := NewMainOfficeEmployee("John", 100)
	jane := NewAuxOfficeEmployee("Jane", 200)

	fmt.Println(john)
	fmt.Println(jane)
}
