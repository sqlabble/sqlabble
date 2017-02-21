package statement

import (
	"github.com/minodisk/sqlabble/operator"
	"github.com/minodisk/sqlabble/token"
)

func Nodeize(stmt Statement) (token.Tokenizer, []interface{}) {
	return stmt.nodeize()
}

// Statement is the interface of the component
// which is the minimum unit constituting SQL.
// All types that implement this interface
// can be built as SQL.
type Statement interface {
	// node() node.Node
	nodeize() (token.Tokenizer, []interface{})
}

// type Expressor interface {
// 	Statement
// 	expression() node.Expression
// }

type Clause interface {
	Statement
	container() (token.Container, []interface{})
	previous() Clause
}

type ColumnOrColumnAs interface {
	Statement
	line() (token.Line, []interface{})
	columnName() string
}

type ColumnOrSubquery interface {
	Statement
	isColumnOrSubquery() bool
}

type ParamOrSubquery interface {
	Statement
	isParamOrSubquery() bool
}

type ParamsOrSubquery interface {
	Statement
	isParamsOrSubquery() bool
}

type Joiner interface {
	Statement
	self() (token.Tokenizer, []interface{})
	Join(Joiner) Joiner
	InnerJoin(Joiner) Joiner
	LeftJoin(Joiner) Joiner
	RightJoin(Joiner) Joiner
	previous() Joiner
}

type ComparisonOrLogicalOperation interface {
	Statement
	// lines() (token.Lines, []interface{})
	operator() operator.Operator
}

type Vals interface {
	// Expressor
	clause() Clause
	previous() Vals
	line() (token.Line, []interface{})
}
