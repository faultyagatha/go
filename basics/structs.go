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
