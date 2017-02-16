package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
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

func (s Set) node() generator.Node {
	cs := clauseNodes(s)
	gs := make([]generator.Node, len(cs))
	for i, c := range cs {
		gs[i] = c.myNode()
	}
	return generator.NewNodes(gs...)
}

func (s Set) myNode() generator.Node {
	gs := make([]generator.Node, len(s.assigns))
	for i, a := range s.assigns {
		gs[i] = a.expression()
	}
	c := generator.NewContainer(
		generator.NewExpression(keyword.Set),
		generator.NewComma(gs...),
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
