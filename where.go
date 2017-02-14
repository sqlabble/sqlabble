package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type where struct {
	prev      clauseNode
	operation operationNode
}

func newWhere(operation operationNode) where {
	return where{
		operation: operation,
	}
}

func (w where) generator() generator.Generator {
	cs := clauseNodes(w)
	gs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		gs[i] = c.clauseGenerator()
	}
	return generator.NewGenerators(gs...)
}

func (w where) clauseGenerator() generator.Generator {
	return generator.NewContainer(
		generator.NewExpression(keyword.Where),
		w.operation.generator(),
	)
}

func (w where) previous() clauseNode {
	return w.prev
}

func (w where) OrderBy(os ...orderNode) orderBy {
	o := newOrderBy(os...)
	o.prev = w
	return o
}
