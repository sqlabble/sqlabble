package sqlabble

import "github.com/minodisk/sqlabble/internal/token"

var (
	CreateTable = token.NewCreateTable
	Select      = token.NewSelect
	InsertInto  = token.NewInsertInto
	Update      = token.NewUpdate

	Column = token.NewColumn
	C      = Column
	Table  = token.NewTable
	T      = Table

	And = token.NewAnd
	Or  = token.NewOr
	Not = token.NewNot
)
