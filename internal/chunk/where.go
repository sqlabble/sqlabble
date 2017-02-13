package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type Where struct {
	prev      grammar.Clause
	operation grammar.Operation
}

func NewWhere(operation grammar.Operation) Where {
	return Where{
		operation: operation,
	}
}

func (w Where) Generator() generator.Generator {
	cs := grammar.Clauses(w)
	gs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		gs[i] = c.ClauseGenerator()
	}
	return generator.NewGenerators(gs...)
}

func (w Where) ClauseGenerator() generator.Generator {
	return generator.NewContainer(
		generator.NewExpression(keyword.Where),
		w.operation.Generator(),
	)
}

func (w Where) Prev() grammar.Clause {
	return w.prev
}

func (w Where) OrderBy(os ...grammar.Order) OrderBy {
	o := NewOrderBy(os...)
	o.prev = w
	return o
}
