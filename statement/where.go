package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type where struct {
	prev      clause
	operation comparisonOrLogicalOperation
}

func NewWhere(operation comparisonOrLogicalOperation) where {
	return where{
		operation: operation,
	}
}

func (w where) node() generator.Node {
	cs := clauseNodes(w)
	ns := make([]generator.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return generator.NewNodes(ns...)
}

func (w where) myNode() generator.Node {
	return generator.NewContainer(
		generator.NewExpression(keyword.Where),
		w.operation.node(),
	)
}

func (w where) previous() clause {
	return w.prev
}

func (w where) GroupBy(col column, columns ...column) groupBy {
	g := NewGroupBy(col, columns...)
	g.prev = w
	return g
}

func (w where) OrderBy(orders ...order) orderBy {
	o := NewOrderBy(orders...)
	o.prev = w
	return o
}

func (w where) Limit(count int) limit {
	l := NewLimit(count)
	l.prev = w
	return l
}
