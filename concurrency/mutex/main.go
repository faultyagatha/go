package main

import (
	"fmt"
	"sync"
	"time"
)

//sharing variables concurrently can cause problems
//coroutines can interfere with each other

//in this example, there are possible interleavings
var (
	i int = 0
	wg sync.WaitGroup
)

func increm() {
	i = i + 1
	wg.Done()
}

//with mutex, there are no interleavings
//Lock() sets the binary semaphore's flag up
//Unlock() sets the binary semaphore's flag down
var mut sync.Mutex

func incremMutex() {
	mut.Lock()
	i = i + 1
	mut.Unlock()
}

// SafeCounter is safe to use concurrently
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	wg.Add(2)
	//both goroutines share i variable
	//if interleaving happens we can't control it
	go increm()
	go increm()
	wg.Wait()
	fmt.Println(i)

	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))

	//restrict possible interleavings to write into
	//the same variable at the same time = MUTUAL EXCLUSION
}
