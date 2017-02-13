package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar/direction"
	"github.com/minodisk/sqlabble/internal/grammar/operator"
)

type Node interface {
	generator() generator.Generator
}

type expressionNode interface {
	Node
	expression() generator.Expression
}

type clauseNode interface {
	Node
	clauseGenerator() generator.Generator
	previous() clauseNode
}

type columnNode interface {
	expressionNode
	columnName() string
}

type tableNode interface {
	expressionNode
	previous() tableNode
	Join(tableNode) tableNode
	InnerJoin(tableNode) tableNode
	LeftJoin(tableNode) tableNode
	RightJoin(tableNode) tableNode
}

type operationNode interface {
	Node
	operator() operator.Operator
}

type comparisonOperationNode interface {
	operationNode
}

type logicalOperationNode interface {
	operationNode
	operations() []operationNode
}

type orderNode interface {
	Node
	direction() direction.Direction
}

type valuesNode interface {
	expressionNode
	clause() clauseNode
	previous() valuesNode
}
