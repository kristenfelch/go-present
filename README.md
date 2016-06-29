# Functional Programming - The Good, the Bad, and the Try at Your Own Risk, 6/29

## My Background

Somewhere between theoretical computer scientist and eager application developer.

Scheme (MIT) -> Java (EMC) -> Web Stack (InvisionApp) -> GoLang (Den)

## My Expectations of GoLang

Trends in other languages seem to be introducing more FP concepts
 - Java 7/8 : Lambda functions and Collections Library
 - Birth of Scala : 2004 - Functional Language on the JVM
 - ES6 and Beyond : Map/Reduce/Filter, Arrow functions
 
## Brief Overview of Functional Programming
 - A Programming Paradigm

### Tenets
 - Computation is the evaluation of mathematical functions
 - Immutable Data and Stateless applications (create versus update)
 - Declarative, using expression evaluation rather than statements.
 
### Characteristics 
 - First class and higher order functions (importance placed on functions and their evaluation,
 rather than objects and updating them)
 - Stateless applications (since we cannot mutate data, or state variables)
 - Pure functions - No side effects to functions (no state)
 - Referential Transparency (no state -> functions of same input always lead to same output)
 - Recursive algorithms rather than iterative. (no state, no counter)
 
### Advantages
 
 - Modularity, of a different kind than OOP.
 - Performance improvements:
     - Memoization - function results with same arguments can be cached
     - Parallelization - no state, so we can perform parts of algorithm in parallel
     - Compiler has freedom to reorder or combine the evaluation of expressions in a program
     - No monitors, race conditions, deadlocks.
 - Testability, based on modularity.  No state to be replicated or mocked.
     - Test the data separate from the math.
 - Learn a new way of thinking.
     
## Takeaways

 - Iteration should be used rather than recursion, due to lack of Tail Recursion Optimization.
 - Higher order functions - Map/Reduce/Filter can be written, but there is a small performance
  tradeoff as well as the necessity to write new functions for each data type.
 - Pure functions with very specific functionality are easy to read, reusable, and allow us to
 take advantage of parallelism via goroutines.

 
 
 
