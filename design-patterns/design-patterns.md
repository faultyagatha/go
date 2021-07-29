## DESIGN PATTERNS

## Creational:

1. `Builder`:
- some objects are simple and can be created in a single constructor call
- other objects require more efforts to create
- having a factory function with 10 arguments is not productive, because we force the user of the API to make many decisions at once
- builder pattern provides an API for constructing an object step by step
- to make a builder fluent in Go (allow chaining), we can return the receiver or a pointer to the receiver
- in some situation, we might need more than one builders (see builder/facets.go)
- functional builder allows easily to extend the builder and add more modifications to it

2. `Factory`:
- way of controlling how an object is constructed
- used when the object creation logic gets too convoluted or when the struct has too many fields and needs to initialise all them correctly
- the object is created in a single invocation that can be outsourced to:
    - a separate function (`Constructor`, `Factory Function`)
    - a separate factory struct
- we can return a struct or interface from a factory function
- can be combined with prototype pattern

3. `Prototype`:
- complicated objects are not designed from scratch, they reiterate existing designs
- an existing (partially constructed) design is a prototype
- we make a copy of the prototype and customise it (requires `deep copy` support)
- a prototype is a partially or fully initialised object that you copy (clone) and make use of it

4. `Singleton`:
- for some componentes it only makes sense to have one in the system:
    - database repository
    - object factory
- in cases when the construction call is expensive
    - we only do it once
    - we give everyone the same instance
- we need to prevent creation of additional copies
- we want to have a lazy instantiation
- often breaks the dependency inversion principle (in singleton, we often depend on the concrete implementation of singleton instead of abstraction)

5. `Adapter`:
- a construct that adapts an existing interface X to conform the required interface Y
- to implement the adapter:
    - determine the API you have and the API you need
    - create a component that aggregates (has a pointer to) the adaptee
    - intermediate representations can pile up: use caching and other optimisations

6. `Facade`:
- provides a simple and easy to understand API or user interface over a large and sophisticated body of code
- we can expose these detailes if we want

7. `Bridge`:
- solves a 'Cartesian product' complexity exposion problem
- decouples abstraction from implementation
- both abstraction and implementation can exist as hierarchies
- stronger form of incapsulation
- can result in a cascading set of methods on a Bridge when many objects are introduced

8. `Composite`:
- objects use other objects' fields/methods via embedding
- composition allows for making compound objects
- composite design is used to treat both single and composite objects uniformly

9. `Decorator`:
- a pattern that facilitates the addition of behaviours to individual objects through embedding
- solves the problem when we want to modify an object with additional functionality but do not want to rewrite or alter existing code (OCP) and keep new functionality separate (SRP)
- interacts with existing structures
- often used to emulate multiple inheritance


## Behavioral:
1. `Iterator`:
- iterator is a type that facilitates the traversal
    - it keeps a pointer to the current element
    - it knows how to move to a different element
- solves the problem of accessing and traversing the elements of an aggregate object without exposing its representation (data structures)
- with iterator, new traversal operations can be defined for an aggregate object without changing its interface
- iterator pattern defines a separate (iterator) object that encapsulates accessing and traversing an aggregate object
- iterator moves along the iterated collection, indicating when last element has been reached
- in go, there is no standard Iterable interface; iteration can be achieved with `range` - in-built iterator for iterable objects

2. `Observer`:
- solves the need to be informed when certain things happen
    - object's field changes
    - object does something
    - some external event takes place
- we want to listen to events and be notified when they take place
- two participants: observable and observer
- observer is an object that wishes to be informed about events happening in the system
- observable is the entity generating the events
- common application in Go: send a notification when any property of a struct has changed

2. `State`:
- a pattern in which the object's behaviour is determined by its state
- changes in state can be explicit or in response to event (see observer pattern)
- an object transitions from one state to another (a trigger, action)
- can define:
    - state entry/exit behaviours
    - action when a particular event causes transition
    - guard conditions enabling/disabling a transition