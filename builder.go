package sqlabble

import "github.com/minodisk/sqlabble/internal/generator"

// Methods exported to make statements.
var (
	CreateTable    = newCreateTable
	Select         = newSelect
	SelectDistinct = newSelectDistinct
	InsertInto     = newInsertInto
	Update         = newUpdate
	Delete         = newDelete

	Column = newColumn
	C      = newColumn
	Table  = newTable
	T      = newTable

	And = newAnd
	Or  = newOr
	Not = newNot

	Union        = newUnion
	UnionAll     = newUnionAll
	Intersect    = newIntersect
	IntersectAll = newIntersectAll
	Except       = newExcept
	ExceptAll    = newExceptAll
)

// Typical builders commonly used to build queries.
var (
	Standard         = NewBuilder(generator.Options{})
	IndentedStandard = NewBuilder(generator.Options{
		Indent: "    ",
	})
	MySQL4 = NewBuilder(generator.Options{
		FlatSets: true,
	})
	IndentedMySQL4 = NewBuilder(generator.Options{
		Indent:   "    ",
		FlatSets: true,
	})
)

// Build builds a query.
func Build(s Statement) (string, []interface{}) {
	return Standard.Build(s)
}

// BuildIndent builds an indented query.
func BuildIndent(s Statement) (string, []interface{}) {
	return IndentedStandard.Build(s)
}

// Builder is a container for storing options
// so that you can build a query with the same options
// over and over again.
type Builder struct {
	context generator.Context
}

// NewBuilder returns a Builder with a specified options.
func NewBuilder(options generator.Options) Builder {
	return Builder{
		context: options.ToContext(),
	}
}

// Build converts a statement into a query and a slice of values.
func (b Builder) Build(statement Statement) (string, []interface{}) {
	return statement.node().ToSQL(b.context)
}
