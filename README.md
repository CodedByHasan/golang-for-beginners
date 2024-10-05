# Golang for Beginners

This repository contains contains simple examples of basic Golang concepts. There examples are obtained from KodeKloud's
[Golang for Beginners](https://learn.kodekloud.com/courses/golang) course.

## Repository Overview

Below is an outline of the repositories, including all packages and a brief description outlining their functionality.

| Package Name  | Decription                                                        |
|---------------|-------------------------------------------------------------------|
| _hello-world_ | [Prints `Hello world` to standard output](./hello-world/main.go). |

### Interfaces

An [interface](https://go.dev/tour/methods/9) specifies a method set and is a way useful way to introduce modularity.
It is like a blueprint for a method set.

Below is a list packages that provide examples of interfaces:

| Package Name | Description                                                                      |
|--------------|----------------------------------------------------------------------------------|
| _interfaces_ | [Two different structs that implement a common interface](./interfaces/main.go). |

### Functions

Below is a list of packages that provide examples of functions:

| Package Name      | Description                                                                                                         |
|-------------------|---------------------------------------------------------------------------------------------------------------------|
| _anonymous_       | [Function that does not contain a name](./functions/anonymous/main.go).                                             |
| _defer-statement_ | [Demonstrates `defer` statement](./functions/defer-statement/main.go).                                              |
| _higher-order_    | [Calculates basic properties of a circle with example of higher order functions](./functions/higher-order/main.go). |
| _recursive_       | [Calculates n! using recursion](./functions/recursive/main.go).                                                     |
| _variadic_        | [Functions that take 0 or more inputs of the same type](./functions/variadic/main.go).                              |
| _examples_        | [Contains example functions](./functions/examples/)                                                                 |

### Pointers

A [pointer](https://go.dev/tour/moretypes/1) holds the memory address of a value.

Below is a list of packages that provide examples of pointers.

| Package Name            | Description                                                                                                          |
|-------------------------|----------------------------------------------------------------------------------------------------------------------|
| _address-and-deference_ | [Address `(&)` and dereference operators `(*)`](./pointers/address-and-dereference/main.go).                         |
| _declare-pointer_       | [Declaring a pointer](./pointers/declare-pointer/main.go).                                                           |
| _dereference-pointer_   | [Dereference pointer](./pointers/dereference-pointer/main.go).                                                       |
| _init-pointer_          | [Initialising pointer using different methods](./pointers/init-pointer/main.go).                                     |
| _pass-by-reference_     | [Pass by reference for different types including `string`, `map` and `slice`](./pointers/pass-by-reference/main.go). |
| _pass-by-value_         | [Pass by value](./pointers/pass-by-value/main.go).                                                                   |

### Structs

A `struct` is a collection of fields.

Below is a list of packages that provide example of structs:

| Package Name          | Description                                                                                     |
|-----------------------|-------------------------------------------------------------------------------------------------|
| _accessing-fields_    | [Accessing fields within a struct](./structs/accessing-fields/main.go).                         |
| _comparing-structs_   | [Comparing structs](./structs/comparing-structs/main.go).                                       |
| _declare-struct_      | [Declare a `struct`](./structs/declare-struct/main.go).                                         |
| _init-struct_         | [Initialising a `struct` using different methods](./structs/init-struct/main.go).               |
| _pass-struct-to-func_ | [Passing `struct` to a function by reference and value](./structs/pass-struct-to-func/main.go). |

### Methods

A [method](https://go.dev/tour/methods/1) is a function with a special receiver argument. The receiver appears in its
own argument list between the `func` keyword and the method name.

Below is a list packages that provide examples of methods:

| Package Name                                  | Description                                                                      |
|-----------------------------------------------|----------------------------------------------------------------------------------|
| [_methods_](./methods/intro/main.go)          | Simple method.                                                                   |
| [_method-set_](./methods/method-sets/main.go) | Method set used to interface with a [`struct`](https://go.dev/tour/moretypes/2). |

### Go routines

A [`goroutine`](https://go.dev/tour/concurrency/1) is a lightweight thread managed by the Go runtime.

| Package Name                                                         | Description                                                       |
|----------------------------------------------------------------------|-------------------------------------------------------------------|
| [_anonymous-go-routine_](./go-routines/anonymous-go-routine/main.go) | Anonymous `go` routine`.                                          |
| [_main-go-routine_](./go-routines/main-go-routine/main.go)           | The `main` go routine`.                                           |
| [_simple-example_](./go-routines/simple-example/main.go)             | Simple example demonstrating a `go` routine.                      |
| [ _wait-groups_](./go-routines/wait-groups/main.go)                  | An example of a [`WaitGroup`](https://pkg.go.dev/sync#WaitGroup). |

### Channels

[Channels](https://go.dev/tour/concurrency/2) are a typed conduit through which you can send and recieve values with the
channel operater `<-`.

| Package Name                                          | Description                                                                          |
|-------------------------------------------------------|--------------------------------------------------------------------------------------|
| [_read-and-write_](./channels/read-and-write/main.go) | Send and receive messages on unbuffered channels.                                    |
| [_buffered_](./channels/buffered/main.go)             | Send and receive messages on[ buffered channels](https://go.dev/tour/concurrency/3). |

## Running Locally

> Assuming you are in . directory.

To run the different packages locally:

```go
go run /path/to/main.go
```

For example, to run `hello-world`:

```go
go run hello-world/main.go
```
