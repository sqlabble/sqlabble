package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type offset struct {
	prev  clause
	count int
}

func newOffset(count int) offset {
	return offset{
		count: count,
	}
}

func (o offset) node() generator.Node {
	cs := clauseNodes(o)
	ns := make([]generator.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return generator.NewNodes(ns...)
}

func (o offset) myNode() generator.Node {
	return generator.NewContainer(
		generator.NewExpression(keyword.Offset),
		generator.ValuesToExpression(o.count),
	)
}

func (o offset) previous() clause {
	return o.prev
}
