# Separated Interface

Defines an interface in a separate package from its implementation.

## Example description

In this example, the Separated Interface pattern is demonstrated by defining the Storage interface in the storage package. The SQLiteStorage struct in the sqlite package and the MemoryStorage struct in the memory package implement this interface, allowing data to be saved and loaded either from an SQLite database or from an in-memory map, respectively. This separation ensures that different storage implementations can be used interchangeably, promoting flexibility and maintainability.

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
