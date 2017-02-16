package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/node"
)

type Set struct {
	prev    Clause
	assigns []Assign
}

func NewSet(assigns ...Assign) Set {
	return Set{
		assigns: assigns,
	}
}

func (s Set) node() node.Node {
	cs := clauseNodes(s)
	gs := make([]node.Node, len(cs))
	for i, c := range cs {
		gs[i] = c.myNode()
	}
	return node.NewNodes(gs...)
}

func (s Set) myNode() node.Node {
	gs := make([]node.Node, len(s.assigns))
	for i, a := range s.assigns {
		gs[i] = a.expression()
	}
	c := node.NewContainer(
		node.NewExpression(keyword.Set),
		node.NewComma(gs...),
	)
	return c
}

func (s Set) previous() Clause {
	return s.prev
}

func (s Set) Where(operation ComparisonOrLogicalOperation) Where {
	w := NewWhere(operation)
	w.prev = s
	return w
}
