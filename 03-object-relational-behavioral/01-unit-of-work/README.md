# Unit Of Work

A Unit of Work keeps track of everything you do during a business transaction that can affect the database. When you're done, it figures out everything that needs to be done to alter the database as a result of your work.

## Example description

The example shows the implementation of a unit of work for managing book transactions in the database. Likewise, 3 scripts have been added to add, modify and delete several books, which execute everything in a single transaction and not in small calls to the database.

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
