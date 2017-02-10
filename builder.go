package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/chunk"
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
)

var (
	CreateTable = chunk.NewCreateTable
	Select      = chunk.NewSelect
	InsertInto  = chunk.NewInsertInto
	Update      = chunk.NewUpdate
	Delete      = chunk.NewDelete

	Column = chunk.NewColumn
	C      = Column
	Table  = chunk.NewTable
	T      = Table

	And = chunk.NewAnd
	Or  = chunk.NewOr
	Not = chunk.NewNot
)

func Build(c grammar.Statement) (string, []interface{}) {
	return c.Generator().Generate(generator.NonBreakingContext)
}

func BuildIndent(c grammar.Statement, prefix, indent string) (string, []interface{}) {
	return c.Generator().Generate(generator.NewContext(prefix, indent))
}
