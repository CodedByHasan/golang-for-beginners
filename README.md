# Golang for Beginners

This repository contains contains simple examples of basic Golang concepts. There examples are obtained from KodeKloud's
[Golang for Beginners](https://learn.kodekloud.com/courses/golang) course.

## Repository Overview

Below is an outline of the repositories, including all packages and a brief description outlining their functionality.

### Hello World

| Package Name                           | Decription                               |
|----------------------------------------|------------------------------------------|
| [_hello-world_](./hello-world/main.go) | Prints `Hello world` to standard output. |

### Interfaces

An [interface](https://go.dev/tour/methods/9) specifies a method set and is a way useful way to introduce modularity.
It is like a blueprint for a method set.

Below is a list packages that provide examples of interfaces:

| Package Name                         | Description                                              |
|--------------------------------------|----------------------------------------------------------|
| [_interfaces_](./interfaces/main.go) | Two different structs that implement a common interface. |

### Functions

A [function](https://go.dev/tour/basics/4) can take zero or more arguments.

Below is a list of packages that provide examples of functions:

| Package Name                                             | Description                                                                     |
|----------------------------------------------------------|---------------------------------------------------------------------------------|
| [_anonymous_](./functions/anonymous/main.go)             | Function that does not contain a name.                                          |
| [_defer-statement_](./functions/defer-statement/main.go) | Demonstrates `defer` statement.                                                 |
| [_higher-order_](./functions/higher-order/main.go)       | Calculates basic properties of a circle with example of higher order functions. |
| [_recursive_](./functions/recursive/main.go)             | Calculates `n!` using recursion.                                                |
| [_variadic_](./functions/variadic/main.go)               | Functions that take 0 or more inputs of the same type.                          |
| [_examples_](./functions/examples/calculate.go)          | Contains example functions.                                                     |

### Pointers

A [pointer](https://go.dev/tour/moretypes/1) holds the memory address of a value.

Below is a list of packages that provide examples of pointers.

| Package Name                                                          | Description                                                                  |
|-----------------------------------------------------------------------|------------------------------------------------------------------------------|
| [_address-and-deference_](./pointers/address-and-dereference/main.go) | Address `(&)` and dereference operators `(*)`.                               |
| [_declare-pointer_](./pointers/declare-pointer/main.go)               | Declaring a pointer.                                                         |
| [_dereference-pointer_](./pointers/dereference-pointer/main.go)       | Dereference pointer.                                                         |
| [_init-pointer_](./pointers/init-pointer/main.go)                     | Initialising pointer using different methods.                                |
| [_pass-by-reference_](./pointers/pass-by-reference/main.go)           | Pass by reference for different types including `string`, `map` and `slice`. |
| [_pass-by-value_](./pointers/pass-by-value/main.go)                   | Pass by value.                                                               |

### Structs

A [struct](https://go.dev/tour/moretypes/2) is a collection of fields.

Below is a list of packages that provide example of structs:

| Package Name                                               | Description                                                                                   |
|------------------------------------------------------------|-----------------------------------------------------------------------------------------------|
| [_accessing-fields_](./structs/accessing-fields/main.go)   | Accessing fields within a struct.                                                             |
| [_comparing-structs_](./structs/comparing-structs/main.go) | Comparing structs.                                                                            |
| _declare-struct_                                           | [Declare a struct](./structs/declare-struct/main.go).                                         |
| _init-struct_                                              | [Initialising a struct using different methods](./structs/init-struct/main.go).               |
| _pass-struct-to-func_                                      | [Passing struct to a function by reference and value](./structs/pass-struct-to-func/main.go). |

### Methods

A [method](https://go.dev/tour/methods/1) is a function with a special receiver argument. The receiver appears in its
own argument list between the `func` keyword and the method name.

Below is a list packages that provide examples of methods:

| Package Name                                  | Description                                   |
|-----------------------------------------------|-----------------------------------------------|
| [_methods_](./methods/intro/main.go)          | Simple method.                                |
| [_method-set_](./methods/method-sets/main.go) | Method set used to interface with a `struct`. |

### Go routines

A [`goroutine`](https://go.dev/tour/concurrency/1) is a lightweight thread managed by the Go runtime.

Below is a list packages that provide examples of go routines:

| Package Name                                                         | Description                                                       |
|----------------------------------------------------------------------|-------------------------------------------------------------------|
| [_anonymous-go-routine_](./go-routines/anonymous-go-routine/main.go) | Anonymous `go` routine`.                                          |
| [_main-go-routine_](./go-routines/main-go-routine/main.go)           | The `main` go routine`.                                           |
| [_simple-example_](./go-routines/simple-example/main.go)             | Simple example demonstrating a `go` routine.                      |
| [ _wait-groups_](./go-routines/wait-groups/main.go)                  | An example of a [`WaitGroup`](https://pkg.go.dev/sync#WaitGroup). |

### Channels

[Channels](https://go.dev/tour/concurrency/2) are a typed conduit through which you can send and recieve values with the
channel operater `<-`.

Below is a list packages that provide examples of channels:

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
