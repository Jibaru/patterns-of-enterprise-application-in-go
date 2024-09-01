# Value Object

A small simple object, like money or a date range, whose equality isn't based on identity.

## Example description

In this example, the Value Object pattern is demonstrated by encapsulating user data within distinct value objects (Email, PhoneNumber, Address, Name, and DateOfBirth). Each value object includes its own validation logic, ensuring the integrity of data before it’s used to construct a User model. The User entity aggregates these value objects.

## How to run?

1. Run:

```
go run ./cmd/app/main.go
```
