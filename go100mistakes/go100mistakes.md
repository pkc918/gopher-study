
### Unintended variable shadowing

1. Variable shadowing occurs when a variable name is redeclared in an inner block.

### Unnecessary nested code

1. Striving to reduce the number of nested blocks, aligning the happy path on the left, and returning as early as possible are concrete means to improve our code's readability.

### Misusing init functions

1. When a package is initialized, all the constant and variable declarations in the package are evaluated. Then, the init functions are executed.
2. We can define multiple init functions per package. When we do, the execution order of the init function inside the package is based on the source files' alphabetical order.
3. Global variables have some server drawbacks; for example:
    - Any functions can alter global variables within the package.
    - Unit tests can be more complicated because a funstion that depends on a global variable won't be isolated anymore.

4. We should be cautions with init functions. They can be helpful in some situations, however, such as defining static configuration, as we saw in this section. Otherwise, and in most cases, we should handle initialization through ad hoc functions.

### Overusing getters and setters

1. Using getters and setters presents some advantages, including these:
    - They encapsulate a behavior associated with getting or setting a field, allowing new functionality to be added later(for example, validating a field, returning a computed value, or wrapping the access to a field around a mutex).
    - They hide the internal representation, giving us more flexibility in what we wxpose.
    - They provide a debugging interception point for when the property changes at run time, making debugging easier.

2. Remember that Go is a unique language designed for many characteristics, including simplicity. However, if we find a need for getters and setters or, as mentioned, foresee a future need while guaranteeing forward compatibility, there's nothing wrong with using them.

### Interface pollution

1. Interface pollution is about overwhelming our code with unnecessaryadstractions, making it harder to understand. It's a common mistake made by developers coming from another language with different habits.
2. Let's look at three concrete use cases where interfaces are usually considered to bring value. Note that the goal isn't to be exhaustive because the more cases we add, the more they would depend on the context. However, these three cases should give us a general idea:

    - Common behavior
    - Decoupling
    - Restricting behavior

#### Common behavior

- Retrieving the number of elements in the collection
- Reporting whether one element must be sorted before another
- Swapping two elements

#### Decoupling
To give us more flexibility, we should decouple `CustomerService` from the actual implementation, which can be done via an interface.

Because storing a customer is now done via an interface, this gives us more flexibility in how we want to test the method. For instance, we can 

- Use the concrete implementation via integration tests
- Use a mock (or any kind of test double) via unit tests
- Or both

#### Restricting behavior
Therefore, we can also use interfaces to restrict a type to a specific behavior for various reasons, such as semantics enforcement.

2. What's the main problem if we overuse interfaces? The answer is that they make the code flow more complex. Adding a useless level of indirection doesn't bring any value; it creates a worthless abstraction mking the code more difficult to read, understand, and reason about. If we don't have a strong reason for adding an interface and it's unclear how an interface makes a code better, we should challenge this interface's purpose. In summary, we should be cautious when creating abstractions in our code--abstractions should be discovered, not created.
3. Don't design with interfaces, discover them.
4. If it's unclear how an interface makes the code better, we should probably consider removing it to make our code simpler.

### Interface on the producer side

An interface should live on the consumer side in most cases. However, in particylar contexts(for example, when we know——not foresee——that an abstraction will be helpful for consumers), we may want to have it on the producer side. If we do, we should strive to keep it as minimal as possible, increasing its reusability potential and making it more easily composable.

### Returning interfaces
> Be conservative in what you do, be liberal in what you accept from others. —— Transmission Control Protocol

If we apply this idiom to Go, it means
- Returning structs instead of interfaces
- Accepting interfaces if possible

All in all, in most cases, we shouldn't return interfaces but concrete implementations. Otherwise, it can make our design more complex due to package dependencies and can restrict flexibility because all the clients woyld have to rely on the same abstraction. Again, the conclusion is similar to the previous sections: if we know (not foresee) that an abstraction will be helpful for clients, we can consider returning an interface. Otherwise, we shouldn't force abstractions; they should be discovered by clients. If a client needs to abstract an implementation for whatever reason, it can still do that on the client's side.