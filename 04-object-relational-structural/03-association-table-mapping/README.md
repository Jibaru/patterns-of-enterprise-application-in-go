# Association Table Mapping

Saves an association as a table with foreign keys to the tables that are linked by the association.

## Example description

This example demonstrates how we can create authors, books, associate them, and retrieve the books by a specific author.

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
