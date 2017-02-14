package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
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

func (s set) generator() generator.Generator {
	cs := clauseNodes(s)
	gs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		gs[i] = c.clauseGenerator()
	}
	return generator.NewGenerators(gs...)
}

func (s set) clauseGenerator() generator.Generator {
	gs := make([]generator.Generator, len(s.assigns))
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
