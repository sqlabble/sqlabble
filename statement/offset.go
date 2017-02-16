package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/node"
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

func (o Offset) node() node.Node {
	cs := clauseNodes(o)
	ns := make([]node.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return node.NewNodes(ns...)
}

func (o Offset) myNode() node.Node {
	return node.NewContainer(
		node.NewExpression(keyword.Offset),
		node.ValuesToExpression(o.count),
	)
}

func (o Offset) previous() Clause {
	return o.prev
}
