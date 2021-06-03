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
- we can return a struct or interface from a factory function
- can be combined with prototype pattern

4. `Prototype`:
- complicated objects are not designed from scratch, they reiterate existing designs
- an existing (partially constructed) design is a prototype
- we make a copy of the prototype and customise it (requires `deep copy` support)
- a prototype is a partially or fully initialised object that you copy (clone) and make use of it

5. `Singleton`:
- for some componentes it only makes sense to have one in the system:
    - database repository
    - object factory
- in cases when the construction call is expensive
    - we only do it once
    - we give everyone the same instance
- we need to prevent creation of additional copies
- we want to have a lazy instantiation
- often breaks the dependency inversion principle (in singleton, we often depend on the concrete implementation of singleton instead of abstraction)