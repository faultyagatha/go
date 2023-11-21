# GO

- [File Organisation](#file-organisation)
- [Variables, Primitive Types and Keywords](#variables-primitive-types-keywords)
- [Control Structures](#control-structures)
- [Advanced Data Types](#advanced-data-types)
- [Functions](#functions)
- [Packages](#packages)
- [Interfaces](#interfaces)
- [Concurrency](#concurrency)
- [Print Formatting](#print-formatting)
- [Debugging](#debugging)
- [Compilation](#Compilation)
- [Tools](#tools)


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

[how to write extended documentation](https://go.dev/src/encoding/gob/doc.go)

### Compile and Run Code

Example: `helloworld.go`

- compile: `go build helloworld.go`
- run `./helloworld`

or use a shortcut: `go run helloworld.go`

> When opening a directory in VSCode that consists of multiple Go projects the following error appears: `gopls requires a module at the root of your workspace...` 

From Go 1.18 onwards there is native support for multi-module workspaces. This is done by having a `go.work` file present in your parent directory.

> In workspace mode, the go. work file will be used to determine the set of main modules used as the roots for module resolution, instead of using the normally-found go.mod file to specify the single main module. 

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

### Change a version of go.mod

`go mod edit -go=1.18`

## Variables, Primitive Types, Keywords

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
  int (will be based on the length of your machine: 32 or 64 bits), `int8`, `int16`, `int32`, `int64`, `uint8` (same as byte), `uint16`, `uint32`, `uint64`
- `byte` (alias for uint8): 
  - the byte type is only used to semantically distinguish between an unsigned integer 8 and a byte. 
  - the range of a byte is 0 to 255 (same as uint8)
  - a []byte can hold non-ASCII characters if they are encoded  as bytes (for example. in UTF-8) but in this case we may not have 1-1 char-byte mapping (UTF-8 codepoints may be represented from 1 to 4 bytes).
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

> NOTE: `A string is a sequence of bytes and not of a Rune`. A string may contain Unicode text encoded in UTF-8. But, the Go source code encodes as UTF-8, therefore, no need to encode the string in UTF-8.

> `UTF-8 encodes all the Unicode in the range of 1 to 4 bytes`, where 1 byte is used for ASCII and the rest for the Rune.

> ASCII contains a total of 256 elements and out of which, 128 are characters and 0-127 are identified as code points. Here, code point refers to the element which represents a single value.

- `complex numbers`:
  - native support
  - `complex128` (64 bit real and imaginary parts) or `complex64` (32 bit real and imaginary parts)

```go
str := "百度一下, 你就知道" 
// an array of the Unicode code points (the number of characters)
fmt.Println("String length", len([]rune(str)))
// the length of the underlying byte array
fmt.Println("Byte length", len(str))
```

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

- removing element(s) from a slice:

1. Non-idiomatic to Go (but the 'right' from cs point of view):
  - create a copy of the slice to avoid modifying an original slice

```go
func OriginalRemoveIndex(arr []int, pos int) []int {
  // Make a copy of the original slice
  ret := make([]int, 0)
  // Append everything before the index
  ret = append(ret, s[:index]...)
  // Append everything after the index
  return append(ret, s[index+1:]...)
}

func main() {
  original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  fmt.Println("original: ", original) //[0 1 2 3 4 5 6 7 8 9]
  result := result(original, 5)

  fmt.Println("original: ", original) //[0 1 2 3 4 5 6 7 8 9]
  fmt.Println("result: ", result) //[0 1 2 3 4 6 7 8 9]

  result[0] = 999
  fmt.Println("original: ", original) //[0 1 2 3 4 5 6 7 9 9]
  fmt.Println("result: ", result) //[999 1 2 3 4 6 7 8 9]
}
```

2. Idiomatic go Go ('re-slicing')
`slice = append(slice[:i], slice[i+1:]...)`
- note, this is expensive (as normally, with any array), so use sparringly

```go
func RemoveIndex(s []int, index int) []int {
    return append(s[:index], s[index+1:]...)
}

func main() {
  original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  fmt.Println(original) //[0 1 2 3 4 5 6 7 8 9]
  // IMPORTANT: the original slice must be reassigned
  // Otherwise --> see above (1.)
  original = RemoveIndex(original, 5)
  fmt.Println(original) //[0 1 2 3 4 6 7 8 9]
}
```

```go
sl := []string{"keep", "keep", "remove", "keep"}
sl = append(sl[:2], sl[3:]...) //note the spread operator
//to remove the range, use the same approach
```

3. Use `Delete()` member function on slices (since Go 1.21)
- this is just a wrapper function to `append(s[:index], s[index+1:]...)`

```go
slice := []int{1, 2, 3, 4}
// IMPORTANT: the original slice must be reassigned
// Otherwise --> see above (1.)
slice = slices.Delete(slice, 1, 2)
fmt.Println(slice) // [1 3 4]
```

### Maps

- hash table or in its simplest form, an array indexed by strings
- keys are unsorted but we can print map using `fmt package` in key-sorted order to ease testing
- keys can be anything other than string
- map is a reference typ
- you can safely pass a map around by value, the underlying map will be updated, not a copy of the map
- `pass-by-value` for a map is a special case: means passing the address of the map, not the contents of the map

#### Create a Map

- define a map: `map[<from type>]<to type>`
- use `make()` when only declaring a map

```go
monthdays := map[string]int {
  "Jan": 31, "Feb": 28, "Mar": 31,
  "Apr": 30, "May": 31, "Jun": 30,
  "Jul": 31, "Aug": 31, "Sep": 30,
  "Oct": 31, "Nov": 30, "Dec": 31, 1
}

// or using make (in this case, the map is empty)
monthdays := make(map[string]int)
```

#### Changing Map

```go
// add a new element to the map monthdays
monthdays["Undecim"] = 30

//  if you use a key that already exists, the value will be silently overwritten
monthdays["Feb"] = 29

// check if the value is present
value, ok := monthdays["Jan"] 
```
- optional second return value indicates whether the key was present in the object
- idiomatic for Go: `“comma ok” form.` when checking values in the maps
- delete values with `delete`

```go
delete(monthdays, "Undecim")
```

- loop through with `for range` (think of `for in` in Javascript)

```go
fruit := map[string]string{"a": "apple", "b": "banana"}
for key, val := range fruit {
  fmt.Printf("%s -> %s\n", key, val)
}
```

- to iterate over map keys in a certain order, first sort the keys as a slice and order them in the way you want:

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

- to create a `Set`, use `map[string]bool` with true for vals in the map.

### Strings

```go
// byte index, char (rune)
for i, c := range "go" {
  fmt.Println(i, c)
}
```

### Pointers

- like in C++ but no pointer arithmetic --> act more like references:
  - `*p++` will be interpreted as `(*p)++`: dereference and then increment the value
- there is `new` but no `delete`
- newly declared pointer points to nothing, has a `nil`-value (nullptr)
- dereferenced with `*`

```go
// value is passed as a copy of the underlying value 
// and here we operate on that copy only
func passedByValue(value int) {
  value = 0
}

// mutate the underlying value by assigning 
// a new int at the referenced address
func passedByReference(ref *int) {
  // dereference and assign 0 to the value
  *ref = 0
}

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

#### Memory Allocation

- the compiler decides where to allocate based on `escape analysis`
- using `new` doesn't imply using the heap (!!)
- can be done with `new` and `make`
- `new(T)` returns *T pointing to a zeroed T
- `make(T)` returns an initialized T

> The built-in function `new` is essentially the same as in other languages: `new(T)` allocates zeroed storage for a new item of type `T` and returns its address, a value of type `*T`. Or in other words, it `returns a pointer to a newly allocated zero value of type T`.

```c++
int *pNum = new int;  // allocated memory but memory location contains junk value
int *pNum1{ new int}; // allocated memory but memory location contains junk value
```

```go
new(Point)    // allocated enough memory to store Point but memory location contains junk value
&Point{}      // the same as above but using & 
&Point{2, 3}  // combines allocation and initialization

new(int)      // allocated enough memory to store int but memory location contains junk value
&int          // illegal (!!) for simple data types

// Works, but it is less convenient to write than new(int)
var i int
&i

p := new(chan int)   // p has type: *chan int
c := make(chan int)  // c has type: chan int

// if Go didn't have new and make but had the built-in function NEW, the code would have looked like this:
p := NEW(*chan int)  // * is mandatory
c := NEW(chan int)
```

> The built-in function `make(T, args)` serves a purpose different from new(T). It `creates slices, maps, and channels only`, and it `returns an initialized` (not zero!) `value of type T`, and `not a pointer: *T`. The reason for the distinction is that these three types are, under the covers, references to data structures that must be initialized before use. A slice, for example, is a three-item descriptor containing a pointer to the data (inside an array), the length, and the capacity; until those items are initialized, the slice is nil. For slices, maps, and channels, make initializes the internal data structure and prepares the value for use.

For example, `make([]int, 10, 100)` allocates an array of 100 ints and then creates a slice structure with length 10 and a capacity of 100 pointing at the first 10 elements of the array. In contrast, `new([]int)` returns a pointer to a newly allocated, zeroed slice structure, that is, a pointer to a nil slice value.

#### Composite Literals
- construct new composite values each time they are evaluated
- can be used with objects (structs), arrays, slices, maps
- consist of the type of the literal followed by a brace-bound list of elements

```go
buffer := [10]string{}             // len(buffer) == 10
intSet := [6]int{1, 2, 3, 5}       // len(intSet) == 6
// The notation ... specifies an array length equal to the maximum element index plus one.
days := [...]string{"Sat", "Sun"}  // len(days) == 2

[...]Point{{1.5, -3.5}, {0, 0}}     // same as [...]Point{Point{1.5, -3.5}, Point{0, 0}}
[][]int{{1, 2, 3}, {4, 5}}          // same as [][]int{[]int{1, 2, 3}, []int{4, 5}}
[][]Point{{{0, 1}, {1, 2}}}         // same as [][]Point{[]Point{Point{0, 1}, Point{1, 2}}}
map[string]Point{"orig": {0, 0}}    // same as map[string]Point{"orig": Point{0, 0}}
map[Point]string{{0, 0}: "orig"}    // same as map[Point]string{Point{0, 0}: "orig"}

type PPoint *Point
[2]*Point{{1.5, -3.5}, {}}          // same as [2]*Point{&Point{1.5, -3.5}, &Point{}}
[2]PPoint{{1.5, -3.5}, {}}          // same as [2]PPoint{PPoint(&Point{1.5, -3.5}), PPoint(&Point{})}

noteFrequency := map[string]float32{
	"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
	"G0": 24.50, "A0": 27.50, "B0": 30.87,
}
```

- pointers and composite literals:

```go
// Option 1:
func NewFile(fd int, name string) *File {
  if fd < 0 {
      return nil
  }
  f := new(File)
  f.fd = fd
  f.name = name
  f.dirinfo = nil
  f.nepipe = 0
  return f
}

// Option 2: using composite literals
func NewFile(fd int, name string) *File {
  if fd < 0 {
      return nil
  }
  f := File{fd, name, nil, 0}
  // It is OK to return the address of a local variable,
  // the storage associated with the variable survives after the function returns.
  return &f	
}
```
- in the example above, the expressions `new(File)` and `&File{}` are equivalent. Note: use of new is discouraged.

### Structs

- like C structs
- methods can be associated with data assigned explicitely:
  - a hidden object (struct instance, for example) will be passed by value implicitely -->
  - use pointer to avoid value copying

#### Calling Methods on Structs

1. Create a function that takes the type as an argument

2. Create a function that works on the type (idiomatic to Go)

```go
type person {
  name string
  age int
}

// using 1-st option
func initPerson1(p *person, n string, a int) { /* */ }

// using 2-d option
func (p *person) initPerson2(n string, a int) { /* */ }

var p *person
// using 1-st option
initPerson1(p, "Stan", 39)
// using 2-d option
p.initPerson2("Stan", 39)
```

> IMPORTANT: `If x is addressable and &x’s method set contains m, x.m() is shorthand for (&x).m()` 

```go
var p person	              // Not a pointer but initPerson2() should be called on a pointer to person -->
p.initPerson2("Stan", 39)   // Go will translate this call to (&p).initPerson2("Stan", 39)
```

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

> Go knows that the method is on the struct since the struct is a parameter. The method is named in the function call defined at the top i.e. `area()`.

> Go also magically converts between values and pointers for method calls. So no need for dereference. You can control the behavior by specifying a pointer receiver type to avoid copying the struct on method calls or to allow the method to mutate the underlying values.

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

- it's idiomatic to initiate a new struct with a factory function.

[when to use a value receiver or a pointer receiver](https://github.com/golang/go/wiki/CodeReviewComments#receiver-type)

> Important: The `receiver type` must be of the form `T` or `*T` where T is a type name. T is called the receiver base type or just base type. `The base type must not be a pointer or interface type` and must be declared in the same package as the method.

### Type Conversion

- only explicit
- not all conversions are allowed

```go
// convert from a string to a slice of bytes 
// because of UTF-8 encoding, some characters in the string may 
// end up in 1, 2, 3 or 4 bytes
mystring := "hello this is string"
byteslice := []byte(mystring)

// convert from a string to a slice of runes
runeslice  := []rune(mystring)

// from a slice of bytes or runes to a string.
b := []byte{'h','e','l','l','o'} // Composite literal.
s := string(b)
i := []rune{257,1024,65}
r := string(i)
```

## Functions

- `pass-by-value`
- func ([receiver]) [name] ([params]) ([return values])
- `multiple return values` (think Typescript or Python tuples):
  - useful to return a value and error
  - removes the need for in-band error returns (such as -1 for EOF) and modifying an argument
  - if you want the return parameters not to be named you only give the types: (int, int)
  - `named return values and naked return`

>  The `return` or result parameters of a Go function can be given names and used as regular variables, just like the incoming parameters. When named, they are initialised to the zero values for their types when the function begins. If the function executes a return statement with no arguments, the current values of the result parameters are returned. Using these features enables you to do more with less code.

```go
// named return
func SumAndMultiplyThenMinus(a, b, c int) (sum, mult int) {
  sum, mult = a+b, a*b
  sum -= c
  mult -= c
  return //named values are implicitly returned
}

// naked return
func doubleReturn(x int) (int, int) {
	return x, x + 1
}
```
- return values can be omitted with `_`:

```go
var sum, _ = doubleReturn(5)
```

- functions can be declared in any order: the compiler scans the entire file before execution, so no need for function declaration
- Go `does not allow nested functions`, but you can work around this with `anonymous functions`.

- `first-class`; mostly like in Javascript (treated like other types) -->
- can be assigned to variables:

```go
// a stores a pointer to the anonymous function
a := func() { 1
	fmt.Println("Hello")
}
```
- can be used as `callbacks`:

```go
func callback(y int, f func(int)) {
  f(y)
}
```

- can be used as `closures`, like in Javascript:

```go
func factory() func() int {
  i := 0

  return func() int {
    i++
    return i
  }
}
```
- variadic functions (`nums ...int`): think `...rest` operator in Javascript
  - under the hood this is a slice of ints
  - we can do slicing on slices

```go
// under the hood this is a slice of ints
func spread(nums ...int) int {
  total := 0
  for _, num := range nums {
    total += num
  }
  return total
}

// 1. call the function with variadic param
spread((1, 3, 6, 4)) // 14

// 2. call the variadic function by slicing 
mySlice := []int{1, 3, 6, 4}
spread(mySlice[:2]) // 4 (1+3)
```

- `no function overloading` (like C, unlike C++)

[more on closures](https://eli.thegreenplace.net/2019/go-internals-capturing-loop-variables-in-closures/)

### Deferring Functions

- functions can be `deferred` with `defer` key-word:
  - we can defer multiple functions
  - deferred functions are executed in LIFO order
  - deferred function can be closures

```go
func ReadWrite() bool {
  file.Open("filename")
  // this will be executed right before
  // the function exits
  defer file.Close() 1
  // Do your thing
  if failureX {
    return false
  }
  if failureY {
    return false
  }
  return true
}

for i := 0; i < 5; i++ {
  defer fmt.Printf("%d ", i) // 4 3 2 1 0.
}

func f() (ret int) {
  // closure that is immediately invoked
  defer func() { 
    ret++
  }()
  return 0 // the function will return 1!
}
```

### Panic and Recover

> `Panic` is a built-in function that stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. To the caller, F then behaves like a call to panic. The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes. Panics can be initiated by invoking panic directly. They can also be caused by runtime errors, such as out-of-bounds array accesses.

- avoid use panic

> `Recover` is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.

```go
func Panic(f func()) (b bool) { 
  defer func() { 
    if x := recover(); x != nil {
      b = true
    }
  }()
  // if this function causes panic,
  // the defer function will be called -->
  // recover will be initiated
  f() 
  return // true 
}

func wrongAccess() {
  var a []int
  a[3] = 5
}
// will cause panic but will be 
// gracefully handled by recover
res := Panic(wrongAccess)
```

## Packages

- a collection of functions and data
- rules:
  - declared with `package` keyword
  - the filename does not have to match the package name
  - naming convention: lowercase, single word names 
  - may consist of multiple files
- functions and data types can be exported from packages (if named with capital letter)
- public functions can be accessed by <package>.FunctionName()
- packages can be used if installed:

```bash
% mkdir $GOPATH/src/even
% cp even.go $GOPATH/src/even
% go build
% go install
```

- in some cases, when go work is used and some other external modules should be installed (e.g., the root folder consists of many subfolders with small projects that use some other modules that must be installed), having go.mod in the root folder will not work --> 

```
├── LICENSE
├── README.md
├── go.work
├── project_1
│   ├── go.mod
│   └── go.sum
|   └── main.go
├── project_2
│   ├── go.mod
│   └── go.sum
|   └── main.go
├── project_3
│   ├── go.mod
│   └── go.sum
|   └── main.go
```

this means, create go.mod for each folder that has some external dependencies and then add them to go.work

```shell
# setup path to project_1 in the workspace
cd project_1
go mod init github/faultyagatha/myrepo/project_1
go get github/somedependency
cd ..
go work use ./project_1

# setup path to project_2 in the workspace
cd project_2
go mod init github/faultyagatha/myrepo/project_2
go get github/somedependency
cd ..
go work use ./project_2

# setup path to project_3 in the workspace ... 
```


## OOP

- polymorphism is achieved with `structs` that can have member functions (sort of). It's done externally, in the code (see examples).

- Go is not a classic OOP language
- there is no type hierarchy

- Go OOP Features:
1. Structs: serve similar purpose to classes
2. Methods: can operate on particular type and mock a member function
3. Embedding: we can embed anonymous types inside each other (see below).
4. Interfaces: have no implementation. Objects that implement all the interface methods automatically implement the interface. There is no inheritance or subclassing or "implements" keyword.

### Encapsulation

- Go encapsulates things at the package level
- names that start with a lowercase letter are only visible within that package
- one can hide anything in a private package and just expose specific types, interfaces, and factory functions.

### Inheritance via Embedding

- Go encourages composition as a way to extend the functionality of types
- composition by embedding an anonymous type is equivalent to implementation inheritance.


```go
// NewMutex is equal to Mutex, 
// but it does not have any of the methods of Mutex
type NewMutex Mutex

// inherited the method set from Mutex
type PrintableMutex struct {Mutex}.
```

- three kinds of embedding in Go:

1. Structs in structs 
2. Interfaces in interfaces
3. Interfaces in structs

[embedding in go](https://eli.thegreenplace.net/2020/embedding-in-go-part-1-structs-in-structs/)

### Polymorphism

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

- uses `interfaces` (method signatures), same as Typescript
- Go interface is similar to pure abstract base class in C++

> Unlike Typescript, in Go, a `type` implements an interface by implementing its methods. There is no explicit declaration of intent, `no implements keyword`. Under the hood, interface values can be thought of as a tuple of a value and a concrete type: `(value, type)`.

> An interface and a type are `structurally equivalent` if they both define a set of methods of the same name, and where methods from each share the same number of parameters and return values, of the same data type.

> Go can convert from one interface type to another. `The conversion is checked at run time`. If the conversion is invalid – if the type of the value stored in the existing interface value does not satisfy the interface to which it is being converted – the program will fail with a run time error.

- we can define as many little interfaces as we want

> `Duck typing` (structural typing): 'If it walks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck.'

> No need to define `implements` means that the interfaces defined in third-party packages can still be included in our own code base (they only need to match the methods).

```go
type S struct { i int }
func (p *S) Get() int  { return p.i }
func (p *S) Put(v int) { p.i = v }

// S is a valid implementation for interface I
type I interface {
  Get() int
  Put(int)
}

func f(p I) { 
  // because p implements I, it must have the Get() method
  fmt.Println(p.Get()) 
  p.Put(1) 3
}
```
- interfaces can be declared inline

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

- `type switch` checks the type of the data at `run time` (similar to C):

```go
type R struct { i int }
func (p *R) Get() int  { return p.i }
func (p *R) Put(v int) { p.i = v }

// use type switch to find the actual type
func f(p I) {
  switch t := p.(type) { 
    case *S: 
      // do something
    case *R: 
      // do somthing
    default: 
  }
}
```
- `.(type)` syntax is `only valid within a switch` statement

- alternative way to find out the type at run-time:

```go
if t, ok := somevar.(I); ok { // If ok is true, t will hold the type of somevar
  // ...
}
```

- interfaces can be empty: 
  - every type satisfies the empty interface: `interface{}`
  - no guarantee of any methods at all: it could contain any type

> Creating a pointer to an interface value is a useless action in Go. It is in fact illegal to create a pointer to an interface value.

- by convention, one-method interfaces are named by the `method name plus the -er suffix`: Reader, Writer, Formatter etc.

### Interfaces and Slices

- we use `interface{}` in the time of uncertainty to be later substituted by some other type

- when the substitution takes place, Go implicitly converts `interface{}` to the type `T` we need in `O(1)` time (since the time is constant, Go hides this operation)

> In Go, there is a general rule that `syntax should not hide complex/costly operations`.

- however, when we use `[]interface{}` as a substitution of the `[]T`, the conversion takes place at `O(n)` because each element of the slice must be converted to a desired type (since the time is not constant and see the general rule above, Go doesn't hide this operation) --> 

- developers need to convert each instance of []interface to []T by themselves by copying the elements individually to the destination slice. 

```go
// This example converts a slice of int to a slice of interface{}
t := []int{1, 2, 3, 4}
s := make([]interface{}, len(t))
for i, v := range t {
    s[i] = v
}

// This function converts a slice of interface{} to a slice of int
func interfaceToInt(el[]interface{}) []int{
	intSlice := make([]int, len(el))
	for i, v := range el {
		if intVar, ok := v.(int); ok { // If ok is true, t will hold the type of intVar
			intSlice[i] = intVar
		}
	}
	return intSlice
}
```

> The one exception to this rule is `converting strings`. When converting a string to and from a []byte or a []rune, Go does O(n) work even though conversions are "syntax".

#### Reasons for `T` and `interface{}` conversion

- type `T` and `interface{}` which holds a value of T have different representations in memory --> they can't be trivially converted.

  - A variable of type `T` is just its value in memory. There is no associated type information (in Go every variable has a single type known at compile time not at run time). It is represented in memory like this:
    - value


  - An `interface{}` holding a variable of type T  takes up two words  (one word for the type of what is contained, the other word for either the contained data or a pointer to it). It is represented in memory like this:
    - pointer to type T 
    - value

--> a slice with length N and with type `[]interface{}` is backed by a chunk of data that is `N*2 words` long. This is different than the chunk of data backing a slice with type `[]MyType` and the same length. Its chunk of data will be `N * sizeof(MyType) words` long.
--> 
Converting []T to []interface{} would involve creating a new slice of interface{} values which is a non-trivial operation since the in-memory layout is completely different.


- example from [stackoverflow](https://stackoverflow.com/questions/12994679/slice-of-struct-slice-of-interface-it-implements)

#### How to Solve

If you want a container for an arbitrary array type, and you plan on changing back to the original type before doing any indexing operations, you can just use an interface{}. The code will be generic (if not compile-time type-safe) and fast.

If you really want a []interface{} because you'll be doing indexing before converting back, or you are using a particular interface type and you want to use its methods, you will have to make a copy of the slice. See the example code above.


### Interfaces and Structs

- cast interface to struct using explicit casting

```go
// queue.go
// Queue node
type node struct {
	Value interface{} 
	next  *node
}

// Queue is a shape of a queue
type Queue struct {
	Len  int
	First *node
	Last *node
}

func (q *Queue) Enqueue(v interface {}) {}

// binarytree.go
// Binary tree node
type Node struct {
	Value byte
	Left *Node
	Right *Node
}

func (t *BinaryTree) CreateUQueue(data []byte) *BinaryTree {
  var q queue.Queue
  	// Add a root node to the tree
	node := &Node{Value: data[0]}
  	// Add this node to the queue
	q.Enqueue(t.Root) // we don't need to cast because interface{} can be anything
  // ...
  curr, _ := q.Dequeue() // curr is of type interface{} -->
  // Add the left node (we need to perform casting from interface to *Node because curr as interface has no member .Left)
  curr.(*Node).Left = &Node{Value: data[i]}
  // ...
}
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

## Print formatting

bool:                    `%t`
int, int8 etc.:          `%d`
uint, uint8 etc.:        `%d`, %#x if printed with %#v
float32, complex64, etc: `%g`
string:                  `%s`
chan:                    `%p`
pointer:                 `%p`
rune:                    `%c`

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

Byte:
- `%s`the uninterpreted bytes of the string or slice
- `%c`	character
 
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

## Debugging

[VS Code](https://www.youtube.com/watch?v=6r08zGi38Tk&list=PLj6YeMhvp2S40Q-TEPEKOeypLvTVd5uME&index=2&ab_channel=VisualStudioCode)

[delve](https://github.com/go-delve/delve/tree/master)

## Compilation

- `cmd/compile` contains the main packages that form the Go compiler
- name `gc` stands for `Go compiler`, whereas uppercase `GC` stands for `garbage collection`.

### Compilation Stages

You may sometimes hear the terms “front-end” and “back-end” when referring to the compiler. Roughly speaking, these translate to the first two and last two phases we are going to list here. A third term, “middle-end”, often refers to much of the work that happens in the second phase.

1. **Parsing**

- `cmd/compile/internal/syntax` (lexer, parser, syntax tree)
- source code is tokenized (lexical analysis), parsed (syntax analysis), and a syntax tree is constructed for each source file.

- each syntax tree is an exact representation of the respective source file, with nodes corresponding to the various elements of the source such as expressions, declarations, and statements. 
- the syntax tree also includes position information which is used for error reporting and the creation of debugging information.

2. **Type checking**

- `cmd/compile/internal/types2` (type checking)
- the `types2` package is a port of `go/types` to use the syntax package’s AST instead of go/ast.

3. **IR construction** (`noding`)

- `cmd/compile/internal/types` (compiler types)
- `cmd/compile/internal/ir` (compiler AST)
- `cmd/compile/internal/typecheck` (AST transformations)
- `cmd/compile/internal/noder` (create compiler AST)

- the compiler middle end uses its own AST definition and representation of Go types carried over from when it was written in C. 
- the next step after type checking is to convert the syntax and types2 representations to ir and types. This process is referred to as `noding`.

There are currently two noding implementations:

    - *irgen* (aka “-G=3” or sometimes “noder2”) is the implementation used starting with Go 1.18
    - *Unified IR* is in-development implementation (enabled with GOEXPERIMENT=unified), which also implements import/export and inlining.

> Up through Go 1.18, there was a third noding implementation (just “noder” or “-G=0”), which directly converted the pre-type-checked syntax representation into IR and then invoked package typecheck’s type checker. This implementation was removed after Go 1.18, so now package typecheck is only used for IR transformations.

4. **Middle end**

- `cmd/compile/internal/deadcode` (dead code elimination)
- `cmd/compile/internal/inline` (function call inlining)
- `cmd/compile/internal/devirtualize` (devirtualization of known interface method calls)
- `cmd/compile/internal/escape` (escape analysis)

- several optimization passes are performed on the IR representation: dead code elimination, (early) devirtualization, function call inlining, and escape analysis.

5. **Walk**

- `cmd/compile/internal/walk` (order of evaluation, desugaring)

The final pass over the IR representation is `walk`, which serves two purposes:

- it decomposes complex statements into individual, simpler statements, introducing temporary variables and respecting order of evaluation. This step is also referred to as `order`.

- it desugars higher-level Go constructs into more primitive ones. For example, switch statements are turned into binary search or jump tables, and operations on maps and channels are replaced with runtime calls.

6. **Generic SSA**

- `cmd/compile/internal/ssa` (SSA passes and rules)
- `cmd/compile/internal/ssagen` (converting IR to SSA)

- in this phase, IR is converted into `Static Single Assignment` (SSA) form, a lower-level intermediate representation with specific properties that make it easier to implement optimizations and to eventually generate machine code from it.

- during this conversion, `function intrinsics` are applied. These are special functions that the compiler has been taught to replace with heavily optimized code on a case-by-case basis.

- certain nodes are also lowered into simpler components during the AST to SSA conversion, so that the rest of the compiler can work with them. For instance, the copy builtin is replaced by memory moves, and range loops are rewritten into for loops. Some of these currently happen before the conversion to SSA due to historical reasons, but the long-term plan is to move all of them here.

- then, a series of machine-independent passes and rules are applied. These do not concern any single computer architecture, and thus run on all GOARCH variants. These passes include dead code elimination, removal of unneeded nil checks, and removal of unused branches. The generic rewrite rules mainly concern expressions, such as replacing some expressions with constant values, and optimizing multiplications and float operations.

7. **Generating machine code**

- `cmd/compile/internal/ssa` (SSA lowering and arch-specific passes)
- `cmd/internal/obj` (machine code generation)

- the machine-dependent phase of the compiler begins with the “lower” pass, which rewrites generic values into their machine-specific variants. For example, on amd64 memory operands are possible, so many load-store operations may be combined.

- the lower pass runs all machine-specific rewrite rules, and thus it currently applies lots of optimizations too.

- once the SSA has been “lowered” and is more specific to the target architecture, the final code optimization passes are `run`. This includes yet another dead code elimination pass, moving values closer to their uses, the removal of local variables that are never read from, and register allocation.

- other important pieces of work done as part of this step include `stack frame layout`, which assigns stack offsets to local variables, and `pointer liveness analysis`, which computes which on-stack pointers are live at each GC safe point.

- at the end of the SSA generation phase, Go functions have been transformed into a series of obj.Prog instructions. These are passed to the assembler (cmd/internal/obj), which turns them into machine code and writes out the final object file. The object file will also contain reflect data, export data, and debugging information.

> To dig deeper into how the SSA package works, including its passes and rules, head to `cmd/compile/internal/ssa/README.md`.

> The `go/*` family of packages, such as `go/parser` and `go/types`, are mostly unused by the compiler. Since the compiler was initially written in C, the go/* packages were developed to enable writing tools working with Go code, such as gofmt and vet. Over time the compiler’s internal APIs have slowly evolved to be more familiar to users of the go/* packages.

## Tools

[linter](https://golangci-lint.run/)

Generate unit tests using VS Code Go Extension:

1. `SHIFT + CMD + P`
2. Go:Generate Unit Tests ... (choose what is needed)
