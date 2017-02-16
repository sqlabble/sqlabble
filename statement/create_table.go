package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/node"
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

func (c CreateTable) node() node.Node {
	return c.container()
}

func (c CreateTable) container() node.Container {
	k := node.NewExpression(keyword.CreateTable)
	if c.ifNotExists {
		k = k.Append(node.NewExpression(keyword.IfNotExists))
	}
	return node.NewContainer(
		k,
		c.table.expression(),
	)
}

func (c CreateTable) Definitions(defs ...Definition) Definitions {
	ds := NewDefinitions(defs...)
	ds.createTable = c
	return ds
}
