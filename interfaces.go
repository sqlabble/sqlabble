package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/operator"
)

// Statement is the interface of the component
// which is the minimum unit constituting SQL.
// All types that implement this interface
// can be built as SQL.
type Statement interface {
	node() generator.Node
}

type expressor interface {
	Statement
	expression() generator.Expression
}

type clause interface {
	Statement
	myNode() generator.Node
	previous() clause
}

type columnOrColumnAs interface {
	expressor
	columnName() string
}

type columnOrSubquery interface {
	Statement
	isColumnOrSubquery() bool
}

type joiner interface {
	expressor
	Join(joiner) joiner
	InnerJoin(joiner) joiner
	LeftJoin(joiner) joiner
	RightJoin(joiner) joiner
	previous() joiner
}

type comparisonOrLogicalOperation interface {
	Statement
	operator() operator.Operator
}

type vals interface {
	expressor
	clause() clause
	previous() vals
}
