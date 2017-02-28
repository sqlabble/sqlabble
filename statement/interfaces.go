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

type Joiner interface {
	Statement
	Join(TableOrAlias) *Join
	InnerJoin(TableOrAlias) *Join
	LeftJoin(TableOrAlias) *Join
	RightJoin(TableOrAlias) *Join
}

type ValOrFuncOrSub interface {
	Statement
	isValOrFuncOrSub() bool
}

type ValsOrSub interface {
	Statement
	isValsOrSub() bool
}

type ValOrColOrFuncOrSub interface {
	Statement
	isValOrColOrFuncOrSub() bool
}

type TableOrAlias interface {
	Statement
	isTableOrAlias() bool
}

type TableOrAliasOrJoiner interface {
	Statement
	isTableOrAliasOrJoiner() bool
}

type ColOrSub interface {
	Statement
	isColOrSub() bool
}

type ColOrAliasOrSub interface {
	Statement
	isColOrAliasOrSub() bool
}

type ColOrAliasOrFuncOrSub interface {
	Statement
	isColOrAliasOrFuncOrSub() bool
}

type ComparisonOrLogicalOperation interface {
	Statement
	keyword() keyword.Operator
}
