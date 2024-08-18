# Concrete Table Inheritance

Represents an inheritance hierarchy of classes with one table per concrete class in the hierarchy.

## Example description

In the example, we created a base Player, a Footballer and Cricketer that inherits from Player, and Bowler that inherits from Cricketer. We use a table for each concrete struct to store them with all of their fields: Footballer, Cricketer and Bowler.

## How to run?

1. Enable CGO_ENABLED flag:

```
go env -w CGO_ENABLED=1
```

2. Install GCC:

   - Windows: https://jmeubank.github.io/tdm-gcc/

3. Run:

```
go run ./cmd/app/main.go
```
