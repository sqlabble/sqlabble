package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type from struct {
	prev  clause
	table joiner
}

func NewFrom(table joiner) from {
	return from{
		table: table,
	}
}

func (f from) node() generator.Node {
	cs := clauseNodes(f)
	ns := make([]generator.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return generator.NewNodes(ns...)
}

func (f from) myNode() generator.Node {
	return generator.NewContainer(
		generator.NewExpression(string(keyword.From)),
		f.table.node(),
	)
}

func (f from) previous() clause {
	return f.prev
}

func (f from) Where(op comparisonOrLogicalOperation) where {
	w := NewWhere(op)
	w.prev = f
	return w
}

func (f from) GroupBy(col column, columns ...column) groupBy {
	g := NewGroupBy(col, columns...)
	g.prev = f
	return g
}

func (f from) OrderBy(orders ...order) orderBy {
	o := NewOrderBy(orders...)
	o.prev = f
	return o
}

func (f from) Limit(count int) limit {
	l := NewLimit(count)
	l.prev = f
	return l
}
