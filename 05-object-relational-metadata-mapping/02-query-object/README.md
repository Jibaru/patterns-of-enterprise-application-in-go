# Query Object

An object that represents a database query. A Query Object is an interpreter, that is, a structure of objects that can form itself into a SQL query. You can create this query by referring to classes and fields rather than tables and columns. In this way those who write the queries can do so independently of the database schema and changes to the schema can be localized in a single place.

## Example description

In the example, we have a single model Product. We defined a struct Query that holds a bunch of Criterias to perform queries against the database for the Product model.

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
