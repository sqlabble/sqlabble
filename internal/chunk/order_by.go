package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type OrderBy struct {
	prev   grammar.Clause
	orders []grammar.Order
}

func NewOrderBy(os ...grammar.Order) OrderBy {
	return OrderBy{orders: os}
}

func (o OrderBy) Generator() generator.Generator {
	cs := grammar.Clauses(o)
	fs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		fs[i] = c.ClauseGenerator()
	}
	return generator.NewGenerators(fs...)
}

func (o OrderBy) ClauseGenerator() generator.Generator {
	fs := make([]generator.Generator, len(o.orders))
	for i, c := range o.orders {
		fs[i] = c.Generator()
	}
	return generator.NewContainer(
		generator.NewExpression(string(keyword.OrderBy)),
		generator.NewComma(fs...),
	)
}

func (o OrderBy) Prev() grammar.Clause {
	return o.prev
}

func (o OrderBy) Limit(offset, limit int) Limit {
	l := NewLimit(offset, limit)
	l.prev = o
	return l
}
