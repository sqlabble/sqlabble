package stmt

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

type Prever interface {
	previous() Prever
	nodeizeSelf() (tokenizer.Tokenizer, []interface{})
}

type Joiner interface {
	Statement
	Join(TableOrAlias) Join
	InnerJoin(TableOrAlias) Join
	LeftJoin(TableOrAlias) Join
	RightJoin(TableOrAlias) Join
}

type Namer interface {
	Statement
	name() string
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

type ColOrAlias interface {
	Statement
	isColOrAlias() bool
}

type JoinerOrAlias interface {
	Statement
	isJoinerOrAlias() bool
}

type ColOrSub interface {
	Statement
	isColOrSub() bool
}

type ColOrAliasOrSub interface {
	Statement
	isColOrAliasOrSub() bool
}

type ValOrColOrAliasOrFuncOrSubOrFormula interface {
	Statement
	isValOrColOrAliasOrFuncOrSubOrFormula() bool
}

type ValOrColOrSubOrFormula interface {
	Statement
	isValOrColOrSubOrFormula() bool
}

type ComparisonOrLogicalOperation interface {
	Statement
	keyword() keyword.Operator
}

type SearchedThenOrElse interface {
	Statement
	isSearchedThenOrElse() bool
}

type SimpleThenOrElse interface {
	Statement
	isSimpleThenOrElse() bool
}

type Selector interface {
	Select(columns ...ValOrColOrAliasOrFuncOrSubOrFormula) Select
}

type Fromer interface {
	From(table JoinerOrAlias) From
}

type Wherer interface {
	Where(op ComparisonOrLogicalOperation) Where
}

type GroupByer interface {
	GroupBy(col Column, columns ...Column) GroupBy
}

type OrderByer interface {
	OrderBy(orders ...Order) OrderBy
}

type Limiter interface {
	Limit(count int) Limit
}

type Offsetter interface {
	Offset(count int) Offset
}
