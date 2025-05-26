package internal

type Person struct {
	FirstName, MiddleName, LastName string
}

// ---------------
// NamesSlice is a simple
// iterable that iterates
// over Person
//
//	---------------
func (p *Person) NamesSlice() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- p.FirstName
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.LastName
	}()
	return out
}

// ---------------
// NamesGenerator is a generator
// that can iterate over Person
//  ---------------

func (p *Person) Names() []string {
	return []string{p.FirstName, p.MiddleName, p.LastName}
}

// ---------------
// PersonNamesIterator is a custom
// iterator that can
// iterate over Person
//  ---------------

type PersonNamesIterator struct {
	person  *Person
	current int
}

// factory function that initialises PersonNamesIterator
func NewPersonNameIterator(person *Person) *PersonNamesIterator {
	return &PersonNamesIterator{person, -1}
}

// function that returns the value
func (p *PersonNamesIterator) MoveNext() bool {
	p.current++
	return p.current < 3
}

// function that checks if there is next value
func (p *PersonNamesIterator) Value() string {
	switch p.current {
	case 0:
		return p.person.FirstName
	case 1:
		return p.person.MiddleName
	case 2:
		return p.person.LastName
	}
	panic("We should not get here")
}
