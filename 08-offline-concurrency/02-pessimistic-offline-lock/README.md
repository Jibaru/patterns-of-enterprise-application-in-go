# Pessimistic Offline Lock

Prevents conflicts between concurrent business transactions by allowing only one business transaction at a time to access data.

## Example description

In this example, the Pessimistic Offline Lock pattern is implemented by locking a Product when it is being accessed for updates. The locked field in the Product model is set to true when a transaction begins and set back to false when the transaction completes. If another transaction tries to lock the product while it is already locked, it fails with an error, preventing concurrent modifications. This approach ensures that only one transaction can modify a product at a time, thus avoiding conflicts but potentially reducing throughput due to the locking mechanism.

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
