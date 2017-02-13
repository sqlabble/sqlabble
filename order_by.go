package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type orderBy struct {
	prev   clauseNode
	orders []orderNode
}

func newOrderBy(os ...orderNode) orderBy {
	return orderBy{orders: os}
}

func (o orderBy) generator() generator.Generator {
	cs := clauseNodes(o)
	fs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		fs[i] = c.clauseGenerator()
	}
	return generator.NewGenerators(fs...)
}

func (o orderBy) clauseGenerator() generator.Generator {
	fs := make([]generator.Generator, len(o.orders))
	for i, c := range o.orders {
		fs[i] = c.generator()
	}
	return generator.NewContainer(
		generator.NewExpression(string(keyword.OrderBy)),
		generator.NewComma(fs...),
	)
}

func (o orderBy) previous() clauseNode {
	return o.prev
}

func (o orderBy) Limit(offset, lim int) limit {
	l := newLimit(offset, lim)
	l.prev = o
	return l
}
