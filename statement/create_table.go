package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type createTable struct {
	ifNotExists bool
	table       joiner
}

func NewCreateTable(table joiner) createTable {
	return createTable{
		ifNotExists: false,
		table:       table,
	}
}

func NewCreateTableIfNotExists(table joiner) createTable {
	return createTable{
		ifNotExists: true,
		table:       table,
	}
}

func (c createTable) node() generator.Node {
	return c.container()
}

func (c createTable) container() generator.Container {
	k := generator.NewExpression(keyword.CreateTable)
	if c.ifNotExists {
		k = k.Append(generator.NewExpression(keyword.IfNotExists))
	}
	return generator.NewContainer(
		k,
		c.table.expression(),
	)
}

func (c createTable) Definitions(defs ...definition) definitions {
	ds := newDefinitions(defs...)
	ds.createTable = c
	return ds
}
