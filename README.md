# GO

## File organisation

root dir `go`

`src` and `bin` inside the project

unique project/package name

the first statement in a go source file must be `package packagename`

_____________

## Generics

Uninitialised values are implicitely initialised with `0`, `nil`, or `''` depending on data type.

`Go is pass by value`; similarly to Javascript, for complex datatypes you pass a pointer to avoid memory oveload.

Polymorphism is achieved with `structs` that can have member functions (sort of). It's done externally, in the code (see examples).

Uses `interfaces` (method signatures), same as Typescript. 

`Concurrency` is build in.


## Functional programming

`first-class`; mostly like in Javascript (treated like other types)

`multiple return values` (think Typescript or Python tuples)

```go
func doubleReturn(x int) (int, int) {
	return x, x + 1
}
```
return values can be omitted with `_`:

```go
var sum, _ = doubleReturn(5)
```

variadic functions (`nums ...int`):

think `...rest` operator in Javascript

```go
func spread(nums ...int) int {
  total := 0
  for _, num := range nums {
    total += num
  }
  return total
}
```

closures

`anonymous functions`, like in Javascript.

```go
func factory() func() int {
  i := 0

  return func() int {
    i++
    return i
  }
}
```

## Control-flow

same as usuall, some difference in the variations of `for` when used with `range`

operators can be omitted with `_`

`for range`: a mapper function for various data structures:
two values are returned for each iteration: the first is the `index`, and the second is a `copy of the element at that index`

```go
for _, item := range lst {
		fmt.Printf("%#v\n", item)
	}
```
in `switch`, the break is done implicitely by a compiler


## Data types 

### Arrays

same as in C++: fixed length, must be known on compilation time

but uninit values are by default initialised to `0`, `nil`, or `''`

initialisation:

```go
var theArray [3]string
theArray[0] = "India"  // Assign a value to the first element
theArray[1] = "Canada" // Assign a value to the second element
theArray[2] = "Japan"  // Assign a value to the third element

//using array literal syntax
x := [5]int{10, 20, 30, 40, 50}   // Intialised with values
var y [5]int = [5]int{10, 20, 30} // Partial assignment
```

mostly used with `slices` (Go's own data type) that are not contrained by length

### Slices

like `dynamically-sized arrays` (think Vector in C++)

there is a catch with storing data (see code examples)

every slice has 3 properties:

- `pointer`: start of the slice
- `length`: the number of elements in the slice: len()
- `capacity`: maximum number of elements: cap()

create an empty slice with non-zero length, use `make`

make a slice of empty strings with length 3

```go
s := make([]string, 3)
```

`append` returns a new slice with 1 or more values; immutable, in contrast to Javascript

```go
s = append(s, "d")
s = append(s, "e")
```

`slice` operator similar to Python

```go
s[1:3]
```

initialise values for a slice

```go
t = []string{"go", "is", "cool"}
```

### Maps

hash table

keys are unsorted but we can print map using fmt package in key-sorted order to ease testing

keys can be anything other than string 

```go
m := map(map[string]int)
```

delete values with `delete`

```go
delete(m, "mykey")
```

optional second return value indicates whether the key was present in the object

```go
value, didKeyExist = m["mykey]
```

loop through with `for range` (think of `for in` in Javascript)

```go
fruit := map[string]string{"a": "apple", "b": "banana"}
for key, val := range fruit {
  fmt.Printf("%s -> %s\n", key, val)
}
```

### Strings

```go
// byte index, char (rune)
for i, c := range "go" {
  fmt.Println(i, c)
}
```

### Pointers

like in C++

value is passed as a copy of the underlying value and here we operate on that copy only
```go
func passedByValue(value int) {
  value = 0
}
```
mutate the underlying value by assigning a new int at the referenced address
```go
func passedByReference(reference *int) {
  *reference = 0
}
```

underlying memory addresses:
```go
func main() {
  i := 0
  fmt.Println(i)

  passedByValue(i)
  fmt.Println(i)

  passedByReference(&i)
  fmt.Println(i)
  fmt.Println(&i)
}

```

### Structs

like C structs

methods can be associated with data assigned explicitely 

in this case, a hidden object (struct instance, for example) will be passed by value implicitely - use pointer to avoid value copying (see the code example).


```go
type rect struct {
  width, height int
}


// Pointer receiver type
func (r *rect) area() int {
  return r.width * r.height
}

// Value receiver type
func (r rect) perim() int {
  return 2*r.width + 2*r.height
}
```

Go knows that the method is on the struct since the struct is a parameter. The method is named in the function call defined at the top i.e. `area()`.

Go also magically converts between values and pointers for method calls. So no need for dereference. You can control the behavior by specifying a pointer receiver type to avoid copying the struct on method calls or to allow the method to mutate the underlying values.

```go
type person {
  name string
  age int
}

func NewPerson(name string) *person {
  p := person{name: name}
  p.age = 42
  return &p
}


func main() {
This syntax creates a new struct.

    fmt.Println(person{"Bob", 20})

    fmt.Println(person{name: "Alice", age: 30})

    fmt.Println(person{name: "Fred"})

    fmt.Println(&person{name: "Ann", age: 40})

    fmt.Println(NewPerson("Jon"))

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    sp := &s
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(sp.age)
}
```

It's idiomatic to initiate a new struct with a factory function.

## Error handling

Go communicates errors with explicit return values. Different than exceptions in other languages.

```go
import (
  "errors"
)

func errorfunc(arg int) (int, error) {
  if arg < 0 {
    return -1, errors.new("Cannot use this number")
  }

  return arg + 3, nil
}

func main() {
  if r, e := errorFunc(42); e != nil {
    // FAIL
  } else {
    // OK
  }
}
```

## Concurrency

`Goroutine` is "a lightweight thread of execution"

like a Javascript promise except actually concurrent (since JS is single-threaded)

```go
func f(arg int) int {
  return arg
}

f(4) // called synchronously

go f(5) // called asynchronously

// with an anonymous function
go func(msg string) {
  fmt.Println(msg)
}("GOING ASYNC ANON")
```

`timer`

think of `setTimeout` in Javascript

```go
import "time"

function timeMe() {
  // returns a channel that will be notified at that time (wait two seconds)
  timer1 := time.NewTimer(2 * time.Second)
  timer1.Stop() // cancel timer
  // sleep
  time.Sleep()
}
```

`ticker`

think of `setInterval` in Javascript

```go
package main
import (
    "fmt"
    "time"
)

func main() {

    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)
    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
}

```

### Other

`panic`

like `throw` in Javascript but it will throw a non-zero exit code and provide a stack trace to stderr.

```go
package main

import "os"

func main() {
	panic("a problem")

	_, err := os.Create("tmp/file")
	if err != nil {
		panic(err)
	}
}
```

`defer`

Like a `finally` in Javascript. Except you defer a function call

You have to check for errors even in a deferred function

example: `defer` the cleanup of a file

```go
func main() {
  f := createFile("/tmp/defer.txt")
  defer closeFile(f)
  writeFile(f)
}
```

`Exit`

`os.Exit(3)` to exit with an explicit exit code. 