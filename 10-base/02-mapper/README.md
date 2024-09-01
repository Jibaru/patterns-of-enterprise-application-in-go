# Mapper

An object that sets up a communication between two independent objects.

## Example description

This example demonstrates the Mapper pattern by connecting independent subsystems (Customer, Lease, Asset) with a Pricing through the PricingMapper class. This setup allows each subsystem to operate independently while still enabling them to work together to calculate a final pricing.

## How to run?

1. Run:

```
go run ./cmd/app/main.go
```
