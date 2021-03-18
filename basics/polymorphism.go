package main

import "fmt"

//GO DOESN'T SUPPORT INHERITANCE
//GO DOESN'T SUPPORT OVERRIDING

//GO USES INTERFACES TO ACCOMPLISH POLYMORPHISM

////// EXAMPLE 1
type Shape2D interface {
	Area() float64
	Perimeter() float64
}

type Triangle struct {
	a float64
	b float64
	c float64
}

//satisfies Shape2D interface
func (t Triangle) Area() float64 {
	return 0.5 * (t.b * t.c)
}

//satisfies Shape2D interface
func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

type Rectangle struct {
	l float64
	w float64
}

//satisfies Shape2D interface
func (r Rectangle) Area() float64 {
	return r.l * r.w
}

//satisfies Shape2D interface
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.l + r.w)
}

//function that makes use of polymorphism
//does the pool's shape fit the yard
func FitInYard(s Shape2D) bool {
	if s.Area() > 100 &&
		s.Perimeter() > 100 {
		return true
	}
	return false
}

////// EXAMPLE 2
type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	numHours    int
	hourlyRate  int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.numHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

// using Polymorphism for calculation based
// on the array of variables of interface type
func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Println("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organisation = $%d", netincome)
}

func main() {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", numHours: 160, hourlyRate: 25}
	incomeStreams := []Income{project1, project2, project3}
	calculateNetIncome(incomeStreams)
}
