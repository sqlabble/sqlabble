package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type having struct {
	prev      clause
	operation comparisonOrLogicalOperation
}

func newHaving(operation comparisonOrLogicalOperation) having {
	return having{
		operation: operation,
	}
}

func (h having) node() generator.Node {
	cs := clauseNodes(h)
	ns := make([]generator.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return generator.NewNodes(ns...)
}

func (h having) myNode() generator.Node {
	return generator.NewContainer(
		generator.NewExpression(string(keyword.Having)),
		h.operation.node(),
	)
}

func (h having) previous() clause {
	return h.prev
}

func (h having) OrderBy(orders ...order) orderBy {
	o := newOrderBy(orders...)
	o.prev = h
	return o
}

func (h having) Limit(count int) limit {
	l := newLimit(count)
	l.prev = h
	return l
}
