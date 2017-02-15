package sqlabble

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type limit struct {
	prev  clause
	count int
}

func newLimit(count int) limit {
	return limit{
		count: count,
	}
}

func (l limit) node() generator.Node {
	cs := clauseNodes(l)
	ns := make([]generator.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return generator.NewParallelNodes(ns...)
}

func (l limit) myNode() generator.Node {
	return generator.NewContainer(
		generator.NewExpression(keyword.Limit),
		generator.ValuesToExpression(l.count),
	)
}

func (l limit) previous() clause {
	return l.prev
}

func (l limit) Offset(count int) offset {
	o := newOffset(count)
	o.prev = l
	return o
}
