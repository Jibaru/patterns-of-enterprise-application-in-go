# Session State

Stores session state on the client.

## Example description

In this example, the Client Session State pattern is implemented using cookies to manage session state. When a user logs in, the session state (user ID) is stored in a cookie on the client side. This cookie is then sent with each request, allowing the server to maintain a stateless architecture. The SessionController handles the session logic, while the Session model provides methods for saving, retrieving, and clearing the session data stored in cookies. The SessionView is responsible for rendering the HTML based on the session state.

## How to run?

1. Run (make sure you have 8080 port available, if not, change it in the main function):

```
go run ./cmd/app/main.go
```
