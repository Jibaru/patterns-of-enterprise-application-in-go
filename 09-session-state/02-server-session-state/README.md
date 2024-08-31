# Server Session State

Keeps the session state on a server system in a serialized form.

## Example description

In this example, the session data is stored on the server side rather than being sent to the client, with only a session identifier being stored in a client-side cookie.

## How to run?

1. Run (make sure you have 8080 port available, if not, change it in the main function):

```
go run ./cmd/app/main.go
```
