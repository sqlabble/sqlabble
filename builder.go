package sqlabble

import "github.com/minodisk/sqlabble/generator"

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

func Build(s Statement) (string, []interface{}) {
	return Standard.Build(s)
}

func BuildIndent(s Statement) (string, []interface{}) {
	return IndentedStandard.Build(s)
}

type Builder struct {
	context generator.Context
}

func NewBuilder(o generator.Options) Builder {
	return Builder{
		context: o.ToContext(),
	}
}

func (b Builder) Build(n Statement) (string, []interface{}) {
	return n.node().ToSQL(b.context)
}
