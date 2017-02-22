package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/tokenizer"
)

// Statement is the interface of the component
// which is the minimum unit constituting SQL.
// All types that implement this interface
// can be built as SQL.
type Statement interface {
	nodeize() (tokenizer.Tokenizer, []interface{})
}

type Clause interface {
	Statement
	self() (tokenizer.Tokenizer, []interface{})
	previous() Clause
}

type Joiner interface {
	Statement
	Join(TableOrAlias) Join
	InnerJoin(TableOrAlias) Join
	LeftJoin(TableOrAlias) Join
	RightJoin(TableOrAlias) Join
	self() (tokenizer.Tokenizer, []interface{})
	previous() Joiner
}

type TableOrAlias interface {
	Statement
	isTableOrAlias() bool
}

type TableOrAliasOrJoiner interface {
	Statement
	isTableOrAliasOrJoiner() bool
}

type ColumnOrAliasOrSubquery interface {
	Statement
	isColumnOrAliasOrSubquery() bool
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

type ComparisonOrLogicalOperation interface {
	Statement
	keyword() keyword.Operator
}
