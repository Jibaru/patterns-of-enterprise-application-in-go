# Service Stub

Removes dependence upon problematic services during testing.

## Example description

This example demonstrates the Service Stub pattern by using a stub (`UserServiceStub`) to simulate external service behavior for testing purposes. The `UserServiceStub` implements the `domain.UserService` interface, providing predefined responses to simulate retrieving user data. The `UserProcessor` relies on this interface, allowing it to work with both the real service and the stub. A test (`user_processor_test.go`) uses the stub to verify that `UserProcessor` correctly handles both existing and non-existing user scenarios without depending on the actual service implementation.

## How to run?

1. Run:

```
go run ./cmd/app/main.go
```

2. Then to see the stub in action, run:

```
go test ./...
```
