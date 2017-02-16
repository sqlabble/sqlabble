package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/node"
)

// Offset skips specified rows before beginning to return rows.
type Offset struct {
	prev  Clause
	count int
}

// NewOffset return a new Offset.
func NewOffset(count int) Offset {
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
