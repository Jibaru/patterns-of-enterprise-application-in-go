# Query Object

Mediates between the domain and data mapping layers using a collection-like interface for accessing domain objects.

## Example description

In the example, we have a single model User. A User can be stored in a UserRepository such as a collection of users.

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
