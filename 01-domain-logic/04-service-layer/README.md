# Service Layer

Defines an application's boundary and its set of available operations from the perspective of interfacing client layers. It encapsulates the application's business logic, controlling transactions and coordinating responses in the implementation of its operations.

## Example description

This example implements the Service Layer pattern to manage users and authentication. The business logic is organized in layers: a service layer (UserService) that handles high-level operations such as registering and authenticating users, and a repository layer (UserRepository) that interacts directly with the database to perform CRUD operations on the data.

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
