package main

import "fmt"

func main() {
	lst, _ := searchAll(5)

	for _, item := range lst {
		fmt.Printf("%#v\n", item)
	}
}

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
