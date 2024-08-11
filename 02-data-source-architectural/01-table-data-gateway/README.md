# Table Data Gateway

A Table Data Gateway holds all the SQL for accessing a single table or view: selects, inserts, updates, and deletes. Other code calls its methods for all interaction with the database.

## Example description

This example implements the Table Data Gateway pattern to manage the employees table. The EmployeeGateway class acts as a gateway, providing centralized methods to perform CRUD (create, read, update, delete) operations on the employees table.

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
