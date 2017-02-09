package sqlabble

import "github.com/minodisk/sqlabble/internal/token"

var (
	Select     = token.NewSelect
	InsertInto = token.NewInsertInto

	Column = token.NewColumn
	C      = Column
	Table  = token.NewTable
	T      = Table

	And = token.NewAnd
	Or  = token.NewOr
)
