package main

import (
	"fmt"
	"time"
)

/*********** Background:
- One core runs 1 thing at a time, how to speed up?
- von Neumann bottleneck: cpu instructions need wait for memory units that slows down the performance
	x = y + z; read y, z from memory; may wait 100+ clock cycles
- power-temperature problem (power wall)
- multi-cores as a solution
- but to make use of cores we need parallelism
	otherwise some of the cores will sit idle while others would work

Parallel execution helps to complete tasks more quickly
processes can cooperate; example: wash dish, dry dish
Parallelism is not necessarily the same a concurrency:
- parallel execution: EXACTLY at the same time
- concurrent execution: start and end time can overlap
- parallel tasks must be executed on different hardware
- concurrent tasks may be executed on the same hardware
*/

/* CONCURRENCY:
- programmer determines which tasks can be executed in parallel
- but mapping of tasks to hardware is out of a programmer's control (think GPU)
Concurrency improves performance, even without parallelism:
- tasks must periodically wait for something
*/

/* PROCESSES and OS:
- instances of running programmes
- have some memory (virtual address space); stack-heap, shared libraries; program counters; data registries...
OS allows many processes to execute concurrently and swithces processes quickly (scheduling):
	OS gives fair access to CPU, memory, etc based on different criteria;

	CONTEXT SWITCH:
	- os sets a timer and a process A
	- when timer is done, os switches to a task B and sets a timer for B

	THREADS:
	- many threads can exist for one process
	- threads have unique context but there is less context than in processes

	INTERLEAVING:
	- order of execution between concurrent tasks is unknown
	- interleaving of instructions between tasks is unknown (is happening on machine level)
	- many interleavings are possible
	- ordering is non-deterministic

	RACE CONDITIONS:
	- outcome depends on non-deterministic interleaving (bad for deterministic programming)
	- races occur due to communication
	- threads are independent but they can share the same data:
	(think of web server, one thread per client but the data can be shared)
*/

/************ CONCURRENCY IS BUILT-IN IN GO:

GOROUTINS:
- a function executing concurrently with other goroutines in the same address space
- like a thread in GO
- many Goroutines execute within a single OS thread

GO RUNTIME SCHEDULER:
- schedules goroutines inside an OS thread
- like a small OS within a single OS thread
- one logical processor mapped to a thread
*/

//one goroutine is created automatically to execute main()
//it is blocking when there are no other goroutines
//when main() ends all other goroutines exit (IMPORTANT!!)

//another goroutine can be created with go keyword
//when there are other goroutines, main() is non-blocking
//goroutine exits silently when it's complete

func main() {
	go fmt.Printf("New routine")
	//this is bad cause we make assumptions about time (non-deterministic)
	time.Sleep(100 * time.Millisecond) //hack to allow "New routine" to be printed
	fmt.Printf("Main routine")         // without the hack above, prints only "Main routine"
}
