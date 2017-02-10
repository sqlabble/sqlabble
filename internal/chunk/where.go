package token

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type Where struct {
	prev      grammar.Clause
	operation grammar.Operation
}

func NewWhere(c grammar.Operation) Where {
	return Where{operation: c}
}

func (w Where) Generator() generator.Generator {
	cs := grammar.Clauses(w)
	fs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		fs[i] = c.Container()
	}
	return generator.NewGenerators(fs...)
}

func (w Where) Container() generator.Container {
	return generator.NewContainer(
		generator.NewExpression(string(keyword.Where)),
		w.operation.Generator(),
	)
}

func (c Where) Child() grammar.Statement {
	return c.operation
}

func (c Where) Prev() grammar.Clause {
	return c.prev
}

func (c Where) OrderBy(os ...grammar.Order) OrderBy {
	o := NewOrderBy(os...)
	o.prev = c
	return o
}
