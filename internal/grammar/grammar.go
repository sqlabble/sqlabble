package grammar

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar/direction"
	"github.com/minodisk/sqlabble/internal/grammar/operator"
)

type Statement interface {
	Generator() generator.Generator
}

type Values interface {
	Statement
	Clause() Clause
	Prev() Values
	Expression() generator.Expression
}

type Clause interface {
	Statement
	Prev() Clause
	Container() generator.Container
}

type Column interface {
	Statement
	ColumnName() string
}

type Table interface {
	Statement
	Prev() Table
	Expression() generator.Expression
	Join(Table) Table
	InnerJoin(Table) Table
	LeftJoin(Table) Table
	RightJoin(Table) Table
}

type Operation interface {
	Statement
	Operator() operator.Operator
}

type ComparisonOperation interface {
	Operation
}

type LogicalOperation interface {
	Operation
	Operations() []Operation
}

type Order interface {
	Statement
	Direction() direction.Direction
}
