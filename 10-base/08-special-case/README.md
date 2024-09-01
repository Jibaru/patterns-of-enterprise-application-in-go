# Special Case

A "subclass" or implementation that provides special behavior for particular cases.

## Example description

This example demonstrates Special Case pattern by defining a domain.OrderInterface that represents typical orders and special cases like NoOrder. The repository always returns types implementing this interface. If an order is not found, the repository returns a NoOrder instance, a special case that implements the same interface as Order, allowing the rest of the system to handle it uniformly without additional error handling or null checks.

## How to run?

1. Run:

```
go run ./cmd/app/main.go
```
