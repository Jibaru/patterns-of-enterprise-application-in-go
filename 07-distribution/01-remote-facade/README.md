# Remote Facade

Provides a coarse-grained facade on fine-grained objects to improve efficiency over a network.

## Example description

This example demonstrates how access to remote address using fine-grained calls by the client, by inside the model is using coarsed-grained calls.

## How to run?

1. Run (make sure you have 8080 port available, if not, change it in the main function in /server):

2. In one terminal execute:

```
go run ./server/cmd/app/main.go
```

3. In another terminal execute two or three times:

```
go run ./client/cmd/app/main.go
```
