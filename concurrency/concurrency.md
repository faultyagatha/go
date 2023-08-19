# Concurrency

- built-in in Go --> easier to use
- in other languages, this is not the case

## Benefits

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

## Parallel Execution

- two programs execure in parallel if they execute at exactly the same time
- processor core is usually made to execute 1 instruction at a time
- for parallel execution, we need 2 cores to execute 2 programs in parallel 

## Concurrent Execution

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