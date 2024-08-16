# Dependent Mapping

Has one class perform the database mapping for a child class.

## Example description

This example demonstrates how to create an album with tracks and how to retrieve the album and its dependent tracks.

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
