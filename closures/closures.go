package main

import (
	"fmt"
	"sync"
)

var db []person = []person{
	{
		0,
		"Cartman",
	},
	{
		1,
		"Stan",
	},
	{
		2,
		"Kyle",
	},
	{
		3,
		"Kenny",
	},
}

type person struct {
	id   int
	name string
}

func getPerson(id int) person {
	for _, v := range db {
		if v.id == id {
			return v
		}
	}
	return person{}
}

func main() {
	// Some ids list somewhere in the code
	ids := make([]int, 10)
	for i := 0; i < 10; i++ {
		ids[i] = i
	}
	ch := make(chan person, 10)
	var wg sync.WaitGroup

	// Put values onto a channel
	for _, id := range ids {
		wg.Add(1)
		// Anonymous function that is immediately called
		go func(id int) {
			defer wg.Done()
			p := getPerson(id)

			ch <- p
		}(id)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	// read values from the channel
	for p := range ch {
		fmt.Printf("%d: %v\n", p.id, p)
	}
}
