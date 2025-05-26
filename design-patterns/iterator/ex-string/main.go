package main

import (
	"fmt"

	"github.com/faultyagatha/design-patterns/iterator/internal"
)

func main() {
	coll := &internal.StringCollection{}
	coll.Add("Go")
	coll.Add("Rust")
	coll.Add("C++")

	iter := coll.CreateIterator()

	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
}
