package main

import (
	"fmt"
	"sync"
)

/*
* Dining Philosophers problem
* http://en.wikipedia.org/wiki/Dining_philosophers_problem
- each chopstick is a mutex
- each philosopher is assiciated with a goroutine and two chopsticks
*/

// //ChopStick represents a chopstick
// type ChopStick struct {
// 	sync.Mutex
// }

// //Philosopher represents a philosopher
// type Philosopher struct {
// 	id                    int
// 	leftChopS, rightChopS *ChopStick
// }

// func (p Philosopher) eat() {
// 	for {
// 		p.leftChopS.Lock()
// 		p.rightChopS.Lock()

// 		fmt.Println(p.id, "eating")

// 		p.rightChopS.Unlock()
// 		p.leftChopS.Unlock()
// 	}
// }

// func main() {
// 	count := 5

// 	cSticks := make([]*ChopStick, count)
// 	for i := 0; i < count; i++ {
// 		cSticks[i] = new(ChopStick)
// 	}
// 	philos := make([]*Philosopher, count)
// 	for i := 0; i < count; i++ {
// 		//left chopstick = philosopher's index
// 		//right chopstick = next after left
// 		//this will cause: fatal error: all goroutines are asleep - deadlock!
// 		//philos[i] = &Philosopher{i, cSticks[i], cSticks[(i+1)%count]}

// 		//each philosopher picks up lowest numbered chopstick first: (i-1+count)%count
// 		philos[i] = &Philosopher{i, cSticks[i], cSticks[(i+1)%count]}

// 		go philos[i].eat()
// 	}
// 	// for i := 0; i < count; i++ {
// 	// 	go philos[i].eat()
// 	// }
// 	// wait endlessly while they're dining
// 	endless := make(chan int)
// 	<-endless
// 	fmt.Println("Everybody finished eating")
// }

func philos(id int, left, right chan bool, wg *sync.WaitGroup) {
	fmt.Printf("Philosopher # %d wants to eat\n", id)
	<-left
	<-right
	left <- true
	right <- true
	fmt.Printf("Philosopher # %d finished eating\n", id)
	wg.Done()
}

func main() {
	const numPhilos = 5
	var forks [numPhilos]chan bool
	for i := 0; i < numPhilos; i++ {
		forks[i] = make(chan bool, 1)
		forks[i] <- true
	}
	var wg sync.WaitGroup
	for i := 0; i < numPhilos; i++ {
		wg.Add(1)
		go philos(i, forks[(i-1+numPhilos)%numPhilos], forks[(i+numPhilos)%numPhilos], &wg)
	}
	wg.Wait()
	fmt.Println("Everybody finished eating")
}
