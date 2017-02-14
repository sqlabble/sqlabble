package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/operator"
)

type Node interface {
	generator() generator.Generator
}

type expression interface {
	Node
	expression() generator.Expression
}

type clause interface {
	Node
	clauseGenerator() generator.Generator
	previous() clause
}

type columnOrColumnAs interface {
	expression
	columnName() string
}

type tableOrTableAs interface {
	expression
	Join(tableOrTableAs) tableOrTableAs
	InnerJoin(tableOrTableAs) tableOrTableAs
	LeftJoin(tableOrTableAs) tableOrTableAs
	RightJoin(tableOrTableAs) tableOrTableAs
	previous() tableOrTableAs
}

type comparisonOrLogicalOperation interface {
	Node
	operator() operator.Operator
}

type comparisonOperation interface {
	comparisonOrLogicalOperation
}

type logicalOperation interface {
	comparisonOrLogicalOperation
	operations() []comparisonOrLogicalOperation
}

type vals interface {
	expression
	clause() clause
	previous() vals
}
