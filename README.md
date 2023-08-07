# GO

- [File Organisation](#file-organisation)
- [Variables, Primitive Types and Keywords](#variables-primitive-types-keywords)
- [Control Structures](#control-structures)
- [Advanced Data Types](#advanced-data-types)
- [Functional Programming](#functional-programming)


## File Organisation

- programs are organised into packages
- unique project/package name
- packages must always be in their own directory (cmd / web / main)
- there can be only one main function in a package

> A `package` is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.

- a repository contains one or more modules

> A `module` is a collection of related Go packages that are released together. A Go repository typically contains only one module, located at the root of the repository. A module can be defined locally without belonging to a repository. However, it's a good habit to organise your code as if you will publish it someday.

- an import path is a string used to import a package
- the first statement in a go source file must be `package packagename`

> A package's import path is its module path joined with its subdirectory within the module. For example, the module github.com/google/go-cmp contains a package in the directory cmp/. That package's import path is github.com/google/go-cmp/cmp. Packages in the standard library do not have a module path prefix.

- root dir `$HOME/go`
- `src` and `bin` inside the project
- executable binary are installed into `$HOME/go/bin/`

> The install directory is controlled by the GOPATH and GOBIN environment variables. If GOBIN is set, binaries are installed to that directory. If GOPATH is set, binaries are installed to the `bin` subdirectory of the first directory in the GOPATH list. Otherwise, binaries are installed to the bin subdirectory of the default GOPATH (`$HOME/go` or `%USERPROFILE%\go`).

> To set the go env to custom dir, use `$ go env -w GOBIN=/somewhere/else/bin`
> To unset a variable previously set by go env -w, use go env -u: `$ go env -u GOBIN`

### Documentation

To install documentation, follow the steps:

1. Add path variables to your ~/.bashrc or ~/.zshrc:

```bash
export GOPATH=$HOME/go # or somewhere else
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```

2. Source your ~/.bashrc or ~/.zshrc:

```bash
source ~/.zshrc # or source ~/.bash_profile
```

3. Run

```bash
mkdir -p $GOPATH $GOPATH/src $GOPATH/pkg $GOPATH/bin # check beforehand if some of these folders are already installed
go install golang.org/x/tools/cmd/godoc@latest
godoc -http=localhost:6060
```

4. Open your browser at `localhost:6060`

### Compile and Run Code

Example: `helloworld.go`

- compile: `go build helloworld.go`
- run `./helloworld`

or use a shortcut: `go run helloworld.go`

> When opening a directory in VSCode that consists of multiple Go projects the following error appears: `gopls requires a module at the root of your workspace...` 

From Go 1.18 onwards there is native support for multi-module workspaces. This is done by having a `go.work` file present in your parent directory.

For a directory structure such as:

```
$ tree /my/parent/dir
/my/parent/dir
├── project-one
│   ├── go.mod
│   ├── project-one.go
│   └── project-one_test.go
└── project-two
    ├── go.mod
    ├── project-two.go
    └── project-two_test.go
```

Create and populate the file by executing `go work`:

```bash
cd /my/parent/dir
go work init
go work use project-one
go work use project-two
```

This will add a go.work file in your parent directory that contains a list of directories you marked for usage:

```
go 1.18

use (
    ./project-one
    ./project-two
)
```

## Variables, Primitive Types and Keywords

- uninitialised values are implicitly initialised with `0`, `nil`, or `''` depending on data type.
- `Go is pass by value`:
    - similarly to Javascript, primitive datatypes are copied
    - for complex datatypes you pass a pointer to avoid memory oveload.

- `postfix types`: types are given `after the variable name`

```go
var a int
// instead of 'int a'
```
- can do `type inference` if a compact syntax is used (`only possible inside of a function`)

```go
var a int                           
var b bool                          
a = 15
b = false

func funcScoped() {
  a := 15      // type is deduced
  b := false  // type is deduced
}
```

- multiple variables of the same type can also be declared on a single line
- parallel assignment is possible with compact syntax

```go
// multiple variables of the same type can also be declared on a single line
var x, y int  

func funcScoped() {
  a, b := 20, 16 // a and b both integer variables and assigns 20 to a and 16 to b
}
```

- strictly typed language with no implicit casting (unlike C++). Casting is always explicit --> assignment between items of different type `requires an explicit conversion`!!

```go
i := 36
j := 36.5
sum := i + int(j) //explicitely converted to int
```

- `constants` can be untyped and typed
- `created at compile time`, and can only be numbers, strings, or booleans

> Any constant in golang, named or unnamed, is untyped unless given a type explicitly. For example an untyped floating-point constant like 4.5 can be used anywhere a floating-point value is allowed. Use untyped constants to temporarily escape from Go’s strong type system until their evaluation in a type-demanding expression.

```go
const untypedInt = 1
const typedInt int = 1
```

### Primitive Data Types

- `bool`
- `int`:
  int (will be based on the length of your machine: 32 or 64 bits), `int8`, `int16`, `int32`, `int64`, `uint8`, `uint16`, `uint32`, `uint64`, `byte` (alias for uint8)
- float:
  note: no float type, `float32` and `float64`
- `string` (IMMUTABLE!!)

> `Go strings are immutable and behave like read-only byte slices` with a few extra properties. 

```go
// Wrong
var s string = "hello"
s[0] = 'c' // compiler error

// Right
s := "hello"
c := []rune(s)  // convert s to an array of runes
c[0] = 'c'      // change the first element of this array
s2 := string(c) // create a new string s2 with the alteration
fmt.Printf("%s\n", s2) 
```

- `rune` (alias for int32) is an UTF-8 encoded code point

> Useful example: iterating over characters in a string. You could loop over each byte (which is only equivalent to a character when strings are encoded in 8-bit ASCII, which they are not in Go!). But to get the actual characters you should use the rune type.

- `complex numbers`:
  - native support
  - `complex128` (64 bit real and imaginary parts) or `complex64` (32 bit real and imaginary parts)

### Go Keywords

- import
- package
- var
- type
- struct
- interface	
- const	
- func 
- select
- defer
- chan
- go
- map	
- case			
- goto	
- switch	
- break		
- continue
- fallthrough
- default	
- return			
- for
- if	
- else	

## Control Structures

- classic, some difference in the variations of `for` when used with `range`
- operators can be omitted with `_`
- `for range`: a mapper function for various data structures:
    - two values are returned for each iteration: 
    - the first is the `index`, and the second is a `copy of the element at that index`

```go
for _, item := range lst {
  // item is the copy of the element
  // in the list
	fmt.Printf("%#v\n", item)
}

// can also use range on strings directly: 
// --> it will break out the individual Unicode characters
// into runes (UTF-8 characters that may be up to 32 bits)
for pos, char := range "Go!" {
  fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
}

// working with list of strings
list := []string{"a", "b", "c", "d", "e", "f"}
for k, v := range list {
  // do something with k and v
}
```
- in `switch`, the break is done implicitely by a compiler
- `fallthrough` is possible

```go
switch i {
  case 0:  fallthrough
  case 1: 1
    f()
  default:
     g() 2
}

// or another mode of the same statement:
switch i {
  // instead of fallthrough, 2 cases:
  case 0, 1: 1
    f()
  default:
    g()
}
```
- since if and switch accept an initialization statement, it’s common to see one used to set up a (local) variable:

```go
if err := SomeFunction(); err == nil {
  // do something if no error occurred
} else {
	return err
}
```
- idiomatic to return err immediately if it's occurred:

```go
if err := SomeFunction(); err != nil {
  return err
}
// do something if no error occurred
```

- `for` loop:
  - `for init; condition; post { }` - traditional loop;
  - `for condition { }` - a while loop;
  - `for { }` - an endless loop

## Advanced Data Types

### Arrays

- same as in C++: fixed length, must be known on compilation time
- but uninit values are by default initialised to `0`, `nil`, or `''`
- `assigning one array to another copies all the elements`.
- if you `pass an array to a function` it will receive `a copy of the array`, not a pointer to it.

```go
// initialisation
var theArray [3]string
theArray[0] = "India"  // Assign a value to the first element
theArray[1] = "Canada" // Assign a value to the second element
theArray[2] = "Japan"  // Assign a value to the third element

// using array literal syntax
x := [5]int{10, 20, 30, 40, 50}   // Intialised with values
var y [5]int = [5]int{10, 20, 30} // Partial assignment

z := [...]int{100, 200, 300}   // Intialised with values, size is determined
// on compilation by counting the elements

// technically, two distinct array types, despite being arrays of ints
// because they hold different size
arr1 int[2]
arr2 int[3]
```
- no negative indexing (unlike Python or JS):

```go
arr int[2]
fmt.Println(arr[-1]) //invalid array index -1 (index must be non-negative)
```
> When declaring arrays you always have to type something in between the square brackets, either a number or three dots (...), when using a composite literal. 

> When using multidimensional arrays, you can use the following syntax: `a := [2][2]int{ {1,2}, {3,4} }`.

- arrays are used on a few occasions (they cannot be resized, not dynamic, their size must be known on compilation time) 
- mostly used as `slices` (Go's own data type) that are not contrained by length (think vector in C++)

### Slices

- like `dynamically-sized arrays` (think Vector in C++)
- a slice is a `pointer to an array`
- slices are `reference types` -->
    - changing the elements of a slice modifies the corresponding elements of its underlying array
    - assign one slice to another, both refer to the same underlying array
    - `slicing does not copy the slice’s data`. It creates a new slice value that points to the original array. This makes slice operations as efficient as manipulating array indices
    - a slice cannot be grown beyond its capacity. Attempting to do so will cause a runtime panic, just as when indexing outside the bounds of a slice or array
    - slices cannot be re-sliced below zero to access earlier elements in the array.

#### Member Functions

every slice has 3 properties:

- `pointer`: start of the slice
- `length`: the number of elements in the slice: len()
- `capacity`: maximum number of elements: cap()

[visual explanation](https://go.dev/blog/slices-intro)

#### Creating Slices

- slices can be created explicitly or via `make()` API:

```go
// create slice explicitly from an array:
var arr[] int = []int{1,2,3}
sl := arr[:3]
// make a slice of empty strings with length 3
s := make([]string, 3)
```

#### Growing Slices

> To increase the capacity of a slice one must create a new, larger slice and copy the contents of the original slice into it. This technique is how dynamic array implementations from other languages work behind the scenes.

- to modify the length of the slice, use `append` and `copy`:
  - `append` returns a new slice with 1 or more values; immutable, in contrast to Javascript
  - the append function appends zero or more values to a slice and returns the result: a slice with the same type as the original. If the original slice isn’t big enough to fit the added values, append will allocate a new slice that is big enough.
  - `copy` copies slice elements from a source to a destination, and returns the number of elements it copied. This number is the minimum of the length of the source and the length of the destination.

```go
// append
s0 := []int{0, 0}
s1 := append(s0, 2) // s1 equal to []int{0, 0, 2}
s2 := append(s1, 3, 5, 7) // s2 equal to []int{0, 0, 2, 3, 5, 7}
// make it clear explicit that you’re appending another slice, instead of a single value
s3 := append(s2, s0...) // s3 equal to []int{0, 0, 2, 3, 5, 7, 0, 0}.

// copy
var a = []int{0, 1, 2, 3, 4, 5, 6, 7}
var s = make([]int, 6)
n1 := copy(s, a[0:]) // []int{0, 1, 2, 3, 4, 5}
n2 := copy(s, s[2:]) // []int{2, 3, 4, 5, 4, 5}
```

- slice `[start:end]` operator similar to Python

```go
s[1:3]
```

- initialise values for a slice

```go
t = []string{"go", "is", "cool"}
```

- convert array to a slice

```go
arr := [3]string{"go", "is", "cool"}
sl := arr[:]
```

- combine two slices together

```go
sl1 := []string{"c", "is", "cool"}
sl2 := []string{"go", "is", "cooler"}
sl1 = append(sl1, sl2...) //note the spread operator
```

- removing element(s) from a slice

`slice = append(slice[:i], slice[i+1:]...)`

```go
sl := []string{"keep", "keep", "remove", "keep"}
sl = append(sl[:2], sl[3:]...) //note the spread operator
//to remove the range, use the same approach
```

### Maps

hash table

keys are unsorted but we can print map using fmt package in key-sorted order to ease testing

keys can be anything other than string

you can safely pass a map around by value, the underlying map will be updated, not a copy of the map.
pass-by-value in the case of a map means passing the address of the map, not the contents of the map

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

to iterate over map keys in a certain order, first sort the keys as a slice and order them in the way you want:

```go
langs: = map[string]int {
  "go": 10,
  "c++": 9,
  "python":8
}

var k []string
for l := range langs {
  k = append(k, l)
}
sort.Strings(k) //will be in alphab order

for _, l := range langs {
  fmt.Println(l, langs[k])
}

```

to create a `Set`, use `map[string]bool` with true for vals in the map.

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

[when to use a value receiver or a pointer receiver](https://github.com/golang/go/wiki/CodeReviewComments#receiver-type)

## Functional Programming

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

`named return values and naked return`

return values can optionally be named. If named return values are used, a return statement without arguments will return those values. This is known as a 'naked' return.

```go
func SumAndMultiplyThenMinus(a, b, c int) (sum, mult int) {
  sum, mult = a+b, a*b
  sum -= c
  mult -= c
  return //named values are implicitly returned
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

- `no function overloading` (like C, unlike C++)

## OOP

- polymorphism is achieved with `structs` that can have member functions (sort of). It's done externally, in the code (see examples).

Go is not a classic OOP language. There is no type hierarchy.
Go OOP Features:
1. Structs: serve similar purpose to classes
2. Methods: can operate on particular type and mock a member function
3. Embedding: we can embed anonymous types inside each other.
4. Interfaces: have no implementation. Objects that implement all the interface methods automatically implement the interface. There is no inheritance or subclassing or "implements" keyword.

`Encapsulation`:
Go encapsulates things at the package level. Names that start with a lowercase letter are only visible within that package. You can hide anything in a private package and just expose specific types, interfaces, and factory functions.

`Inheritance`:
composition by embedding an anonymous type is equivalent to implementation inheritance.

`Polymorphism`:
a variable of type interface can hold any value which implements the interface. This property of interfaces is used to achieve polymorphism in Go.

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

## Interfaces

- uses `interfaces` (method signatures), same as Typescript. 

> Unlike Typescript, in Go, a `type` implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword. Under the hood, interface values can be thought of as a tuple of a value and a concrete type: `(value, type)`.

An interface and a type are `structurally equivalent` if they both define a set of methods of the same name, and where methods from each share the same number of parameters and return values, of the same data type.

In Go, we can define as many little interfaces as we want.
> `Duck typing` (structural typing): 'If it walks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck.'

No need to define `implements` means that the interfaces defined in third-party packages can still be included in our own code base (they only need to match the methods).

```go
type NewUser struct {
  Email    string
  Password string
}

type DBUser struct {
  ID           int
  Email        string
  PasswordHash string
}

//declare the interface inline
func CreateUser(user *User, db interface {
  SaveUser(*DBUser) error
}) error {
  var dbUser DBUser
  // Validate the user first...
  if user.Email == "" {
    return fmt.Errorf("email is required")
  }
  dbUser.Email = strings.ToLower(user.Email)

  if user.Passwword == "" {
    return fmt.Errorf("password is required")
  }
  dbUser.Password = pretendBCrypt(user.Password)

  // Then save the user with the db interface
  return db.SaveUser(&dbUser)
}

func pretendBcrypt(pw string) string { return pw }

type DataStore struct {}

//use the SaveUser declared ob inline interface
func (ds *DataStore) SaveUser(u *DBUser) error {
  return nil
}

func (ds *DataStore) DeleteUser(id int) error {
  return nil
}

func (ds *DataStore) CreateWidget(w *DBWidget) error {
  return nil
}

// ... + more methods
```

## Concurrency

- `concurrency` is build in
- `goroutine` is "a lightweight thread of execution"

- goroutine has a simple model: 
it is a function executing in parallel with other goroutines in the same address space. It is lightweight, costing little more than the allocation of stack space. And the stacks start small, so they are cheap, and grow by allocating (and freeing) heap storage as required.

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

`channels`: a way of communicating between coroutines;
a typed conduit through which you can send and receive values with the channel operator, `<-`

```go
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```

created with `make()`

```go
intChannel := make(chan int)
```

`mutex`: mutual exclusion; a way to prevent interleavings and ensure that different coroutines won't write into the same variable at the same time; uses a `binary semaphore` (flags up and down)

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

`Unbuffered channel`: the sender will block until the receiver receives the data from the channel; the receiver will also block until the sender sends the data into the channel

`Buffered channel`: the sender will block when there is not empty slot of the channel; the receiver will block the channel when it is empty.

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

## Print formatting

bool:                    `%t`
int, int8 etc.:          `%d`
uint, uint8 etc.:        `%d`, %#x if printed with %#v
float32, complex64, etc: `%g`
string:                  `%s`
chan:                    `%p`
pointer:                 `%p`

General:
- `%v`	the value in a default format
- `%+v` when printing structs, the plus flag adds field names
- `%#v` a Go-syntax representation of the value
- `%T`	a Go-syntax representation of the type of the value
- `%%`	a literal percent sign; consumes no value

Boolean
- `%t`	the word true or false

Integer
- `%b`	base 2
- `%c`	the character represented by the corresponding Unicode code point
- `%d`	base 10
- `%o`	base 8
- `%O`	base 8 with 0o prefix
- `%q`	a single-quoted character literal safely escaped with Go syntax.
- `%x`	base 16, with lower-case letters for a-f
- `%X`	base 16, with upper-case letters for A-F
- `%U`	Unicode format: U+1234; same as "U+%04X"

Floating-point and complex constituents:
- `%b`	decimalless scientific notation with exponent a power of two
- `%e`	scientific notation, e.g. -1.234456e+78
- `%E`	scientific notation, e.g. -1.234456E+78
- `%f`	decimal point but no exponent, e.g. 123.456
- `%F`	synonym for `%f`
- `%g`	%e for large exponents, %f otherwise. Precision is discussed below.
- `%G`	%E for large exponents, %F otherwise
- `%x`	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
- `%X`	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20

String and slice of bytes:
- `%s`	the uninterpreted bytes of the string or slice
- `%q`	a double-quoted string safely escaped with Go syntax
- `%x`	base 16, lower-case, two characters per byte
- `%X`	base 16, upper-case, two characters per byte

Slice:
- `%p`  base 16 notation, with leading 0x

Pointer:
- `%p`	base 16 notation, with leading 0x
- `%b, %d, %o, %x` and `%X` verbs also work with pointers, formatting the value exactly as if it were an integer.

```go
//simple default formatting
//does not support formatting directives
i := 5
fmt.Println("Hello, playground", i)  // Hello, playground 5

//using formatting directives
fmt.Printf("Hello, playground %d\n", i) // Hello, playground 5

//%v can figure out the type of the var
fmt.Printf("Hello, playground %v\n", i) // Hello, playground 5

coral := [4]string{"blue coral", "staghorn coral", "pillar coral", "elkhorn coral"}
//"%q\n" is the formatting directive, sets ""
fmt.Printf("%q\n", coral) // ["blue coral" "staghorn coral" "pillar coral" "elkhorn coral"]
```

## Tips. Tricks, Interview Questions

1. How to swap values:

simple: a,b, = b,a

```go
func swap(a, b string) (string, string) {
  return b, a
}
```

2. Mind the difference:
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
3. Concat strings:
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

4. Check if map contains a key:
```go
if val, ok := dict["someval"]; ok {
  //do something
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

6. Reverse a slice of ints:
```go
func reverse(s []int) {
  for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
    s[i], s[j] = s[j], s[i]
  }
}
```

7. Print const (iotas):

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
8. FIFO Queue: Push / Pop

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

9. LIFO Stack: Push / Pop

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

## Resources

[interfaces](https://fast4ward.online/posts/a-guide-to-interfaces-in-go/)

[The language tour](https://tour.golang.org/)
[learn how to organize your Go workspace](https://golang.org/doc/code.html)
[be more effective at writing Go](https://golang.org/doc/effective_go.html)
[Go spec](https://golang.org/ref/spec)
[Go articles](https://golang.org/doc/#articles)

Websites:
[great resources for Gophers in general](https://blog.gopheracademy.com)
[weekly podcast of Go awesomeness](http://gotime.fm)
[examples of how to do things in Go](https://gobyexample.com)
[how to use SQL databases in Go](http://go-database-sql.org) 
[tips on how to write more idiomatic Go code](https://dmitri.shuralyov.com/idiomatic-go) 
[will help you avoid gotchas in Go](https://divan.github.io/posts/avoid_gotchas) 
[tutorials to help you get started in Go](https://golangbot.com)
[a collection of articles around various aspects of Go](https://tutorialedge.net) 

Videos:
[related to Go from various authors](http://gophervids.appspot.com)

Books:
[The Go Programming Language](http://gopl.io/)
[Go in action](https://www.manning.com/books/go-in-action)

Other:
[how to organize your Go project](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1#.ds38va3pp)
[walkthrough the various standard library packages](https://medium.com/go-walkthrough)
[Go resources](https://github.com/golang/go/wiki#learning-more-about-go)

https://github.com/dariubs/GoBooks