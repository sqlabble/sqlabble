package sqlabble

import "github.com/minodisk/sqlabble/internal/generator"

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

var (
	Standard         = Builder{generator.NewContext("", "")}
	IndentedStandard = Builder{generator.NewContext("", "    ")}
	MySQL4           = Builder{generator.NewContext("", "")}
	IndentedMySQL4   = Builder{generator.NewContext("", "    ")}
)

func Build(n Statement) (string, []interface{}) {
	return NewBuilder("", "").Build(n)
}

func BuildIndent(n Statement, prefix, indent string) (string, []interface{}) {
	return NewBuilder(prefix, indent).Build(n)
}

type Builder struct {
	context generator.Context
}

func NewBuilder(prefix, indent string) Builder {
	return Builder{
		context: generator.NewContext(prefix, indent),
	}
}

func NewBuilderForMySQL4(prefix, indent string) Builder {
	return Builder{
		context: generator.
			NewContext(prefix, indent).
			SetFlatSetOperation(true),
	}
}

func (b Builder) Build(n Statement) (string, []interface{}) {
	return n.node().ToSQL(b.context)
}
