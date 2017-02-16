package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/node"
)

type Where struct {
	prev      Clause
	operation ComparisonOrLogicalOperation
}

func NewWhere(operation ComparisonOrLogicalOperation) Where {
	return Where{
		operation: operation,
	}
}

func (w Where) node() node.Node {
	cs := clauseNodes(w)
	ns := make([]node.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return node.NewNodes(ns...)
}

func (w Where) myNode() node.Node {
	return node.NewContainer(
		node.NewExpression(keyword.Where),
		w.operation.node(),
	)
}

func (w Where) previous() Clause {
	return w.prev
}

func (w Where) GroupBy(column Column, columns ...Column) GroupBy {
	g := NewGroupBy(column, columns...)
	g.prev = w
	return g
}

func (w Where) OrderBy(orders ...Order) OrderBy {
	o := NewOrderBy(orders...)
	o.prev = w
	return o
}

func (w Where) Limit(count int) Limit {
	l := NewLimit(count)
	l.prev = w
	return l
}
