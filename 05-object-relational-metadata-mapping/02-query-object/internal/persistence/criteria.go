package persistence

type Op string

const (
	GreaterThanOp       Op = "gt"
	LowerThanOp         Op = "lt"
	LowerThanEqualsOp   Op = "lte"
	GreaterThanEqualsOp Op = "gte"
	EqualsOp            Op = "eq"
	NotEqualsOp         Op = "neq"
	IncludesOp          Op = "includes"
)

type Criteria struct {
	Operator Op
	Field    string
	Value    interface{}
}

type Query []Criteria
