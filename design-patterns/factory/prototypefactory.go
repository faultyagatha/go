package main

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	CTO
	Manager
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "developer", 60000}
	case CTO:
		return &Employee{"", "developer", 100000}
	case Manager:
		return &Employee{"", "manager", 100000}
	default:
		panic("unsupported role")
	}
}

func main() {
	m := NewEmployee(CTO)
	m.Name = "Stan"
}
