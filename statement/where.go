package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
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

func (w Where) node() generator.Node {
	cs := clauseNodes(w)
	ns := make([]generator.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return generator.NewNodes(ns...)
}

func (w Where) myNode() generator.Node {
	return generator.NewContainer(
		generator.NewExpression(keyword.Where),
		w.operation.node(),
	)
}

func (w Where) previous() Clause {
	return w.prev
}

func (w Where) GroupBy(col Column, columns ...Column) GroupBy {
	g := NewGroupBy(col, columns...)
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
