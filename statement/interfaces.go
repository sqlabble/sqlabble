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
	nodeize() (token.Tokenizer, []interface{})
}

type Clause interface {
	Statement
	self() (token.Tokenizer, []interface{})
	previous() Clause
}

type TableOrAlias interface {
	Statement
	isTableOrAlias() bool
}

type TableOrAliasOrJoiner interface {
	Statement
	isTableOrAliasOrJoiner() bool
}

type ColumnOrColumnAsOrSubquery interface {
	Statement
	isColumnOrColumnAsOrSubquery() bool
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
	Join(TableOrAlias) Join
	InnerJoin(TableOrAlias) Join
	LeftJoin(TableOrAlias) Join
	RightJoin(TableOrAlias) Join
	previous() Joiner
}

type ComparisonOrLogicalOperation interface {
	Statement
	operator() operator.Operator
}
