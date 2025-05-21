package main

import "fmt"

//Vertex type to demo structs
type Vertex struct {
	X int
	Y int
}

//User type to demo structs
type User struct {
	FirstName string
	LastName  string
	Email     string
	IsAdmin   bool
}

//assign a function to the struct
//and access a struct member (see below)
func (m *User) printFirstName() string {
	return m.FirstName
}

// INITIALISATION: ARRAY OF STRUCTS (I)
func initArrayOfStructs() {
	var users = []User{
		{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@email.com",
			IsAdmin:   true,
		},
		{
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "jane.doe@email.com",
			IsAdmin:   true,
		},
	}
	fmt.Println(users)
}

// INITIALISATION: ARRAY OF STRUCTS (II)

func initArrayOfStructsAnonym() {
	var users = []struct {
		firstName string
		lastName  string
		email     string
		isAdmin   bool
	}{
		{"John", "Doe", "john.doe@email.com", true},
		{"Jane", "Doe", "jane.doe@email.com", true},
	}
	fmt.Println(users)
}

//INITIALISATION: ARRAY OF STRUCTS (III)
//very common practice in Go

var users []User

func init() {
	users = []User{
		{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@email.com",
			IsAdmin:   true,
		},
		{
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "jane.doe@email.com",
			IsAdmin:   true,
		},
	}
}

func main() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2} //can be also initialised using new()
	v.X = 4
	fmt.Println(v)

	user := User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@email.com",
		IsAdmin:   true,
	}
	//shortcut to access a member var
	fmt.Println("The username is : ", user.printFirstName())
}
