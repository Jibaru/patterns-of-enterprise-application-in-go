# Model View Controller

Splits user interface interaction into three distinct roles.

## Example description

In the example, we have a todo-list web app. We have the list, create and edit views. Each view allows you to list, create, update or delete a todo using a controller. The controller calls model functions to update the state of the model and provide a way to show this state to the views.

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
