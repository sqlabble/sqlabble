package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type CreateTable struct {
	table grammar.Table
}

func NewCreateTable(table grammar.Table) CreateTable {
	return CreateTable{
		table: table,
	}
}

func (c CreateTable) Generator() generator.Generator {
	return c.Container()
}

func (c CreateTable) Container() generator.Container {
	return generator.NewContainer(
		generator.NewExpression(string(keyword.CreateTable)),
		c.table.Expression(),
	)
}

func (c CreateTable) Definitions(definitions ...Definition) Definitions {
	ds := NewDefinitions(definitions...)
	ds.createTable = c
	return ds
}
