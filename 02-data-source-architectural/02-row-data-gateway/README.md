# Row Data Gateway

A Row Data Gateway gives you objects that look exactly like the record in your record structure but can be accessed with the regular mechanisms of your programming language. All details of data source access are hidden behind this interface.

## Example description

This example shows an implementation of the pattern to manage individual employees in a database. Each employee is represented by an instance of the EmployeeGateway structure, which contains methods to insert, update, delete, and retrieve the associated employee data.

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
