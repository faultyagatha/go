# Closures

`Closures` allow us to encapsulate and maintain the state of variables alongside the functions that use them. 

## Implementation Across Languages

### Go
- closures are implemented as anonymous functions that can capture variables from their lexical scope
- the function pointer to a closure, along with any captured variables, is stored on the heap
- this enables closures to outlive the scope in which they were created, ensuring the captured variables' values are preserved
- Go's garbage collector manages the memory, releasing the resources when the closure and its captured variables are no longer referenced.

```go
func main() {
  x := 10
  addX := func(y int) int {
    return x + y
  }
  result := addX(5)
  fmt.Println(result) // Output: 15
}
```

### JavaScript 
- implements closures by maintaining references to the variables within the closure's lexical scope. 
- function objects, including their internal [[Environment]] reference, are stored in memory. 
- captured variables are stored in the environment, allowing them to persist even after the parent function has finished executing. 
- JavaScript's garbage collector automatically manages memory, ensuring that closures and their captured variables are cleaned up when no longer needed.

```js
function createMultiplier(factor) {
  return function(x) {
      return x * factor;
  };
}
const double = createMultiplier(2);
console.log(double(5)); // Output: 10
```

### C++
- closures are introduced through lambda expressions, which are essentially syntactic sugar for function objects. 
- the function pointer to a lambda, along with its captured variables, is stored on the stack or heap, depending on how the lambda is used. 
- captured variables are stored within the lambda object's context.
- the lifetime of the function and captured variables depends on where they are stored - stack-based closures go out of scope when the block ends, while heap-based closures are managed manually or through smart pointers.

```c++
#include <iostream>
#include <functional>

int main() {
  int x = 10;
  std::function<int(int)> addX = [&x](int y) {
    return x + y;
  };
  int result = addX(5);
  std::cout << result << std::endl; // Output: 15
}

```

## Closures in Go

- used as anonymous functions, named callbacks, etc
- can be immediately called (as in JS)
- Important (!!): `Don't Capture Indexers!!`

> The `value of a variable in a closure` is the `value at the time the variable is used`, not the value at the time it was captured.

```go
ch := make(chan person, 10)
var wg sync.WaitGroup

// put values onto a channel
for _, id := range ids {
  wg.Add(1)
  go func() {
    defer wg.Done()
    p, err := getPerson(id)
    if err != nil {
      // Here the value will always be 9 (with some exceptions of concurrency).
      // It happens because of the way closures work in Go:
      // if we capture an indexer or an iterator, we often get the
      //  final value of that indexer.
      log.Printf("err calling getPerson(%d): %v", id, err)
      return
    }
    ch <- p
  }()
}

wg.Wait()
close(ch)

// read values from the channel
for p := range ch {
  fmt.Printf("%d: %v\n", p.ID, p)
}
```


