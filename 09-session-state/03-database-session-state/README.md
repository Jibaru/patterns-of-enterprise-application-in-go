# Database Session State

Stores session data as committed data in the database.

## Example description

In this example, the session data is stored on the server side in a database rather than being sent to the client, with only a session identifier being stored in a client-side cookie.

## How to run?

1. Run (make sure you have 8080 port available, if not, change it in the main function):

```
go run ./cmd/app/main.go
```
