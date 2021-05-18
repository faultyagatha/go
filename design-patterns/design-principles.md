## SOLID Principles:

- S - SRP (Single Responsibility Principle)
- O - OCP (Open Closed Principle)
- L - LSP (Liskov Substitution Principle)
- I - ISP (Interface Segregation Principle)
- D - DIP (Dependency Inversion Principle)

1. `SRP` (Single Responsibility Principle)

Every function, class, module, or service should have a single clearly defined responsibility. In other words, a class/function/module should have one and only one reason to change.

- makes the code more organized
- improves the readability of the code
- contributes to reusability: if you have short and focused functions/classes, you’ll be able to reuse them easily. But if you have a single function that does too many things then you wouldn’t be able to use it if you only need a fraction of the functionality implemented by the function.

Antipattern: `God Object`: object or package that packed with everything

2. `OCP` (Open/Closed Principle)

We should build our functions/classes/modules in such a way that they are open for extension, but closed for modification.

- Open for Extension

We should be able to add new features to the classes/modules without breaking existing code. This can be achieved using inheritance and composition.

- Closed for Modification

We should strive not to introduce breaking changes to the existing functionality, because that would force us to refactor a lot of existing code and write a whole bunch of tests to make sure that the changes work.

3. `LSP` (Liskov Substitution Principle)

Every child/derived class should be substitutable for their parent/base class without altering the correctness of the program. In other words, the objects of your subclass should behave in the same way as the objects of your superclass.

```java
class Bird {
    public void fly() {
        System.out.println("Bird flying...");
    }
    public void eat() {
        System.out.println("Bird eating...");
    }
}

class Penguin extends Bird {
    public void fly() {
       throw new UnsupportedOperationException("Can't fly.");
    }
}

class LSPTest {
    public static void main(String[] args) {
        Bird bird = new Bird();
        bird.fly();
    }
}
```

According to LSP, if you have a piece of code that uses a Bird object, then you should be able to replace it with a Penguin object and the code will still behave the same.

But, the above example violates the Liskov Substitution Principle. You can’t replace an object of the Bird class with an object of the Penguin class. If you do that, the program will throw an exception.

To fix:

```java
class Bird {
    public void eat() {
        System.out.println("Bird eating...");
    }
}

class FlightBird extends Bird {
    public void fly() {
        System.out.println("Bird flying...");
    }
}

class FlightlessBird extends Bird {

}
```

4. `ISP` (Interface Segregation Principle)

A client should never be forced to depend on methods it does not use. Make your interfaces small and focused.

You should split large interfaces into more specific ones that are focused on a specific set of functionalities so that the clients can choose to depend only on the functionalities they need.

5. `DIP` (Dependency Inversion Principle)

Avoid tight coupling between software modules. High-level modules should not depend on low-level modules, but only on their abstractions. In simple words, you should use interfaces instead of concrete implementations wherever possible.

This decouples a module from the implementation details of its dependencies. The module only knows about the behavior on which it depends, not how that behavior is implemented. This allows you to change the implementation whenever you want without affecting the module itself.
