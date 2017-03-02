package sqlabble

import (
	"github.com/minodisk/sqlabble/stmt"
)

// Methods exported to make stmts.
var (
	CreateTable            = stmt.NewCreateTable
	CreateTableIfNotExists = stmt.NewCreateTableIfNotExists
	Select                 = stmt.NewSelect
	SelectDistinct         = stmt.NewSelectDistinct
	InsertInto             = stmt.NewInsertInto
	Update                 = stmt.NewUpdate
	Delete                 = stmt.NewDelete

	SimpleCase   = stmt.NewSimpleCase
	SimpleWhen   = stmt.NewSimpleWhen
	SearchedCase = stmt.NewSearchedCase
	SearchedWhen = stmt.NewSearchedWhen

	Column   = stmt.NewColumn
	C        = stmt.NewColumn
	Table    = stmt.NewTable
	T        = stmt.NewTable
	Param    = stmt.NewParam
	P        = stmt.NewParam
	Params   = stmt.NewParams
	Ps       = stmt.NewParams
	Subquery = stmt.NewSubquery
	S        = stmt.NewSubquery

	And       = stmt.NewAnd
	Or        = stmt.NewOr
	Not       = stmt.NewNot
	Exists    = stmt.NewExists
	NotExists = stmt.NewNotExists

	Union        = stmt.NewUnion
	UnionAll     = stmt.NewUnionAll
	Intersect    = stmt.NewIntersect
	IntersectAll = stmt.NewIntersectAll
	Except       = stmt.NewExcept
	ExceptAll    = stmt.NewExceptAll
)
