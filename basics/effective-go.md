# Effective Go

## Packages

- lower case, single-word names; there should be no need for underscores or mixedCaps
- the package name is the base name of its source directory; the package in src/encoding/base64 is imported as "encoding/base64" but has name base64, not encoding_base64 and not encodingBase64.

## Getters
- no support for getters and setters but it's ok to have them
- idiomatic practice:
  - no `Get` into the getter's name
  - e.g., if you have a field called `owner` (lower case, unexported), the getter method should be called `Owner` (upper case, exported), not GetOwner. 
  - the use of upper-case names for export provides the hook to discriminate the field from the method
  - a setter function, if needed, will likely be called SetOwner

  ```go
  owner := obj.Owner()
  if owner != user {
    obj.SetOwner(user)
  }
  ```

  ## Interface

  - one-method interfaces are `named` by the `method name plus an -er` suffix or similar modification to construct an agent noun: Reader, Writer, Formatter, CloseNotifier etc.
  - follow the same naming patters (as in standard lib)
  - use MixedCaps or mixedCaps rather than underscores to write multiword names

  ## Semicolons

  - if the last token before a newline is an identifier (words like int and float64), a basic literal such as a number or string constant, or one of the tokens
```go
break continue fallthrough return ++ -- ) }
```
the lexer always inserts a semicolon after the token

- you cannot put the opening brace of a control structure (if, for, switch, or select) on the next line. If you do, a semicolon will be inserted before the brace, which could cause unwanted effects (!!)

## Switch

- more general than C's
- the expressions need not be constants or even integers, the cases are evaluated top to bottom until a match is found, and if the switch has no expression it switches on true. --> 
- It is idiomatic—to write an if-else-if-else chain as a switch.

```go
func unhex(c byte) byte {
  switch {
    case '0' <= c && c <= '9':
      return c - '0'
    case 'a' <= c && c <= 'f':
      return c - 'a' + 10
    case 'A' <= c && c <= 'F':
      return c - 'A' + 10
  }
  return 0
}
```

- no automatic fall through, but cases can be presented in comma-separated lists

```go
func shouldEscape(c byte) bool {
  switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
      return true
  }
  return false
}
```

- not necessary to use `break` but can be used; the same applies to `continue`

- it is possible to break out of a surrounding loop, not the switch, by putting a label on the loop and "breaking" to that label

```go
Loop:
  for n := 0; n < len(src); n += size {
    switch {
    case src[n] < sizeOne:
      if validateOnly {
        break
      }
      size = 1
      update(src[n])
    case src[n] < sizeTwo:
      if n+1 >= len(src) {
        err = errShortInput
        break Loop
      }
      if validateOnly {
        break
      }
      size = 2
      update(src[n] + src[n+1]<<shift)
    }
  }
```

- `type switch` uses the syntax of a type assertion with the keyword type inside the parentheses

```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
  default:
    fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
  case bool:
    fmt.Printf("boolean %t\n", t)             // t has type bool
  case int:
    fmt.Printf("integer %d\n", t)             // t has type int
  case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
  case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}

```


[effective go](https://go.dev/doc/effective_go)
