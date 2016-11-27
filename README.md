# golang

maintains *MUST READ* articles (I'v read)

## Quick Tutorial

- [A Taste of Go](https://talks.golang.org/2014/taste.slide#40)
- [Golang Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet)
- [Go for Java Programmers](https://talks.golang.org/2015/go-for-java-programmers.slide)
- [Golang Naming Convention](https://talks.golang.org/2014/names.slide)

## Language Design 

> When reading code, it should be clear what the program will do.  
> When writing code, it should be clear how to make the program do what you want.
- [Rob Pike: Simplicity is Complicated](https://www.youtube.com/watch?v=rFejpH_tAHM)

> Leading edge language features don't usually address what you really want.  
> Golang is designed for large applications, large teams and large dependencies.
- **Rob Pike: Go at Google** [Article](https://talks.golang.org/2012/splash.article), [Video](https://www.infoq.com/presentations/Go-Google)

- [Golang FAQ: Design](https://golang.org/doc/faq#Design)
- [Golang FAQ: Types](https://golang.org/doc/faq#types)
- [Campoy: Functional Go?](https://www.youtube.com/watch?v=ouyHp2nJl0I)

## Tools

- [Go Tooling in Action](https://www.youtube.com/watch?v=uBjoTxosSys)
- **Debugging Go programs with Delve** [Article](https://blog.gopheracademy.com/advent-2015/debugging-with-delve/), [Video](https://www.youtube.com/watch?v=InG72scKPd4)

## Idiomatic Go

- **Rob Pike: Go Proverb** [Page](https://go-proverbs.github.io/), [Video](https://www.youtube.com/watch?v=PAAkCSZUG1c)
- [Campoy: Understanding nil](https://www.youtube.com/watch?v=ynoY2xz-F8s)
- **Twelve Go Best Practices**: [Slide](https://talks.golang.org/2013/bestpractices.slide#1), [Video](https://www.youtube.com/watch?v=8D3Vmm1BGoY)
- **Go best practices, six years in:** [Article](https://peter.bourgon.org/go-best-practices-2016/), [Video](https://www.infoq.com/presentations/go-patterns)

## Concurrency

> *Concurrency* is about **structure**, while *Paralleism* is about **execution**

- [Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)
- **Rob Pike: Concurrency is not Parallelism** [Slide](https://talks.golang.org/2012/waza.slide), [Video](https://www.youtube.com/watch?v=B9lP-E4J_lc)
- [Curious Channels](https://dave.cheney.net/2013/04/30/curious-channels)

## Error Handling

> Error values in Go aren’t special, they are just values like any other, and so you have the entire language at your disposal.

- [Go Blog: Defer, Panic, and Recover](https://blog.golang.org/defer-panic-and-recover)
- [Go Blog: Error Handling and Go](https://blog.golang.org/error-handling-and-go)
- [Go Blog: Errors are Values](https://blog.golang.org/errors-are-values)
- **Dave Cheney: Don't just check errors, handle them gracefully** [Article](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully), [Video](https://www.youtube.com/watch?v=lsBF58Q-DnY)

> Go solves the exception problem by not having exceptions.  
> ...
> The decision to not include exceptions in Go is an example of its simplicity and orthogonality. Using multiple return values and a simple convention, Go solves the problem of letting programmers know when things have gone wrong and reserves panic for the truly exceptional.

- [Dave Cheney: Why Go gets exceptions right](https://dave.cheney.net/2012/01/18/why-go-gets-exceptions-right)
- [Dave Cheney: Inspecting errors](https://dave.cheney.net/2014/12/24/inspecting-errors)
- [Dave Cheney: Error handling vs. exceptions redux](https://dave.cheney.net/2014/11/04/error-handling-vs-exceptions-redux)
- [Dave Cheney: Errors and Exceptions, redux](https://dave.cheney.net/2015/01/26/errors-and-exceptions-redux)

> Knowing the difference between which errors to ignore and which to check is why we’re paid as professionals.

- [Dave Cheney: Constant errors](https://dave.cheney.net/2016/04/07/constant-errors)
- [Dave Cheney: Stack traces and the errors package](https://dave.cheney.net/2016/06/12/stack-traces-and-the-errors-package)

## Interface

> If C++ and Java are about type hierarchies and the taxonomy of types, **Go is about composition**

> Q. Hey gophers, what was the best/worst moment of your experienes lenaring golang?  
> A. **The worst was interface, but the best was also interface**

- [Stackoverflow: What's the mearning of interface{} ?](http://stackoverflow.com/questions/23148812/go-whats-the-meaning-of-interface)
- [How to use interfaces in Go](http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go)

## Struct

- [Dave Cheney: Struct composition with Go](https://dave.cheney.net/2015/05/22/struct-composition-with-go)
- [Dave Cheney: The empty struct](https://dave.cheney.net/2014/03/25/the-empty-struct)

## Pointer

- [Dave Cheney: Pointers in Go](https://dave.cheney.net/2014/03/17/pointers-in-go)
- [Things I Wish Someone Had Told Me About Go](http://openmymind.net/Things-I-Wish-Someone-Had-Told-Me-About-Go/)
- [Dave Cheney: Go has both make and new functions, what gives?](https://dave.cheney.net/2014/08/17/go-has-both-make-and-new-functions-what-gives)
- [Dave Cheney: Should methods be declared on T or *T](https://dave.cheney.net/2016/03/19/should-methods-be-declared-on-t-or-t)

## Map, Slice

- [Go Blog: Map in Action](https://blog.golang.org/go-maps-in-action)
- [Go Blog: Slices Usage and Internals](https://blog.golang.org/go-slices-usage-and-internals)

## Logging

- [The Hunt for a Logger Interface](http://go-talks.appspot.com/github.com/ChrisHines/talks/structured-logging/structured-logging.slide#1)
- [Logging v. instrumentation](https://peter.bourgon.org/blog/2016/02/07/logging-v-instrumentation.html)
- [Dave Cheney: Let’s talk about logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging)

## Encoding, JSON

- [JSON, interface, and go generate](https://www.youtube.com/watch?v=YgnD27GFcyA)



