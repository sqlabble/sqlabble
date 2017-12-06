package keyword

// Operator is reserved keyword in SQL statement.
type Operator string

// Logical operators.
const (
	And Operator = "AND"
	Or  Operator = "OR"
	Not Operator = "NOT"

	Unique Operator = "UNIQUE"

	Like   Operator = "LIKE"
	RegExp Operator = "REGEXP"

	Between Operator = "BETWEEN"

	In    Operator = "IN"
	NotIn Operator = "NOT IN"

	Is    Operator = "IS"
	IsNot Operator = "IS NOT"
)

// Scalar Comparison operators.
const (
	Eq    Operator = "="
	NotEq Operator = "!="
	Gt    Operator = ">"
	Gte   Operator = ">="
	Lt    Operator = "<"
	Lte   Operator = "<="
)

// Non-Scalar Comparison operators.
const (
	EqAll     Operator = "= ALL"
	NotEqAll  Operator = "!= ALL"
	GtAll     Operator = "> ALL"
	GteAll    Operator = ">= ALL"
	LtAll     Operator = "< ALL"
	LteAll    Operator = "<= ALL"
	EqAny     Operator = "= ANY"
	NotEqAny  Operator = "!= ANY"
	GtAny     Operator = "> ANY"
	GteAny    Operator = ">= ANY"
	LtAny     Operator = "< ANY"
	LteAny    Operator = "<= ANY"
	Exists    Operator = "EXISTS"
	NotExists Operator = "NOT EXISTS"
)

// Arithmetic operators.
const (
	Add Operator = "+"
	Sub Operator = "-"
	Mul Operator = "*"
	Div Operator = "/"
	IntegerDiv Operator = "DIV"
	Mod Operator = "%"
)

// Alias operators.
const (
	As Operator = "AS"
)

// Set operators.
const (
	Union        = "UNION"
	UnionAll     = "UNION ALL"
	Intersect    = "INTERSECT"
	IntersectAll = "INTERSECT ALL"
	Except       = "EXCEPT"
	ExceptAll    = "EXCEPT ALL"
)
