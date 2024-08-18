# Inheritance Mappers

A structure to organize database mappers that handle inheritance hierarchies.

## Example description

In the example, we created a base Player, a Footballer and Cricketer that inherits from Player, and Bowler that inherits from Cricketer. Each player type has its own table in the database, and the corresponding mapper class is responsible for loading and saving data to those tables. The PlayerMapper is the base mapper, while FootballerMapper, CricketerMapper, and BowlerMapper extend its functionality.

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
