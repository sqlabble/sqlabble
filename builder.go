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

func Build(n Node) (string, []interface{}) {
	return NewBuilder("", "").Build(n)
}

func BuildIndent(n Node, prefix, indent string) (string, []interface{}) {
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

func (b Builder) Build(n Node) (string, []interface{}) {
	return n.generator().Generate(b.context)
}
