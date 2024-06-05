# Let-s-Go
This repository contains my learnings for each day about the fundamentals of the Go language.

## Day 1
- **What is GO language?**
    - Go is a procedural programming language. It was developed in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google but launched in 2009 as an open-source programming language.Programs are assembled by using packages, for efficient management of dependencies.It is a statically typed, concurrent and garbage-collected programming language.
- **Inspiration**
    - Modern programming languages are getting complicated each day. It seems such that the languages are taking features from each other and converging into a single language. But this convergence comes with the cost of complexity. Although it may seem like each feature is making it simpler to code but they are also adding the additional burden of complexity behind the scenes.As Rob pike describes it, "Bloat without distinction". Additionally, there are multiple ways to do the same thing. One developer might traverse through a loop in one way and another developer might do it in another way. This makes the code less readable for the develoeprs. Go's inspiration was to make a language that is simple, readable and easy to maintain.
- **Feature Overview**
    - **Simplicity**: Go is simple to read and write. It has a clean and easy to understand syntax.
    - **Fast Compile Times**: Go is a compiled language. It is compiled to machine code. The compilation time is very fast.
    - **Garbage Collection**: Go has garbage collection. It automatically manages memory.
    - **Concurrency**: Go has goroutines which are lightweight threads. They are used to run functions concurrently. Also there are channels to help communicate between goroutines.
    - **Static Typing**: Go is statically typed. It helps in catching errors at compile time.
    - **Cross-Platform**: Go is cross-platform. It can be compiled on different platforms.
    - **Packages**: Go has a package system which makes it easy to manage dependencies.
    - **Interface**: Go has interfaces which can be implemented by struct and non-struct types.
    These features will be discussed in details below.
- **Principles**
    Much like objected oriented programming languages , Go also has a set of principles that it follows.
    - **Don't communicate by sharing memory, share memory by communicating** 
        Go uses channels to communicate between different goroutines.This is in contrast to traditional threading models used in different languages where the threads communicate using share memory. Instead of doing that, passing references to data from one goroutine to another ensures that only one goroutine has access to the data at a given time.
    - **The bigger the interface, the weaker the abstraction**
        The smaller the interface is , the more helpful it is. If an interface has only one method, it is more helpful than an interface with multiple methods. This is because the smaller the interface, the more specific it is to the type that it is implemented on.In Go, interfaces are not implemented explicitly rather they are implementedimplicitly.If a type implements all the methods of an interface it means that the interface is implemented by the type."If a bird eats like a duck,swims like a duck and quacks like a duck, then it is a duck".
    - **Errors are values**
        Treating errors as values makes it easier and cleaner to handle errors. In Go, errors are values that can be returned from functions. If a function returns an error, it is the responsibility of the caller to handle the error.
- **How a Go Code is Compiled**
    When a Go code is compiled it is first transformed to bytecode. This bytecode is compiled to machine code by the Go compiler. The resutting machine code is stored in a executable file which can then be loaded into memory by the Go loader. While loading the executable to memory it does various tasks such as mapping to memory, initializing data structures required, security checks etc.
- **Features**
    - Simplicity
        Go is simple to read and write.There are only 20 keywords present in the language. There are lesser built-in data types. This makes Go an easy to understand language. The syntax is clean and easy so that any developer can just tell what is happening in the code by seeing it. In other languages there are a lot of ways to do the same task. Go wants to get rid of that confusion. Instead the syntax is defined in such a way so that developers can easily read and write code.
    - Fast Compile Time
        Go is very well known for fast compile times. It is obviously faster than interpreted languages but it is also faster than a lot of compiled languages like C,C++ . Let's discuss how Go achieves fast compile times.
        - Dependency management
            Go has a package based system. This enables handling dependencies easily.Go does not allow unused dependencies and circular dependencies.
        - No symbol table for parsing
            Compiled languages like C++ often use symbol tables for compiling and also parsing the syntax to make a parse tree.Go does not use symbol tables for parsing. This makes the parsing faster than C,C++.
        - No header files
            Instead of using header files, Go uses a package system. This makes it easier to manage dependencies and thus improve performance.
        - No virtual machine
            Java and C# use virtual machines to run the code. Go does not use a virtual machine. It compiles the code to machine code directly.This makes it portable to different platforms as well as improves the compilation time as there is no middle man needed.
    - Packages
         Go is not an object oriented language.It uses packages to handle dependencies.Packaging is also done in other programming languages but we can choose to not do it if we don't want to.Instead,in Go it is a feature of the language itself. It can be compared to implementing the idea of microservices in source code level.We can think of each package as a different micro service which can be accessed through an api.
         Every package should have a purpose and its name should be very specific to what it does.It can be compared to the single responsibility principle in OOP where each method should have only one responsibility.In this case each package should have a very specific purpose.
         Each package should respect its impact on resources and performance of the user.
         Packages should be designed while keeping reusability in mind.This makes a package protable.Each package should be as decoupled as possible.