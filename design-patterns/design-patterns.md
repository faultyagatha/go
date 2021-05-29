## DESIGN PATTERNS

## Creational:
1. `Singleton`:

2. `Builder`:
- some objects are simple and can be created in a single constructor call
- other objects require more efforts to create
- having a factory function with 10 arguments is not productive, because we force the user of the API to make many decisions at once
- builder pattern provides an API for constructing an object step by step
- to make a builder fluent in Go (allow chaining), we can return the receiver or a pointer to the receiver
- in some situation, we might need more than one builders (see builder/facets.go)
- functional builder allows easily to extend the builder and add more modifications to it

3. `Factory`:
- way of controlling how an object is constructed 
- used when the object creation logic gets too convoluted or when the struct has too many fields and needs to initialise all them correctly
- the object is created in a single invocation that can be outsourced to:
    - a separate function (`Constructor`, `Factory Function`)
    - a separate factory struct
- we can return struct or interface from a factory function
s
4. `Prototype`:
...