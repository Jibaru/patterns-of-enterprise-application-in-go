# Implicit Lock

Allows framework or layer supertype code to acquire offline locks.

## Example description

In this example, the Implicit Lock pattern is demonstrated by using database transactions to manage concurrency without manually controlling locks. The OrderService creates and updates orders within transactions (sql.Tx), relying on the database to implicitly lock the relevant rows during the operation. This ensures that concurrent access to the same data is handled safely and efficiently, avoiding conflicts or data corruption. The simplicity of this approach makes it well-suited for scenarios where the underlying database provides robust transaction management.

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
