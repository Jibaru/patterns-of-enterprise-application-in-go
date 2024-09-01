# Plugin

Links implementations during configuration rather than compilation.

## Example description

This example demonstrates the use of Plugin pattern by defining a `Plugin` interface that allows extending functionalities via interchangeable components. The system includes a `Processor` that manages registered plugins and processes data using them. Each plugin, like `JSONPlugin` and `XMLPlugin`, implements the `Plugin` interface to format data in different ways (e.g., JSON or XML). This pattern enables easy addition of new plugins without altering the core processor, promoting modularity and extensibility.

## How to run?

1. Run:

```
go run ./cmd/app/main.go
```
