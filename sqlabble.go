package sqlabble

import "github.com/minodisk/sqlabble/statement"

// Methods exported to make statements.
var (
	CreateTable            = statement.NewCreateTable
	CreateTableIfNotExists = statement.NewCreateTableIfNotExists
	Select                 = statement.NewSelect
	SelectDistinct         = statement.NewSelectDistinct
	InsertInto             = statement.NewInsertInto
	Update                 = statement.NewUpdate
	Delete                 = statement.NewDelete

	Column = statement.NewColumn
	C      = statement.NewColumn
	Table  = statement.NewTable
	T      = statement.NewTable

	And = statement.NewAnd
	Or  = statement.NewOr
	Not = statement.NewNot

	Union        = statement.NewUnion
	UnionAll     = statement.NewUnionAll
	Intersect    = statement.NewIntersect
	IntersectAll = statement.NewIntersectAll
	Except       = statement.NewExcept
	ExceptAll    = statement.NewExceptAll
)
