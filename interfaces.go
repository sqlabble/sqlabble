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
	nodeMine() generator.Node
	previous() clause
}

type columnOrColumnAs interface {
	expressor
	columnName() string
}

type tableOrTableAs interface {
	expressor
	Join(tableOrTableAs) tableOrTableAs
	InnerJoin(tableOrTableAs) tableOrTableAs
	LeftJoin(tableOrTableAs) tableOrTableAs
	RightJoin(tableOrTableAs) tableOrTableAs
	previous() tableOrTableAs
}

type comparisonOrLogicalOperation interface {
	Statement
	operator() operator.Operator
}

// type comparisonOperation interface {
// 	comparisonOrLogicalOperation
// }
//
// type logicalOperation interface {
// 	comparisonOrLogicalOperation
// 	operations() []comparisonOrLogicalOperation
// }

type vals interface {
	expressor
	clause() clause
	previous() vals
}
