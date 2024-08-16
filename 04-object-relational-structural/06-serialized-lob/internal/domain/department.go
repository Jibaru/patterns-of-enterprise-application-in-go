package domain

// Department represents a department entity.
type Department struct {
	Name     string
	Children []*Department // Child departments (hierarchical relationship)
}
