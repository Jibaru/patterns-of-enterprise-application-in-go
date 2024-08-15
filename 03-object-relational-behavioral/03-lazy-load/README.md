# Lazy Load

A Lazy Load interrupts the loading process for the moment, leaving a marker in the object structure so that if the data is needed it can be loaded only when it is used. As many people know, if you're lazy about doing things you'll win when it turns out you don't need to do them at all.

## Example description

This example implements the Lazy Load pattern in an application that manages customers, their orders, and order lines. The Customer struct allows you to load a customer's orders lazily, that is, only when they are necessary. Likewise, orders load their order lines only when they are accessed for the first time. This optimizes performance, especially in scenarios where only part of the information is required.

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
