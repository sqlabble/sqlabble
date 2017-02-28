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

	SimpleCase   = statement.NewSimpleCase
	SimpleWhen   = statement.NewSimpleWhen
	SearchedCase = statement.NewSearchedCase
	SearchedWhen = statement.NewSearchedWhen

	Column   = statement.NewColumn
	C        = statement.NewColumn
	Table    = statement.NewTable
	T        = statement.NewTable
	Param    = statement.NewParam
	P        = statement.NewParam
	Params   = statement.NewParams
	Ps       = statement.NewParams
	Subquery = statement.NewSubquery
	S        = statement.NewSubquery

	And       = statement.NewAnd
	Or        = statement.NewOr
	Not       = statement.NewNot
	Exists    = statement.NewExists
	NotExists = statement.NewNotExists

	Union        = statement.NewUnion
	UnionAll     = statement.NewUnionAll
	Intersect    = statement.NewIntersect
	IntersectAll = statement.NewIntersectAll
	Except       = statement.NewExcept
	ExceptAll    = statement.NewExceptAll
)
