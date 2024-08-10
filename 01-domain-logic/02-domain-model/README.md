# Domain Model

The Domain Model pattern is an approach in which business logic is organized around a rich model of objects that represent key domain concepts. These objects encapsulate not only data, but also related behavior, and collaborate with each other to perform business operations.

## Example description

Simulates a purchase order management system, where the Domain Model pattern is used to model entities such as Order, Customer, and Product, and manage the creation and processing of orders.

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
