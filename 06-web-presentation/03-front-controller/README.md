# Front Controller

A controller that handles all requests for a Web site.

## Example description

This example demonstrates how the Front Controller pattern can be used in Go to centralize request handling and delegate specific tasks to command objects. The pattern ensures that all incoming requests are processed through a single point, making it easier to apply common behavior and reducing redundancy across the application.

## How to run?

1. Enable CGO_ENABLED flag:

```
go env -w CGO_ENABLED=1
```

2. Install GCC:

   - Windows: https://jmeubank.github.io/tdm-gcc/

3. Run (make sure you have 8080 port available, if not, change it in the main function):

```
go run ./cmd/app/main.go
```
