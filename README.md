# Golang for Beginners

This repository contains contains simple examples of basic Golang concepts. There examples are provided by KodeKloud's
[Golang for Beginners](https://learn.kodekloud.com/courses/golang) course.

## Repository Overview

| Package Name | Decription                               |
| ------------ | ---------------------------------------- |
| hello-world  | Prints `Hello world` to standard output. |

### Functions

Below is a list of packages that provide examples of functions.

| Package Name      | Description                                                                       |
| ----------------- | --------------------------------------------------------------------------------- |
| _defer-statement_ | Demonstrates `defer` statement.                                                   |
| _higher-order_    | Program calculates basic properties of a circle. Contains higher order functions. |

### Pointers

Below is a list of packages that provide examples of pointers.

| Package Name            | Description                                                                  |
| ----------------------- | ---------------------------------------------------------------------------- |
| _address-and-deference_ | Address `(&)` and dereference operators `(*)`.                               |
| _declare-pointer_       | Declaring a pointer.                                                         |
| _dereference-pointer_   | Dereference pointer.                                                         |
| _init-pointer_          | Initialising pointer using different methods.                                |
| _pass-by-reference_     | Pass by reference for different types including `string`, `map` and `slice`. |
| _pass-by-value_         | Pass by value.                                                               |

### Structs

Below is a list of packages that provide example of structs:

| Package Name       | Description                                                                       |
| ------------------ | --------------------------------------------------------------------------------- |
| _accessing-fields_ | [Accessing fields within a struct](./structs/accessing-fields/main.go).           |
| _declare-struct_   | [Declare a `struct`](./structs/declare-struct/main.go).                           |
| _init-struct_      | [Initialising a `struct` using different methods](./structs/init-struct/main.go). |

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
