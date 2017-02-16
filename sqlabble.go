package sqlabble

import (
	"github.com/minodisk/sqlabble/builder"
	"github.com/minodisk/sqlabble/statement"
)

// Methods exported to make statements.
var (
	CreateTable    = statement.NewCreateTable
	Select         = statement.NewSelect
	SelectDistinct = statement.NewSelectDistinct
	InsertInto     = statement.NewInsertInto
	Update         = statement.NewUpdate
	Delete         = statement.NewDelete

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

// Typical builders commonly used to build queries.
var (
	Standard         = builder.NewBuilder(builder.Options{})
	IndentedStandard = builder.NewBuilder(builder.Options{
		Indent: "    ",
	})
	MySQL4 = builder.NewBuilder(builder.Options{
		FlatSets: true,
	})
	IndentedMySQL4 = builder.NewBuilder(builder.Options{
		Indent:   "    ",
		FlatSets: true,
	})
)

// Build builds a query.
func Build(s statement.Statement) (string, []interface{}) {
	return Standard.Build(s)
}

// BuildIndent builds an indented query.
func BuildIndent(s statement.Statement) (string, []interface{}) {
	return IndentedStandard.Build(s)
}
