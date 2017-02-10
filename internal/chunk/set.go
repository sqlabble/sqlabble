package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type Set struct {
	prev    grammar.Clause
	assigns []Assign
}

func NewSet(assigns ...Assign) Set {
	return Set{
		assigns: assigns,
	}
}

func (s Set) Generator() generator.Generator {
	cs := grammar.Clauses(s)
	gs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		gs[i] = c.Container()
	}
	return generator.NewGenerators(gs...)
}

func (s Set) Container() generator.Container {
	gs := make([]generator.Generator, len(s.assigns))
	for i, a := range s.assigns {
		gs[i] = a.Expression()
	}
	c := generator.NewContainer(
		generator.NewExpression(keyword.Set),
		generator.NewComma(gs...),
	)
	return c
}

func (s Set) Prev() grammar.Clause {
	return s.prev
}

func (s Set) Where(operation grammar.Operation) Where {
	w := NewWhere(operation)
	w.prev = s
	return w
}
