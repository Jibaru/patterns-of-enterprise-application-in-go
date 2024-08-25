# Page Controller

An object that handles a request for a specific page or action on a Web site.

## Example description

This example demonstrates the Page Controller pattern by creating separate controllers for different pages and actions within a simple blog application. Each controller is responsible for the logic and rendering of its respective page.

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
