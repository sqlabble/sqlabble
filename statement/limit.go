package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/node"
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

func (l Limit) node() node.Node {
	cs := clauseNodes(l)
	ns := make([]node.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return node.NewNodes(ns...)
}

func (l Limit) myNode() node.Node {
	return node.NewContainer(
		node.NewExpression(keyword.Limit),
		node.ValuesToExpression(l.count),
	)
}

func (l Limit) previous() Clause {
	return l.prev
}

func (l Limit) Offset(count int) Offset {
	o := NewOffset(count)
	o.prev = l
	return o
}
