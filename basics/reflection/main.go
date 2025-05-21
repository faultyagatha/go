package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string "namestring"
	age int
}

func ShowTag(i interface{}) {
	switch t := reflect.TypeOf(i); t.Kind() {
	case reflect.Ptr:
		// Elem returns a type’s element type. It panics if the type’s Kind is not Array, Chan, Map, Ptr, or Slice
		// The struct StructField has a Tag member which returns the tag-name as a string. 

		tag := t.Elem().Field(0).Tag // Field(0).Tag gives namestr.
		fmt.Printf("Tag: %s\n", tag)
	}
}

// must be called with *Person
func Set(i interface{}) {
	switch i.(type) {
	case *Person:
			r := reflect.ValueOf(i)
			r.Elem().Field(0).SetString("Ada Lovelace")
	}
}

func main() {
	var p Person = Person{"Alan Turing", 19}
	ShowTag(p)
	Set(p)
	fmt.Println("Finished.\n")
}