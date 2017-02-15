package sqlabble

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type set struct {
	prev    clause
	assigns []assign
}

func newSet(assigns ...assign) set {
	return set{
		assigns: assigns,
	}
}

func (s set) node() generator.Node {
	cs := clauseNodes(s)
	gs := make([]generator.Node, len(cs))
	for i, c := range cs {
		gs[i] = c.myNode()
	}
	return generator.NewParallelNodes(gs...)
}

func (s set) myNode() generator.Node {
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

func (s set) previous() clause {
	return s.prev
}

func (s set) Where(operation comparisonOrLogicalOperation) where {
	w := newWhere(operation)
	w.prev = s
	return w
}
