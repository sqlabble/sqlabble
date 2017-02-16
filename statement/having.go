package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type Having struct {
	prev      Clause
	operation ComparisonOrLogicalOperation
}

func NewHaving(operation ComparisonOrLogicalOperation) Having {
	return Having{
		operation: operation,
	}
}

func (h Having) node() generator.Node {
	cs := clauseNodes(h)
	ns := make([]generator.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return generator.NewNodes(ns...)
}

func (h Having) myNode() generator.Node {
	return generator.NewContainer(
		generator.NewExpression(string(keyword.Having)),
		h.operation.node(),
	)
}

func (h Having) previous() Clause {
	return h.prev
}

func (h Having) OrderBy(orders ...Order) OrderBy {
	o := NewOrderBy(orders...)
	o.prev = h
	return o
}

func (h Having) Limit(count int) Limit {
	l := NewLimit(count)
	l.prev = h
	return l
}
