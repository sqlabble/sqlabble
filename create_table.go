package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type createTable struct {
	table tableOrTableAs
}

func newCreateTable(table tableOrTableAs) createTable {
	return createTable{
		table: table,
	}
}

func (c createTable) generator() generator.Generator {
	return c.container()
}

func (c createTable) container() generator.Container {
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
