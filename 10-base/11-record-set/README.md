# Record Set

An in-memory representation of tabular data.

## Example description

By using the sql.Rows type, we describe the use of record set, and how to dump this in-memory data to get all users from a database.

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
