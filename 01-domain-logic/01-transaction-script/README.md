# Transaction Script

The Transaction Script pattern is an approach where business logic is organized into scripts or procedures that handle the operations necessary for a particular transaction. It is a simple and straightforward pattern, suitable for applications that do not have complex business logic.

## Example description

Simulates a bank account management system where the Transaction Script pattern is used to transfer money between accounts. The code will connect to a SQLite database to perform debit and credit operations.

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
