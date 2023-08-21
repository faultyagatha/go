# Tips. Tricks, Interview Questions

## How to swap values

simple: a,b, = b,a

```go
func swap(a, b string) (string, string) {
  return b, a
}
```

## Structs and pointers
```go
type Vertex struct {
	X int
	Y int
}

//returns a copy of a struct
func funcOne() Vertex {
  return Vertex{X: 1}
}

//returns a pointer to the struct
func funcTwo() *Vertex {
  return &Vertex{}
}

//overrides a value in a struct passed to a func
func funcThree(v *Vertex) {
  v.X = 1
}
```
## Concat strings
```go
import (
    "strings"
    "fmt"
)

func main() {
  var str strings.Builder
  for i := 0; i < 10; i++ {
    str.WriteString("hello")
  }
}
```

## Check if map contains a key

```go
if val, ok := dict["someval"]; ok {
  //do something
}
```

```go
type Node struct {
  Next  *Node
  Value interface{}
}

var first *Node

visited := make(map[*Node]bool)
for n := first; n != nil; n = n.Next {
  if visited[n] {
    fmt.Println("cycle detected")
    break
  }

  visited[n] = true
  fmt.Println(n.Value)
}
```

5. Copy map:
```go
mapOne := map[string]bool{"A": true, "B": true}
mapTwo := make(map[string]bool)
for i, v:= range mapOne {
  mapTwo[i] = v
}
```

## Reverse a slice of ints

```go
func reverse(s []int) {
  for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
    s[i], s[j] = s[j], s[i]
  }
}
```

## Print const (iotas)

```go
type State int

//integers under the hood
const (
	OffHook State = iota
	Connecting
	Connected
	OnHold
	OnHook
)

// ---------------
// String allows to handle
// const ints as strings
//  ---------------
func (s State) String() string {
	switch s {
	case OffHook:
		return "OffHook"
	case Connecting:
		return "Connecting"
	case Connected:
		return "Connected"
	case OnHold:
		return "OnHold"
	case OnHook:
		return "OnHook"
	}
	return "Unknown"
}

func main() {
  var state State //do something with it
  ...
  fmt.Println("The state is currently", state)
}
```
## FIFO Queue: Push / Pop

```go
var queue = []int
//push
queue = append(queue, 1)
queue = append(queue, 9)
queue = append(queue, 19)
//pop
var first int
first, queue = queue[0], queue[1:]
```

## LIFO Stack: Push / Pop

```go
var stack = []int
//push
stack = append(stack, 1)
stack = append(stack, 9)
stack = append(stack, 19)
//pop
var last int
last, stack = stack[len(stack) - 1], stack[:len(stack) - 1]
```

## Convert string 

- to a slice of bytes:

```go
// convert from a string to a slice of bytes 
// because of UTF-8 encoding, some characters in the string may 
// end up in 1, 2, 3 or 4 bytes
mystring := "hello this is string"
byteslice := []byte(mystring)

- to a slice of runes:

// convert from a string to a slice of runes
runeslice  := []rune(mystring)

// from a slice of bytes or runes to a string.
b := []byte{'h','e','l','l','o'} // Composite literal.
s := string(b)
i := []rune{257,1024,65}
r := string(i)
```

## When to use string vs []bytes

- se `string` by default when you're working with text. 
- use `[]byte` instead if one of the following conditions applies:
  - the mutability of a []byte will significantly reduce the number of allocations needed
  - you are dealing with an API that uses []byte, and avoiding a conversion to string will simplify your code.

## DB Querying 

```go
func doQuery() {
  dbTables := []string{"job_m", "job_s", "job_i"}

  for _, table := range dbTables {
    err := db.QueryRow("SELECT id FROM ? WHERE target = ?", table, target).Scan(&jobID)
    if errors.Is(err, sql.ErrNoRows) {
      continue
    }

    if err != nil {
       return fmt.Errorf("whatever")
    }
    // if found, do whatever
  }
}
```