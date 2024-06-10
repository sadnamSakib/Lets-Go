# Let's Go
This repository contains my learnings for each day about the fundamentals of the Go language.

- ## **What is GO language?**
    - Go is a procedural programming language. It was developed in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google but launched in 2009 as an open-source programming language.Programs are assembled by using packages, for efficient management of dependencies.It is a statically typed, concurrent and garbage-collected programming language.
- ## **Inspiration**
    - Modern programming languages are getting complicated each day. It seems such that the languages are taking features from each other and converging into a single language. But this convergence comes with the cost of complexity. Although it may seem like each feature is making it simpler to code but they are also adding the additional burden of complexity behind the scenes.As Rob pike describes it, "Bloat without distinction". Additionally, there are multiple ways to do the same thing. One developer might traverse through a loop in one way and another developer might do it in another way. This makes the code less readable for the develoeprs. Go's inspiration was to make a language that is simple, readable and easy to maintain.
- ## **Feature Overview**
    - **Simplicity**: Go is simple to read and write. It has a clean and easy to understand syntax.
    - **Fast Compile Times**: Go is a compiled language. It is compiled to machine code. The compilation time is very fast.
    - **Garbage Collection**: Go has garbage collection. It automatically manages memory.
    - **Concurrency**: Go has goroutines which are lightweight threads. They are used to run functions concurrently. Also there are channels to help communicate between goroutines.
    - **Static Typing**: Go is statically typed. It helps in catching errors at compile time.
    - **Cross-Platform**: Go is cross-platform. It can be compiled on different platforms.
    - **Packages**: Go has a package system which makes it easy to manage dependencies.
    - **Interface**: Go has interfaces which can be implemented by struct and non-struct types.
    These features will be discussed in details below.
    - **Structs**: Go has structs which are used to define custom data types. They are similar to classes in object oriented programming languages.
- ## **Principles**
    Much like objected oriented programming languages , Go also has a set of principles that it follows.
    - ### **Don't communicate by sharing memory, share memory by communicating** 
        Go uses channels to communicate between different goroutines.This is in contrast to traditional threading models used in different languages where the threads communicate using share memory. Instead of doing that, passing references to data from one goroutine to another ensures that only one goroutine has access to the data at a given time.
    - ### **The bigger the interface, the weaker the abstraction**
        The smaller the interface is , the more helpful it is. If an interface has only one method, it is more helpful than an interface with multiple methods. This is because the smaller the interface, the more specific it is to the type that it is implemented on.In Go, interfaces are not implemented explicitly rather they are implementedimplicitly.If a type implements all the methods of an interface it means that the interface is implemented by the type."If a bird eats like a duck,swims like a duck and quacks like a duck, then it is a duck".
    - ### **Errors are values**
        Treating errors as values makes it easier and cleaner to handle errors. In Go, errors are values that can be returned from functions. If a function returns an error, it is the responsibility of the caller to handle the error.
- ## **How a Go Code is Compiled**
    When a Go code is compiled it is first transformed to bytecode. This bytecode is compiled to machine code by the Go compiler. The resutting machine code is stored in a executable file which can then be loaded into memory by the Go loader. While loading the executable to memory it does various tasks such as mapping to memory, initializing data structures required, security checks etc.
- ## **Features**
    - ### **Simplicity** :
        Go is simple to read and write.There are only 20 keywords present in the language. There are lesser built-in data types. This makes Go an easy to understand language. The syntax is clean and easy so that any developer can just tell what is happening in the code by seeing it. In other languages there are a lot of ways to do the same task. Go wants to get rid of that confusion. Instead the syntax is defined in such a way so that developers can easily read and write code.
    - ### **Fast Compile Time** : 
        Go is very well known for fast compile times. It is obviously faster than interpreted languages but it is also faster than a lot of compiled languages like C,C++ . Let's discuss how Go achieves fast compile times.
        - Dependency management
            Go has a package based system. This enables handling dependencies easily.Go does not allow unused dependencies and circular dependencies.
        - No symbol table for parsing
            Compiled languages like C++ often use symbol tables for compiling and also parsing the syntax to make a parse tree.Go does not use symbol tables for parsing. This makes the parsing faster than C,C++.
        - No header files
            Instead of using header files, Go uses a package system. This makes it easier to manage dependencies and thus improve performance.
        - No virtual machine
            Java and C# use virtual machines to run the code. Go does not use a virtual machine. It compiles the code to machine code directly.This makes it portable to different platforms as well as improves the compilation time as there is no middle man needed.
    - ### **Packages** : 
         Go is not an object oriented language.It uses packages to handle dependencies.Packaging is also done in other programming languages but we can choose to not do it if we don't want to.Instead,in Go it is a feature of the language itself. It can be compared to implementing the idea of microservices in source code level.We can think of each package as a different micro service which can be accessed through an api.
         Every package should have a purpose and its name should be very specific to what it does.It can be compared to the single responsibility principle in OOP where each method should have only one responsibility.In this case each package should have a very specific purpose.
         Each package should respect its impact on resources and performance of the user.
         Packages should be designed while keeping reusability in mind.This makes a package portable.Each package should be as decoupled as possible.
    - ### **Concurrency** : 
        One of the highlighting features of Go is its concurrency. Concurrent programming is not unique to Go ,other programming languages use threads to handle concurrency. But Go uses goroutines which are basically lightweight threads, therefore we can create thousands of these if we want. Now let's take a look at how concurrency is handled in Go.
        - Goroutines
                A goroutine is a function that is capable of running concurrently with other functions. To create a goroutine we use the keyword go followed by the function name. This will create a new goroutine which will run concurrently with the main goroutine.

                              
                package main
                import "fmt"

                func f(n int) {
                for i := 0; i < 10; i++ {
                    fmt.Println(n, ":", i)
                }
                }

                func main() {
                go f(0)
                var input string
                fmt.Scanln(&input)
                }

      
       Normally if we had just called the function f(0) it would have run in the main goroutine. But by using go f(0) we are creating a new goroutine which will run concurrently with the main goroutine.
            - Channels
                Channels provide a way for two goroutines to communicate with one another and synchronize their execution.
                Some rules that channels follow:
                - A channel type is represented with the keyword chan
                - the <- operator is used to send and receive messages from the channel
                - Suppose a <- b , b waits until a is ready to receive the message from b
                - We can specify channel direction which can be send only, receive only or both
                - Channels are asynchronous unless they are buffered channels. Buffered channels are asynchronous and does not wait for sending or recieveing messages unless the channel is full.Buffered channels are created using a second paramater inside the make function.
            - Select
                
        Normally if we had just called the function f(0) it would have run in the main goroutine. But by using go f(0) we are creating a new goroutine which will run concurrently with the main goroutine.
        - Channels
            Channels provide a way for two goroutines to communicate with one another and synchronize their execution.
            Some rules that channels follow:
            - A channel type is represented with the keyword chan
            - the <- operator is used to send and receive messages from the channel
            - Suppose a <- b , b waits until a is ready to receive the message from b
            - We can specify channel direction which can be send only, receive only or both
            - Channels are asynchronous unless they are buffered channels. Buffered channels are asynchronous and does not wait for sending or recieveing messages unless the channel is full.Buffered channels are created using a second paramater inside the make function.
        - Select
                The select statement works much like a switch case statement but for channels. It waits for any of the cases to be ready and executes the corresponding case. If multiple cases are ready, it randomly selects one of them.
    - ### **Garbage Collection** : 
        Go is a garbage collected language. This means that the memory is managed automatically by the Go runtime. The garbage collector runs in the background and reclaims memory that is no longer being used by the program. Languages like C don't heave this feature. In those languages the developer has to manually allocate and deallocate memory. This can lead to memory leaks and other problems.Go uses Tricolor mark and Sweep algorithm for garbage collection.
    - ### **Static Typing** : 
        Go is a statically typed language.The types are checked at compile time rather than runtime.This is another reason for which the compilation is faster. Even after being strictly typed, Go allows for type inference. This means that we don't have to explicitly mention the type of a variable. The compiler can infer the type of the variable from the value assigned to it at compile time.
        This feature makes it safer to handle errors as the compiler can catch type errors at compile time itself.Go has three basic data types which are bool, numeric and string. Numeric represents integer,floating point and complex types.
    - ### **Cross platform** : 
        Go is a cross platform language. This means that we can compile the code on one platform and run it on another platform. This is possible because Go compiles the code to machine code. This makes it easier to run the code on different platforms.There are other languages like Java and C# which are also cross platform but they use virtual machines to run the code. Go does not use a virtual machine. It compiles the code to machine code directly.
    - ### **Interface** : 
        Go is not an object oriented language and therefore it does not have the concept of classes or inheritance.Instead Go uses interfaces to achieve polymorphism. An interface is a set of methods that a type must implement if the type has to implement the interface. In Go we don't need to explicity implement an interface. It is done implicitly by the language if a type has implemented all the methods of an interface.There is also a concept of empty interface which is an interface with no methods. This can be used to store any type of value.
    - ### **Structs** : 
        Like C and C++ Go has a struct type. A struct is a collection of fields. It is used to define custom data types. It is similar to classes in object oriented programming languages except that it does not have any methods.
- ## **Comparison with other languages**
    In this section we want to compare Go with some other languages like Java, Python and C++.
    |  | Go | Java | Python | C++
    |--------|-|--------|------|----
    | Design | Simple and Minimial. Package based design. Shorter syntax and fewer language features. | Object oriented language with rich features. Lot of boilerplate code. | Object oriented language with massive standard library. Easy to read and write for beginners. | Object oriented language with rich features. 
    | Concurrency | Uses Goroutines and channels  | Uses threads | Uses threads | Uses threads
    | Speed | Incredible  | Good  | Slow | Fast
    | Memory Footprint | Low | High | High | Low 
    | Garbage Collection | Yes | Yes | Yes | No
    | Static Typing | Yes | Yes | No | Yes
    | Cross Platform | Yes | Yes | Yes | Yes
    | Object Oriented | No | Yes | Yes | Yes
    | Virtual Machine | No | Yes | Yes | No

- ## **Use Cases of Go**
    | | |
    |--------|--------|
    | Web Development | Go is used to make fast and scalable web application backends. |
    | Cloud & Network Services | With a strong ecosystem of tools and APIs for major cloud platforms , it is easy to build services with Go |
    | CLIs | Go is used to create fast and elegant CLIs |
    | DevOps & SRE | Because of fast build times and lean syntax, Go is built to support DevOps and SRE. |
        
- ## **Popular Companies that Use Go**
    - Google 
    - Uber
    - Twitch
    - Dropbox
    - Docker
    - Soundcloud

# Tutorial

## **Day 1**
- ### **Hello World**
    - To print Hello World in Go, we use the fmt package.
    ```go
    package main
    import "fmt"
    func main() {
        fmt.Println("Hello, World!")
    }
    ```
    

    


        
### References
- [Concurrency in Go](https://www.golang-book.com/books/intro/10)
- [dotGo 2015 - Rob Pike ](https://www.youtube.com/watch?v=rFejpH_tAHM&list=PLEcwzBXTPUE9V1o8mZdC9tNnRZaTgI-1P)
- [Gopherfest 2015 - Rob Pike ](https://www.youtube.com/watch?v=PAAkCSZUG1c&list=PLEcwzBXTPUE9V1o8mZdC9tNnRZaTgI-1P&index=4)
- [Package Oriented Design - William Kennedy](https://www.youtube.com/watch?v=spKM5CyBwJA&list=PLEcwzBXTPUE9V1o8mZdC9tNnRZaTgI-1P&index=9)
- [Why Golang compiles so fast? ](https://www.youtube.com/watch?v=BYsspvUqCbM)
- [Golang vs Java](https://www.matellio.com/blog/golang-vs-java/#:~:text=Go%20excels%20in%20terms%20of,may%20require%20more%20memory%20resources)
- [Golang vs Python](https://www.tftus.com/blog/python-vs-golang-which-language-to-choose#:~:text=Python%20is%20acclaimed%20for%20its,simplicity%2C%20performance%2C%20and%20extensibility)
- [Golang vs C++](https://codedamn.com/news/developer-tips/golang-vs-c-which-is-best-for-you)
- [Use cases of Golang](https://go.dev/solutions/use-cases)



