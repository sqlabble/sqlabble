package sqlabble

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type where struct {
	prev      clause
	operation comparisonOrLogicalOperation
}

func newWhere(operation comparisonOrLogicalOperation) where {
	return where{
		operation: operation,
	}
}

func (w where) node() generator.Node {
	cs := clauseNodes(w)
	gs := make([]generator.Node, len(cs))
	for i, c := range cs {
		gs[i] = c.nodeMine()
	}
	return generator.NewNodes(gs...)
}

func (w where) nodeMine() generator.Node {
	return generator.NewContainer(
		generator.NewExpression(keyword.Where),
		w.operation.node(),
	)
}

func (w where) previous() clause {
	return w.prev
}

func (w where) OrderBy(os ...order) orderBy {
	o := newOrderBy(os...)
	o.prev = w
	return o
}
