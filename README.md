# Golang for Beginners

This repository contains contains simple examples of [Golang](https://go.dev/) concepts.

## Repository Overview

Below is an outline of all examples contained in this repository.

### Hello World

| Name                                   | Decription                               |
| -------------------------------------- | ---------------------------------------- |
| [_hello-world_](./hello-world/main.go) | Prints `Hello world` to standard output. |

### Interfaces

An [interface](https://go.dev/tour/methods/9) specifies a method set and is a useful way to introduce modularity.
It is like a blueprint for a method set.

Examples of interfaces can be found below:

| Name                                 | Description                                              |
| ------------------------------------ | -------------------------------------------------------- |
| [_interfaces_](./interfaces/main.go) | Two different structs that implement a common interface. |

### Functions

A [function](https://go.dev/tour/basics/4) can take zero or more arguments.

Examples of functions can be found below:

| Name                                                     | Description                                                                     |
| -------------------------------------------------------- | ------------------------------------------------------------------------------- |
| [_anonymous_](./functions/anonymous/main.go)             | Function that does not contain a name.                                          |
| [_defer-statement_](./functions/defer-statement/main.go) | Demonstrates `defer` statement.                                                 |
| [_higher-order_](./functions/higher-order/main.go)       | Calculates basic properties of a circle with example of higher order functions. |
| [_recursive_](./functions/recursive/main.go)             | Calculates `n!` using recursion.                                                |
| [_variadic_](./functions/variadic/main.go)               | Functions that take 0 or more inputs of the same type.                          |
| [_examples_](./functions/examples/main.go)               | Contains example functions.                                                     |

### Pointers

A [pointer](https://go.dev/tour/moretypes/1) holds the memory address of a value.

Examples of pointers can be found below:

| Name                                                                  | Description                                                                  |
| --------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| [_address-and-deference_](./pointers/address-and-dereference/main.go) | Address `(&)` and dereference operators `(*)`.                               |
| [_declare-pointer_](./pointers/declare-pointer/main.go)               | Declaring a pointer.                                                         |
| [_dereference-pointer_](./pointers/dereference-pointer/main.go)       | Dereference pointer.                                                         |
| [_init-pointer_](./pointers/init-pointer/main.go)                     | Initialising pointer using different methods.                                |
| [_pass-by-reference_](./pointers/pass-by-reference/main.go)           | Pass by reference for different types including `string`, `map` and `slice`. |
| [_pass-by-value_](./pointers/pass-by-value/main.go)                   | Pass by value.                                                               |

### Structs

A [`struct`](https://go.dev/tour/moretypes/2) is a collection of fields.

Example of structs can be found below:

| Name                                                           | Description                                            |
| -------------------------------------------------------------- | ------------------------------------------------------ |
| [_accessing-fields_](./structs/accessing-fields/main.go)       | Accessing fields within a `struct`.                    |
| [_comparing-structs_](./structs/comparing-structs/main.go)     | Comparing structs.                                     |
| [_declare-struct_](./structs/declare-struct/main.go)           | Declare a `struct`.                                    |
| [_init-struct_](./structs/init-struct/main.go)                 | Initialising a `struct` using different methods.       |
| [_pass-struct-to-func_](./structs/pass-struct-to-func/main.go) | Passing `struct` to a function by reference and value. |

### Methods

A [method](https://go.dev/tour/methods/1) is a function with a special receiver argument. The receiver appears in its
own argument list between the `func` keyword and the method name.

Examples of methods can be found below:

| Name                                          | Description                                   |
| --------------------------------------------- | --------------------------------------------- |
| [_methods_](./methods/intro/main.go)          | Simple method.                                |
| [_method-set_](./methods/method-sets/main.go) | Method set used to interface with a `struct`. |

### Goroutines

A [`goroutine`](https://go.dev/tour/concurrency/1) is a lightweight thread managed by the Go runtime.

Examples of goroutines can be found below:

| Name                                                                            | Description                                                       |
| ------------------------------------------------------------------------------- | ----------------------------------------------------------------- |
| [_anonymous-go-routine_](./go-routines/anonymous-go-routine/main.go)            | Anonymous goroutine.                                              |
| [_main-go-routine_](./go-routines/main-go-routine/main.go)                      | The `main` goroutine.                                             |
| [_simple-example_](./go-routines/simple-example/main.go)                        | Simple example demonstrating a `go` routine.                      |
| [ _wait-groups_](./go-routines/wait-groups/main.go)                             | An example of a [`WaitGroup`](https://pkg.go.dev/sync#WaitGroup). |
| [ _print-letters-and-numbers_](./go-routines/print-letters-and-numbers/main.go) | Using goroutines to prints letter and numbers to the screen.      |

### Channels

[Channels](https://go.dev/tour/concurrency/2) are a typed conduit through which you can send and receive values with the
channel operater `<-`.

Examples of channels can be found below:

| Name                                                        | Description                                                                                                   |
| ----------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| [_read-and-write_](./channels/read-and-write/main.go)       | Send and receive messages on unbuffered channels.                                                             |
| [_buffered_](./channels/buffered/main.go)                   | Send and receive messages on [buffered channels](https://go.dev/tour/concurrency/3).                          |
| [_closing-a-channel_](./channels/closing-a-channel/main.go) | Example of closing a buffered channel.                                                                        |
| [_for-range_](./channels/for-range/main.go)                 | Using a `for-range` to receive data from a buffered channel.                                                  |
| [_go-routine-leak_](./channels/go-routine-leak/main.go)     | A goroutine leak cause by a goroutine that remains active indefinitely.                                       |
| [_select-statement_](./channels/select-statement/main.go)   | Using a [`select`](https://go.dev/tour/concurrency/5) statement to wait on multiple communication operations. |
| [_closure-in-a-loop_](./channels/closure-in-a-loop/main.go) | A goroutine [closure](https://go.dev/tour/moretypes/25) executed in a `for` loop                              |
| [_time-out-code_](./channels/timeout-code/main.go)          | Adding a [timeout](https://go.dev/wiki/Timeouts) in a `select` statement when waiting on receiver/sender.     |
| [_rx-and-tx-data-1_](./channels/rx-and-tx-data-1/main.go)   | Sending and received data using channels.                                                                     |
| [_rx-and-tx-data-2_](./channels/rx-and-tx-data-2/main.go)   | Another example of sending and receiving data using channels.                                                 |

## Running Locally

> Assuming you are in . directory.

To run the different examples locally:

```go
go run /path/to/main.go
```

For example, to run `hello-world`:

```go
go run hello-world/main.go
```

## Acknowledgements

There examples are obtained from KodeKloud's [Golang for Beginners](https://learn.kodekloud.com/courses/golang)
and [Advanced Golang](https://learn.kodekloud.com/courses/advanced-golang) courses.
