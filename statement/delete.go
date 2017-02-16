package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/node"
)

type Delete struct{}

func NewDelete() Delete {
	return Delete{}
}

func (d Delete) node() node.Node {
	cs := clauseNodes(d)
	fs := make([]node.Node, len(cs))
	for i, c := range cs {
		fs[i] = c.myNode()
	}
	return node.NewNodes(fs...)
}

func (d Delete) myNode() node.Node {
	return node.NewContainer(
		node.NewExpression(keyword.Delete),
	)
}

func (d Delete) previous() Clause {
	return nil
}

func (d Delete) From(t Table) From {
	f := NewFrom(t)
	f.prev = d
	return f
}
