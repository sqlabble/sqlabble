package sqlabble

import "github.com/minodisk/sqlabble/internal/chunk"

var (
	CreateTable = chunk.NewCreateTable
	Select      = chunk.NewSelect
	InsertInto  = chunk.NewInsertInto
	Update      = chunk.NewUpdate
	Delete      = chunk.NewDelete

	Column = chunk.NewColumn
	C      = Column
	Table  = chunk.NewTable
	T      = Table

	And = chunk.NewAnd
	Or  = chunk.NewOr
	Not = chunk.NewNot
)
