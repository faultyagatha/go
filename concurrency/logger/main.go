package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func logMessages1(count int) {
	for i := 0; i < count; i++ {
		log.Printf("1: Logging item #%d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func logMessages2(count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		log.Printf("2: Logging item #%d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	/**
	Output:
	2023/11/24 14:15:26 Logging item #0
	2023/11/24 14:15:27 Logging item #1
	2023/11/24 14:15:28 Logging item #2
	Done
	*/
	go logMessages1(10)

	var wg sync.WaitGroup
	wg.Add(1)
	go logMessages2(10, &wg)

	// logMessages1 will only have 3 seconds to be executed (comment out the wg to see in action)
	// logMessages2 will be executed fully because wg will take care of it
	time.Sleep(3 * time.Second)
	wg.Wait()
	fmt.Println("Done")
}
