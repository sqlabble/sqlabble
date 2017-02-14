package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type from struct {
	prev  clause
	table tableOrTableAs
}

func newFrom(table tableOrTableAs) from {
	return from{
		table: table,
	}
}

func (f from) node() generator.Node {
	cs := clauseNodes(f)
	fs := make([]generator.Node, len(cs))
	for i, c := range cs {
		fs[i] = c.nodeMine()
	}
	return generator.NewNodes(fs...)
}

func (f from) nodeMine() generator.Node {
	return generator.NewContainer(
		generator.NewExpression(string(keyword.From)),
		f.table.node(),
	)
}

func (f from) previous() clause {
	return f.prev
}

func (f from) Where(op comparisonOrLogicalOperation) where {
	w := newWhere(op)
	w.prev = f
	return w
}

func (f from) OrderBy(orders ...order) orderBy {
	o := newOrderBy(orders...)
	o.prev = f
	return o
}

func (f from) GroupBy(col column, columns ...column) groupBy {
	g := newGroupBy(col, columns...)
	g.prev = f
	return g
}
