# Table Module

This pattern organizes business logic into modules that operate on database tables. Each module acts as a service that provides operations related to a specific table.

## Example description

Modules are created for each table (ContractModule, ProductModule, RevenueRecognitionModule), encapsulating the business logic related to each table.

The ContractModule module handles the business logic to recognize revenue based on product type. For "Software" type products, revenue is recognized in three installments; for "Service" type products, revenue is recognized immediately. The RevenueRecognitionModule module manages the insertion and retrieval of revenue recognitions.

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
