package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/chunk"
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
)

var (
	CreateTable    = chunk.NewCreateTable
	Select         = chunk.NewSelect
	SelectDistinct = chunk.NewSelectDistinct
	InsertInto     = chunk.NewInsertInto
	Update         = chunk.NewUpdate
	Delete         = chunk.NewDelete

	Column = chunk.NewColumn
	C      = Column
	Table  = chunk.NewTable
	T      = Table

	And = chunk.NewAnd
	Or  = chunk.NewOr
	Not = chunk.NewNot

	Union = chunk.NewUnion
)

func Build(s grammar.Statement) (string, []interface{}) {
	return NewBuilder("", "").Build(s)
}

func BuildIndent(s grammar.Statement, prefix, indent string) (string, []interface{}) {
	return NewBuilder(prefix, indent).Build(s)
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
			SetFlatUnion(true),
	}
}

func (b Builder) Build(c grammar.Statement) (string, []interface{}) {
	return c.Generator().Generate(b.context)
}
