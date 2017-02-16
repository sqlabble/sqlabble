package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/node"
)

type From struct {
	prev  Clause
	table Joiner
}

func NewFrom(table Joiner) From {
	return From{
		table: table,
	}
}

func (f From) node() node.Node {
	cs := clauseNodes(f)
	ns := make([]node.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return node.NewNodes(ns...)
}

func (f From) myNode() node.Node {
	return node.NewContainer(
		node.NewExpression(string(keyword.From)),
		f.table.node(),
	)
}

func (f From) previous() Clause {
	return f.prev
}

func (f From) Where(op ComparisonOrLogicalOperation) Where {
	w := NewWhere(op)
	w.prev = f
	return w
}

func (f From) GroupBy(col Column, columns ...Column) GroupBy {
	g := NewGroupBy(col, columns...)
	g.prev = f
	return g
}

func (f From) OrderBy(orders ...Order) OrderBy {
	o := NewOrderBy(orders...)
	o.prev = f
	return o
}

func (f From) Limit(count int) Limit {
	l := NewLimit(count)
	l.prev = f
	return l
}
