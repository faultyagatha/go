package main

import (
	"fmt"

	"github.com/faultyagatha/design-patterns/iterator/internal"
)

func main() {
	// 1. Person iterator
	p := internal.Person{"Alexander", "Graham", "Bell"}
	//custom iterator:
	for it := internal.NewPersonNameIterator(&p); it.MoveNext(); {
		fmt.Println(it.Value())
	}
	/*
		for _, name := range p.Names() {
			fmt.Println(name)
		}
	*/
}
