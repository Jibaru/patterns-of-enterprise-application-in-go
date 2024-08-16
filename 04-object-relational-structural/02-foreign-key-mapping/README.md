# Foreign Key Mapping

Maps an association between objects to a foreign key reference between tables.

## Example description

This example demonstrates how we can create an artist, add albums linked to that artist using a foreign key, and retrieve all albums associated with that artist.

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
