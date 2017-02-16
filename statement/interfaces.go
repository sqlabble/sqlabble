package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/operator"
)

// Statement is the interface of the component
// which is the minimum unit constituting SQL.
// All types that implement this interface
// can be built as SQL.
type Statement interface {
	node() generator.Node
}

type Expressor interface {
	Statement
	expression() generator.Expression
}

type Clause interface {
	Statement
	myNode() generator.Node
	previous() Clause
}

type ColumnOrColumnAs interface {
	Expressor
	columnName() string
}

type columnOrSubquery interface {
	Statement
	isColumnOrSubquery() bool
}

type Joiner interface {
	Expressor
	Join(Joiner) Joiner
	InnerJoin(Joiner) Joiner
	LeftJoin(Joiner) Joiner
	RightJoin(Joiner) Joiner
	previous() Joiner
}

type ComparisonOrLogicalOperation interface {
	Statement
	operator() operator.Operator
}

type Vals interface {
	Expressor
	clause() Clause
	previous() Vals
}
