package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// ---------------
// Functional factory
//  ---------------
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// ---------------
// Struct factory
//  ---------------
type EmployeeFactory struct {
	position     string
	annualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.position, f.annualIncome}
}

func NewEmployeeStructFactory(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func main() {
	//use the functional factory
	developerFactory := NewEmployeeFactory("developer", 60000)
	managerFactory := NewEmployeeFactory("manager", 80000)

	developer := developerFactory("Sara")
	manager := managerFactory("Kora")
	fmt.Println(developer)
	fmt.Println(manager)

	//use the struct factory
	ctoFactory := NewEmployeeStructFactory("CTO", 100000)
	cto := ctoFactory.Create("Mona")
	fmt.Println(cto)
}
