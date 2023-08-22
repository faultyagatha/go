# Concurrency

[Concurrency and Parallelism](#concurrency-and-parallelism)
[Goroutines](#goroutines)

- built-in in Go --> easier to use
- in other languages, this is not the case

## Concurrency and Parallelism

### Benefits 

- tasks may complete quicker
- example: two piles of dishes will be washed faster by 2 dish-washers but we could also speed this task by using sequential wash (e.g., wash dish, dry dish)
- some things can be executed sequentially --> this is also quicker but this is not parallel execution
- but some things cannot be parallelised 
- speedup may be achieved by designing faster processors

> `Van Neumann Bottleneck` is the delay to access memory. Cache partly solves this problem.

> Moore predicted that transistor density would double every two years. `Smaller transistors switch faster` --> Exponential increase in density would lead to exponential increase in speed.

Increasing transistors use less power but density scaling is much faster --> increased power consumption --> high power leads to higher temperatures. -->

Moore's Law doesn't work anymore becase of the `power wall`.

> `Dennard Scaling`: voltage should scale with transistor size to keep power consumption and temperature law. But voltage must stay above threshold voltage. Pluse noise problem occurs. --> 

We cannot increase much clock frequency because of Dennard Scaling but we can increase density (the N of cores). -->

We need parallel execution to exploit multi-core systems.

### Parallel Execution

- two programs execure in parallel if they execute at exactly the same time
- processor core is usually made to execute 1 instruction at a time
- for parallel execution, we need 2 cores to execute 2 programs in parallel 

### Concurrent Execution

- concurrent execution is not necessarily the same as parallel execution 
- concurrent: 
  - start and end times overlap 
  - may or not be executed on different hardware
  - only 1 task is actually executed at a time
- parallel
  - 2 tasks are literally executed at the same time
  - must be executed on different hardware

Programmer determines which tasks CAN be executed in parallel. What WILL be executed in parallel depens on how the tasks are mapped to hardware. 

Even in 1-core computer, concurrency may be an improvement because of `hidden latency`:
- tasks must periodically wait for something (e.g., memory, io response, hardware communication)

```go
x := y + z // where y and z are read from memory

// the instruction is simple to be executed in 1 clock cycle
// but we may be waiting 100+ clock cycles just to get the data

// we can hide this latency to do something instead of waiting 
```

### Process

- the concept of concurrency comes from OS
- process is an instance of a running program

#### Every Process has:

1. Memory 
  - virtual address space
  - code, stack, heap, shared libs

2. Registers:
  - program counter
  - data registers
  - stack ptr ...

#### OS

- allows many processes to execute concurrently 
- processes are switched quickly (handled by OS)
- scheduling task is the task to decide which process runs when
- user has an impression of parallelism
- OS must give processes fair access to resources

### Scheduling

- OS schedules processes for execution
- this gives an illusion of parallel execution
- there are different algorithms for scheduling based on task priorities 
- `context switch` is a change of control flow from one process to another
- OS must save the state (`context`) of the process before it does the switch 
- when the process goes back to the process OS restores the context

### Threads vs Processes

> `Thread` is a lightweight process (it has less context than a process) and shares some context with other processes. 

- multiple threads can exist on 1 process
- threads have unique context and shared context
- when switching from 1 thread to 2 thread in 1 process, context switch happens much faster

## Goroutines

> `Goroutine` is a thread in Go

- many Goroutines execute within a single OS thread
- scheduling process is done by go `runtime scheduler`

### Go Runtime Scheduler

- schedules goroutines inside an OS thread
- like a little goroutines OS inside a single OS thread
- `logical processor` is `mapped to a thread` --> there is `no parallel processing` 
- but we can have several logical processors and allow to mock parallel processing

### Interleavings

- debugging concurrent programming is hard
- we don't know at which state we've been in the execution state in different tasks
- the order of the task execution in different concurrent tasks is unknown
- interleavings of instructions between tasks is unknown
- many interleavings are possible
- interleavings happen at the machine level code

### Race Conditions 

> `Race Condition` happens when outcome of the program depends on non-deterministic ordering

- race conditions must be avoided
- they `occur due to communication`

### Task Communication

- threads are largely independent but not completely independent
- example 1: web server (one thread per client but they may share the data)
- example 2: image processing (one thread per pixel block but they may share the pixel values of the neighbours)

### Creating a Goroutine

- one goroutine is created automatically to execute `main()`
- if there are no other goroutines:
  - `main goroutine is blocking`
  - `when main() ends all other goroutines exit` (IMPORTANT!!) `even if go routines are not finished`

- another goroutine can be created with `go keyword`
- `when there are other goroutines, main() is non-blocking`
- goroutine exits silently when it's complete


```go
// --- BLOCKING MAIN:

func foo() {
  fmt.Printf("Foo function")
}

// now main is blocking
func main() {
  fmt.Printf("Main routine")
  a := 1
  foo()
  // the code here will be executed
  // ONLY AFTER foo() finishes execution
  a = 1
}
```

```go
// --- NON-BLOCKING MAIN:

func foo() {
  fmt.Printf("Foo routine")
}

// now main is non-blocking
// because we introduced another goroutine
func main() {
  fmt.Printf("Main routine")
  a := 1
  go foo()
  // the code here will be executed
  // BEFORE or AFTER or IN-BETWEEN the
  // execution of the foo() goroutine
  a = 1
}
```

- common bad-bad hack: get the main routine sleep for some time:
`time.Sleep(200 * time.Millisecond)`
- it's bad because it makes the assumption about the time necessary to execute other routines 

--> need to have `formal synchronisation constructs`

### Synchronisation

> `Synchronisation` is using global events whose execution is viewed by all threads, simultaneously

- synchronisation is used to restrict ordering and bad interleavings
- to make synchronisation possible, we need some global event that is visible to all threads at the same moment
- we can introduce condition execution that follows this global event

- synchronisation reduces efficiency because we're limiting scheduling --> 
- it reduces performance but it is absolutely necessary to restric bad interleavings and get deterministic results

```go
// pseudo code
// -- TASK 1
x := 1
x = x + 1

[GLOBAL EVENT]

if GLOBAL EVENT {
  print x
}
```

#### WaitGroups

- `sync package` contains functions to synchronise between goroutines
- sync.WaitGroups forces a goroutine to wait for other goroutines
- contains an internal counter (waiting semaphore)
  - increment counter for each goroutine to wait for 
  - decrement counter when each goroutine completes
  - waiting goroutine cannot continue until counter is 0

- sync.WaitGroups ethods:
  - `wg.Add()` increments the counter
  - `wg.Done()` decrements the counter
  - `wg.Wait()` blocks until the counter == 0; wait is then be passed to main to notify that main can continue

### Communication

- goroutines usually work together to perform a bigger task
- they often need to send data to collaborate

----
EXAMPLE: find the product of 4 int
  - make 2 goroutines, each multiplies a pair
  - main goroutine multiplies the 2 results

--> 
- need to send ints from main routine to the two sub-routines
- need to send results from sub-routines back to main routine
- naive implementation:

```go
package main

import (
	"fmt"
	"sync"
)

func mult1(a int, b int, res *int, wg *sync.WaitGroup) {
	fmt.Printf("mult1 routine\n")
	*res = a * b
	wg.Done()
}

func mult2(a int, b int, res *int, wg *sync.WaitGroup) {
	fmt.Printf("mult1 routine\n")
	*res = a * b
	wg.Done()
}

func main() {
	var res1 int
	var res2 int
	var wg sync.WaitGroup
	wg.Add(1)
	go mult1(2, 3, &res1, &wg)
	wg.Add(1)
	go mult2(3, 3, &res2, &wg)
	wg.Wait()
	fmt.Printf("Main routine %d\n", res1)
	fmt.Printf("Main routine %d\n", res2)
	res3 := res1 * res2
	fmt.Printf("Main routine %d\n", res3)
}
```
- but the example above is simple, real-live cases are complicated
- we need a way of comminicating between goroutines
----

#### Channels

- `channels` transfer data between goroutines
- channels are types to transfer types data
- `make()` creates a channel
- send and receive data using `<-` operator

```go
// create a channel
c := make(chan int)

// send data on a channel
c <- 3
// receive data from a channel
x := <- c
```