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

3. `Factory`:

4. `Prototype`:
...