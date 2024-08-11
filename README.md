# Golang for Beginners

This repository contains contains simple examples of basic Golang concepts. There examples are provided by KodeKloud's
[Golang for Beginners](https://learn.kodekloud.com/courses/golang) course.

## Repository Overview

| Package Name | Decription                               |
| ------------ | ---------------------------------------- |
| hello-world  | Prints `Hello world` to standard output. |

### Functions

Below is a list of packages that provide examples of functions.

| Package Name    | Description                                                                       |
| --------------- | --------------------------------------------------------------------------------- |
| defer-statement | Demonstrates `defer` statement.                                                   |
| higher-order    | Program calculates basic properties of a circle. Contains higher order functions. |

### Pointers

Below is a list of packages that provide examples of pointers.

| Package Name          | Description                                    |
| --------------------- | ---------------------------------------------- |
| address-and-deference | Address `(&)` and dereference operators `(*)`. |
| declare-pointer       | Declaring a pointer.                           |
| init-pointer          | Initialising pointer using different methods.  |

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
