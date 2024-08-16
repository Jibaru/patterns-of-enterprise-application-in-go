# Single Table Inheritance

Represents an inheritance hierarchy of classes as a single table that has columns for all the fields of the various classes.

## Example description

In the example, we created a base Player, a Footballer and Cricketer that inherits from Player, and Bowler that inherits from Cricketer. We simplify the database schema by using one table to store multiple related entity types, while still maintaining the distinction between different types of players.

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
