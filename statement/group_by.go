package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type GroupBy struct {
	prev    Clause
	columns []Column
}

func NewGroupBy(column Column, columns ...Column) GroupBy {
	return GroupBy{
		columns: append([]Column{column}, columns...),
	}
}

func (g GroupBy) node() generator.Node {
	cs := clauseNodes(g)
	ns := make([]generator.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return generator.NewNodes(ns...)
}

func (g GroupBy) myNode() generator.Node {
	gs := make([]generator.Node, len(g.columns))
	for i, c := range g.columns {
		gs[i] = c.node()
	}
	return generator.NewContainer(
		generator.NewExpression(string(keyword.GroupBy)),
		generator.NewComma(gs...),
	)
}

func (g GroupBy) previous() Clause {
	return g.prev
}

func (g GroupBy) Having(operation ComparisonOrLogicalOperation) Having {
	l := NewHaving(operation)
	l.prev = g
	return l
}

func (g GroupBy) OrderBy(orders ...Order) OrderBy {
	o := NewOrderBy(orders...)
	o.prev = g
	return o
}

func (g GroupBy) Limit(count int) Limit {
	l := NewLimit(count)
	l.prev = g
	return l
}
