package main

import "fmt"

// sorts only integers
// but what if we need to sort strings too?
// use
func bubblesort(n []int) {
	for i := 0; i < len(n)-1; i++ {
			for j := i + 1; j < len(n); j++ {
					if n[j] < n[i] {
							n[i], n[j] = n[j], n[i]
					}
			}
	}
}

// instead of many functions for each data structure,
// create an interface that would be used on these data structures
type Sorter interface {
	Len() int           // len() as a method.
	Less(i, j int) bool // p[j] < p[i] as a method.
	Swap(i, j int)      // p[i], p[j] = p[j], p[i] as a method.
}

// create new types that will be implementing
// the sorter interface
type Sortablei []int
type Sortables []string

// implement functions from the sorter interface
// this is important to be able to use the sorter 
func (p Sortablei) Len() int               {return len(p)}
func (p Sortablei) Less(i int, j int) bool {return p[j] < p[i]}
func (p Sortablei) Swap(i int, j int)      {p[i], p[j] = p[j], p[i]}

func (p Sortables) Len() int               {return len(p)}
func (p Sortables) Less(i int, j int) bool {return p[j] < p[i]}
func (p Sortables) Swap(i int, j int)      {p[i], p[j] = p[j], p[i]}

// generic function that works on sorter interface
func Sort(x Sorter) { 
	// both string and int now have Len method
	for i := 0; i < x.Len() - 1; i++ { 
		for j := i + 1; j < x.Len(); j++ {
			// both string and int now have Less and Swap methods
			if x.Less(i, j) {
					x.Swap(i, j)
			}
		}
	}
}

func main() { 
	ints := Sortablei{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
	strings := Sortables{"nut", "ape", "elephant", "zoo", "go"}
	
	Sort(ints)
	fmt.Printf("%v\n", ints)
	Sort(strings)
	fmt.Printf("%v\n", strings)
}