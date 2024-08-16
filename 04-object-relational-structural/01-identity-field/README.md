# Identity Field

Saves a database ID field in an object to maintain identity between an in-memory object and a database row.

## Example description

This example demonstrates how the Identity Field works by creating a new device and automatically generating its ID.

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
