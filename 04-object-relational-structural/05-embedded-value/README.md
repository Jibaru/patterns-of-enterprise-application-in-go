# Embedded Value

An Embedded Value maps the values of an object to fields in the record of the object's owner.

## Example description

This example demonstrates how to create an Employment entity with its embedded values in the employments table and how to retrieve it.

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
