package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type createTable struct {
	table tableNode
}

func newCreateTable(table tableNode) createTable {
	return createTable{
		table: table,
	}
}

func (c createTable) generator() generator.Generator {
	return c.Container()
}

func (c createTable) Container() generator.Container {
	return generator.NewContainer(
		generator.NewExpression(string(keyword.CreateTable)),
		c.table.expression(),
	)
}

func (c createTable) Definitions(defs ...definition) definitions {
	ds := newDefinitions(defs...)
	ds.createTable = c
	return ds
}
