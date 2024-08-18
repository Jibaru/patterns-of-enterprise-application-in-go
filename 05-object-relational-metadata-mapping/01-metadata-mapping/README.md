# Metadata Mapping

Holds details of object-relational mapping in metadata.

## Example description

In the example, we have customers and products, and for each one, json and tag mappers that can save each of theses structs into database tables using the defined metadata to associates table and columns with struct name and fields.

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
