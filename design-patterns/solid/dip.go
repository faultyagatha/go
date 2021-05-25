package main

import "fmt"

// ---------------
// Dependency Inversion Principle
// HLM (high-level module) should not
// depend on LLM (low-level module)
// Both should depend on abstractions
//  ---------------

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
	// other useful stuff here
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

//LLM
type Relationships struct {
	relations []Info
}

func (rs *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range rs.relations {
		if v.relationship == Parent &&
			v.from.name == name {
			result = append(result, rs.relations[i].to)
		}
	}

	return result
}

func (rs *Relationships) AddParentAndChild(parent, child *Person) {
	rs.relations = append(rs.relations,
		Info{parent, Parent, child})
	rs.relations = append(rs.relations,
		Info{child, Child, parent})
}

// ---------------
// the implementation below
// breaks DIP because Research (HLM)
// uses the internals of Relationships (LLM)
//  ---------------

/*
type Research struct {
	relationships Relationships //breaks DIP
}

func (r *Research) Investigate() {
	relations := r.relationships.relations
	for _, rel := range relations {
		if rel.from.name == "John" &&
			rel.relationship == Parent {
			fmt.Println("John has a child called", rel.to.name)
		}
	}
}
*/

// ---------------
// better implementation
// that doesn't break DIP
//  ---------------

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

//HLM
type Research struct {
	browser RelationshipBrowser // low-level
}

func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	// low-level module
	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	research := Research{&relationships}
	research.Investigate()
}
