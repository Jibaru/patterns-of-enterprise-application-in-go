# Coarse-Grained Lock

Locks a set of related objects with a single lock.

## Example description

In this example, the Coarse-Grained Lock pattern is implemented by applying a single sync.Mutex lock to the entire CustomerService class, which manages both Customer and Address entities. When a customer or an associated address is updated, the lock ensures that the entire customer record and its associated addresses are locked. This approach simplifies concurrency control by ensuring that all related objects are managed under a single lock, reducing the risk of deadlocks but potentially decreasing the system's concurrency performance.

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
