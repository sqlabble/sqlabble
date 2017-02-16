package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type Limit struct {
	prev  Clause
	count int
}

func NewLimit(count int) Limit {
	return Limit{
		count: count,
	}
}

func (l Limit) node() generator.Node {
	cs := clauseNodes(l)
	ns := make([]generator.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return generator.NewNodes(ns...)
}

func (l Limit) myNode() generator.Node {
	return generator.NewContainer(
		generator.NewExpression(keyword.Limit),
		generator.ValuesToExpression(l.count),
	)
}

func (l Limit) previous() Clause {
	return l.prev
}

func (l Limit) Offset(count int) Offset {
	o := newOffset(count)
	o.prev = l
	return o
}
