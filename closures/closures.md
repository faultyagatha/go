# Closures

`Closures` allow us to encapsulate and maintain the state of variables alongside the functions that use them. 

Closures transcend language boundaries, showcasing how different programming paradigms and design philosophies influence their implementation. Understanding how closures work in Go, JavaScript, and C++ provides us with a broader perspective on how these languages tackle the challenges of maintaining state and context within functions.


In Go, closures are implemented as anonymous functions that can capture variables from their lexical scope. The function pointer to a closure, along with any captured variables, is stored on the heap. This enables closures to outlive the scope in which they were created, ensuring the captured variables' values are preserved. Go's garbage collector manages the memory, releasing the resources when the closure and its captured variables are no longer referenced.


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

JavaScript implements closures by maintaining references to the variables within the closure's lexical scope. Function objects, including their internal [[Environment]] reference, are stored in memory. Captured variables are stored in the environment, allowing them to persist even after the parent function has finished executing. JavaScript's garbage collector automatically manages memory, ensuring that closures and their captured variables are cleaned up when no longer needed.

```js
function createMultiplier(factor) {
  return function(x) {
      return x * factor;
  };
}
const double = createMultiplier(2);
console.log(double(5)); // Output: 10
```

In C++, closures are introduced through lambda expressions, which are essentially syntactic sugar for function objects. The function pointer to a lambda, along with its captured variables, is stored on the stack or heap, depending on how the lambda is used. Captured variables are stored within the lambda object's context. The lifetime of the function and captured variables depends on where they are stored - stack-based closures go out of scope when the block ends, while heap-based closures are managed manually or through smart pointers.

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