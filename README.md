# Golang for Beginners

This repository contains contains simple examples of basic Golang concepts. There examples are provided by KodeKloud's
[Golang for Beginners](https://learn.kodekloud.com/courses/golang) course.

## Repository Overview

| Package Name | Decription                                                        |
| ------------ | ----------------------------------------------------------------- |
| hello-world  | [Prints `Hello world` to standard output](./hello-world/main.go). |

### Functions

Below is a list of packages that provide examples of functions.

| Package Name      | Description                                                                                                         |
| ----------------- | ------------------------------------------------------------------------------------------------------------------- |
| _defer-statement_ | [Demonstrates `defer` statement](./functions/defer-statement/main.go).                                              |
| _higher-order_    | [Calculates basic properties of a circle with example of higher order functions](./functions/higher-order/main.go). |

### Pointers

Below is a list of packages that provide examples of pointers.

| Package Name            | Description                                                                                                          |
| ----------------------- | -------------------------------------------------------------------------------------------------------------------- |
| _address-and-deference_ | [Address `(&)` and dereference operators `(*)`](./pointers/address-and-dereference/main.go).                         |
| _declare-pointer_       | [Declaring a pointer](./pointers/declare-pointer/main.go).                                                           |
| _dereference-pointer_   | [Dereference pointer](./pointers/dereference-pointer/main.go).                                                       |
| _init-pointer_          | [Initialising pointer using different methods](./pointers/init-pointer/main.go).                                     |
| _pass-by-reference_     | [Pass by reference for different types including `string`, `map` and `slice`](./pointers/pass-by-reference/main.go). |
| _pass-by-value_         | [Pass by value](./pointers/pass-by-value/main.go).                                                                   |

### Structs

A `struct` is a collection of fields. Below is a list of packages that provide example of structs:

| Package Name          | Description                                                                                     |
| --------------------- | ----------------------------------------------------------------------------------------------- |
| _accessing-fields_    | [Accessing fields within a struct](./structs/accessing-fields/main.go).                         |
| _comparing-structs_   | [Comparing structs](./structs/comparing-structs/main.go).                                       |
| _declare-struct_      | [Declare a `struct`](./structs/declare-struct/main.go).                                         |
| _init-struct_         | [Initialising a `struct` using different methods](./structs/init-struct/main.go).               |
| _pass-struct-to-func_ | [Passing `struct` to a function by reference and value](./structs/pass-struct-to-func/main.go). |

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
