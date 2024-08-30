# Application Controller

A centralized point for handling screen navigation and the flow of an application.

## Example description

This example demonstrates the Application Controller pattern by centralizing the flow logic in a single ApplicationController that determines which domain command to execute and which view to render based on the incoming request. The InputController handles HTTP requests and delegates decisions to the ApplicationController, which simplifies managing complex workflows and reduces code duplication. Domain commands encapsulate business logic, while views are responsible for rendering the final HTML output. This separation of concerns makes the application more maintainable and flexible.

## How to run?

1. Run (make sure you have 8080 port available, if not, change it in the main function):

```
go run ./cmd/app/main.go
```
