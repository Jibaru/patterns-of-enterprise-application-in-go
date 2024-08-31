# Optimistic Offline Lock

Prevents conflicts between concurrent business transactions by detecting a conflict and rolling back the transaction.

## Example description

In this example, the Optimistic Offline Lock pattern is implemented using a version field in the Product model. The version is incremented every time a product is updated. When an update is attempted, the service checks whether the version in the database matches the version provided by the transaction. If they don't match, it means another transaction has modified the data, and a conflict is detected, causing the transaction to fail. This approach helps to avoid lost updates and ensures data consistency in scenarios where multiple transactions might attempt to update the same resource simultaneously.

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
