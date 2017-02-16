package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/node"
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

func (g GroupBy) node() node.Node {
	cs := clauseNodes(g)
	ns := make([]node.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return node.NewNodes(ns...)
}

func (g GroupBy) myNode() node.Node {
	gs := make([]node.Node, len(g.columns))
	for i, c := range g.columns {
		gs[i] = c.node()
	}
	return node.NewContainer(
		node.NewExpression(string(keyword.GroupBy)),
		node.NewComma(gs...),
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
