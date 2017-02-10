package grammar

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar/direction"
	"github.com/minodisk/sqlabble/internal/grammar/operator"
)

type Statement interface {
	Generator() generator.Generator
}

type Container interface {
	Statement
	Container() generator.Container
}

type Expression interface {
	Statement
	Expression() generator.Expression
}

type Clause interface {
	Container
	Prev() Clause
}

type Column interface {
	Expression
	ColumnName() string
}

type ColumnAlias interface {
	Column
	Alias() string
}

type Table interface {
	Expression
	Prev() Table
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

type Values interface {
	Expression
	Clause() Clause
	Prev() Values
}
