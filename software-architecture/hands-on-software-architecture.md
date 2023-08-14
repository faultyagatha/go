# Hands-on Software Architecture with Go [Notes]

[Introduction](#introduction)

## Introduction

### Components

package code into components based on:
- high cohesion (component performs a single related task)
- low coupling (components should have less dependency between themselves)

interactions and contracts between components become API-driven
`services` are components that don't share memory and can only communicate via network calls

### Concurrency

> Traditional approach to multithreading is the `shared memory and locks` paradigm for communication. 

Shared memory leads to:
- deadlocks and corruption if a process crashes or behaves not as expected
- difficult to recover from failure

> `Communicating Sequential Processes` is a formal language describing patterns of interaction in concurrent systems.

- key concept is a `process`
- code inside the process in sequential
- at some point in time that code can start another process
- these processes communicate with each other
- CSP promotes the `message-passing` paradigm of comminication -->

> Go uses the `message-passing` paradigm of comminication using the concept of channels. 

`Channels` are queues with a logical interface of `send()` and `recv()`
- Go channels can be blocking
- passing a pointer between channels is idiomatic in Go
- procedures are called `goroutines`: one can spawn a goroutine out of the function and have it execute independently

### Polymorphism

typically is implemented as either:
- tables for all method calls prepared statically on compile time (C++)
- a method lookup at each call (JS, Python)

Go implements method tables but computes them at runtime. 
Interfaces are represented as a pointer-pair: 
- *p1: information about the type and method tables (i-table)
- *p2: references the associated data

### Modules

`All code in the package should be private unless explicitly needed by other client packages.`

[Convention over configuration](https://en.wikipedia.org/wiki/Convention_over_configuration)
