package sqlabble

import (
	"github.com/sqlabble/sqlabble/stmt"
)

// Methods exported to make statements.
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
	Table    = stmt.NewTable
	Val      = stmt.NewVal
	Vals     = stmt.NewVals
	Subquery = stmt.NewSubquery

	Assign    = stmt.NewAssign
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

	Wildcard = stmt.NewWildcard()

	Count = stmt.NewCount
	Max   = stmt.NewMax

	Add        = stmt.NewAdd
	Sub        = stmt.NewSub
	Mul        = stmt.NewMul
	Div        = stmt.NewDiv
	IntegerDiv = stmt.NewIntegerDiv
	Mod        = stmt.NewMod

	IFNULL = stmt.NewIFNULL
)
