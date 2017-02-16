package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type CreateTable struct {
	ifNotExists bool
	table       Joiner
}

func NewCreateTable(table Joiner) CreateTable {
	return CreateTable{
		ifNotExists: false,
		table:       table,
	}
}

func NewCreateTableIfNotExists(table Joiner) CreateTable {
	return CreateTable{
		ifNotExists: true,
		table:       table,
	}
}

func (c CreateTable) node() generator.Node {
	return c.container()
}

func (c CreateTable) container() generator.Container {
	k := generator.NewExpression(keyword.CreateTable)
	if c.ifNotExists {
		k = k.Append(generator.NewExpression(keyword.IfNotExists))
	}
	return generator.NewContainer(
		k,
		c.table.expression(),
	)
}

func (c CreateTable) Definitions(defs ...Definition) Definitions {
	ds := NewDefinitions(defs...)
	ds.createTable = c
	return ds
}
