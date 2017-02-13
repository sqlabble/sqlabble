package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type Having struct {
	prev      grammar.Clause
	operation grammar.Operation
}

func NewHaving(operation grammar.Operation) Having {
	return Having{
		operation: operation,
	}
}

func (w Having) Generator() generator.Generator {
	cs := grammar.Clauses(w)
	fs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		fs[i] = c.ClauseGenerator()
	}
	return generator.NewGenerators(fs...)
}

func (w Having) ClauseGenerator() generator.Generator {
	return generator.NewContainer(
		generator.NewExpression(string(keyword.Having)),
		w.operation.Generator(),
	)
}

func (c Having) Child() grammar.Statement {
	return c.operation
}

func (c Having) Prev() grammar.Clause {
	return c.prev
}
