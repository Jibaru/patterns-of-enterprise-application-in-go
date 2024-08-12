# Identity Map

An Identity Map keeps a record of all objects that have been read from the database in a single business transaction. Whenever you want an object, you check the Identity Map first to see if you already have it.

## Example description

The example shows the implementation of an identity map for store the registered drones in the database. A service use this identtiy map to store the registered drones and check whether a drone exists in the identity map before calling it from the database when find is called.

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
