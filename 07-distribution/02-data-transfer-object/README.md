# Data Transfer Object

An object that carries data between processes in order to reduce the number of method calls.

## Example description

In this example, the Data Transfer Object (DTO) pattern is used to decouple the internal representation of data from the way it is exposed to other layers of the application, such as the presentation layer. The DTO (UserDTO) serves as a simple container to transfer data between the database layer and the service layer.

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
