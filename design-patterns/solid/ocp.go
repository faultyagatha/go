package main

import "fmt"

type Color int

type Size int

const (
	silver Color = iota
	gray
	black
)

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

//FilterByColor filters product by color
func (f *Filter) FilterByColor(
	products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// ---------------
// further filtering options that will inevitably
// come out will start breaking OCP because we will need
// to add more member functions to Filter
// and will be modifying our object
// so the code that was tested before
// will stop to be reliable
//  ---------------

//  ---------------
// here starts a better solution that doesn't break OCP
// if we want to add a new filter, we make a new
// Specification type that conforms the interface
// in this case, we won't modify any structure
// created before
//  ---------------

// Specification tests whether a product satisfies criteria
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type CRPFilter struct{}

func (f *CRPFilter) Filter(
	products []Product, s Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		//check first if the specification is satisfied
		if s.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	laptop := Product{"Laptop", silver, medium}
	ipad := Product{"IPad", silver, medium}
	iphone := Product{"IPhone", silver, small}
	products := []Product{laptop, iphone, ipad}
	//old version of filtering that breaks OCP
	fmt.Printf("Silver Products (old):\n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, silver) {
		fmt.Printf("%s Product is silver\n", v.name)
	}

	//use new Filtering with Specification
	fmt.Printf("Silver Products (new):\n")
	silverSpec := ColorSpecification{silver}
	orpf := CRPFilter{}
	for _, v := range orpf.Filter(products, silverSpec) {
		fmt.Printf("%s Product is silver\n", v.name)
	}
}
