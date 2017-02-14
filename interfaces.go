package sqlabble

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/operator"
)

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
