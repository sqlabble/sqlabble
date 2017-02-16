package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type Offset struct {
	prev  Clause
	count int
}

func newOffset(count int) Offset {
	return Offset{
		count: count,
	}
}

func (o Offset) node() generator.Node {
	cs := clauseNodes(o)
	ns := make([]generator.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return generator.NewNodes(ns...)
}

func (o Offset) myNode() generator.Node {
	return generator.NewContainer(
		generator.NewExpression(keyword.Offset),
		generator.ValuesToExpression(o.count),
	)
}

func (o Offset) previous() Clause {
	return o.prev
}
