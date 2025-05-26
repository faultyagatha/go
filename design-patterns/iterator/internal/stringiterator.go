package internal

// Iterator interface
type Iterator interface {
	HasNext() bool
	Next() string
}

// StringCollection holds a list of strings
type StringCollection struct {
	items []string
}

// Add appends an item to the collection
func (s *StringCollection) Add(item string) {
	s.items = append(s.items, item)
}

// CreateIterator returns a new iterator for the collection
func (s *StringCollection) CreateIterator() Iterator {
	return &StringIterator{
		collection: s,
		index:      0,
	}
}

// StringIterator is a concrete iterator for StringCollection
type StringIterator struct {
	collection *StringCollection
	index      int
}

// HasNext checks if there are more elements
func (i *StringIterator) HasNext() bool {
	return i.index < len(i.collection.items)
}

// Next returns the next element
func (i *StringIterator) Next() string {
	if i.HasNext() {
		item := i.collection.items[i.index]
		i.index++
		return item
	}
	return ""
}
