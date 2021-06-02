package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func main() {
	john := Person{"John",
		&Address{"123 London Rd", "London", "UK"}}

	//jane := john

	// ---------------
	// Shallow copy:
	// address pointer is not copyied,
	// it still points to Jonh
	//  ---------------
	//jane.Name = "Jane" //ok
	//jane.Address.StreetAddress = "321 Baker St" //not ok

	//fmt.Println(john.Name, john.Address)
	//fmt.Println(jane.Name, jane. Address)

	// ---------------
	// Deep copy:
	// copyies address pointer
	// but it doesn't scale well:
	// if we have nested objects,
	// all must be copyied manually
	//  ---------------
	jane := john
	jane.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		john.Address.Country}

	jane.Name = "Jane" //ok

	jane.Address.StreetAddress = "321 Baker St" //ok

	fmt.Println(john.Name, john.Address)
	fmt.Println(jane.Name, jane.Address)
}
