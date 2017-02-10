package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

func NewOrderBy(os ...grammar.Order) OrderBy {
	return OrderBy{orders: os}
}

type OrderBy struct {
	prev   grammar.Clause
	orders []grammar.Order
}

func (o OrderBy) Generator() generator.Generator {
	cs := grammar.Clauses(o)
	fs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		fs[i] = c.Container()
	}
	return generator.NewGenerators(fs...)
}

func (o OrderBy) Container() generator.Container {
	fs := make([]generator.Generator, len(o.orders))
	for i, c := range o.orders {
		fs[i] = c.Generator()
	}
	return generator.NewContainer(
		generator.NewExpression(string(keyword.OrderBy)),
		generator.NewComma(fs...),
	)
}

func (c OrderBy) Prev() grammar.Clause {
	return c.prev
}
