package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type groupBy struct {
	prev    clause
	columns []column
}

func NewGroupBy(col column, columns ...column) groupBy {
	return groupBy{
		columns: append([]column{col}, columns...),
	}
}

func (g groupBy) node() generator.Node {
	cs := clauseNodes(g)
	ns := make([]generator.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return generator.NewNodes(ns...)
}

func (g groupBy) myNode() generator.Node {
	gs := make([]generator.Node, len(g.columns))
	for i, c := range g.columns {
		gs[i] = c.node()
	}
	return generator.NewContainer(
		generator.NewExpression(string(keyword.GroupBy)),
		generator.NewComma(gs...),
	)
}

func (g groupBy) previous() clause {
	return g.prev
}

func (g groupBy) Having(operation comparisonOrLogicalOperation) having {
	l := NewHaving(operation)
	l.prev = g
	return l
}

func (g groupBy) OrderBy(orders ...order) orderBy {
	o := NewOrderBy(orders...)
	o.prev = g
	return o
}

func (g groupBy) Limit(count int) limit {
	l := NewLimit(count)
	l.prev = g
	return l
}
