# Layer Supertype

A type that acts as the supertype for all types in its layer.

## Example description

In this example, the Layer Supertype pattern is applied through the Entity struct for the entities layer, which contains common fields (ID, CreatedAt, UpdatedAt) and methods (UpdateTimestamp). Both Customer and Order structs embed Entity, inheriting these common properties and behaviors. This approach reduces redundancy and ensures that all entities within the layer have consistent behavior.

## How to run?

1. Run:

```
go run ./cmd/app/main.go
```
