package main

import "fmt"

type Contacts []Contact

type Contact struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Married bool   `json:"married"`
	Age     int    `json:"age"`
}

func searchContactByID(id int) (Contact, error) {
	return Contact{Name: fmt.Sprintf("Name#%d", id)}, nil
}

func searchAll(limit int) (Contacts, error) {
	var err error
	var con Contact
	var lst Contacts
	for i := 1; i <= limit; i++ {
		if con, err = searchContactByID(i); err != nil {
			fmt.Println("searchContactByID", err)
			continue
		}
		lst = append(lst, con)
	}
	return lst, nil
}

type MyInt int

//this is how in Go we associate methods with data
//note (mi MyInt) - receiver type to be called on
//in the receiver type, there is a hidden arg that is
//passed automatically by value
//in this case, v will be passed
func (mi MyInt) Double() int {
	return int(mi * 2)
}

func main() {
	lst, _ := searchAll(5)

	for _, item := range lst {
		fmt.Printf("%#v\n", item)
	}

	v := MyInt(3)
	fmt.Println(v.Double())
}
