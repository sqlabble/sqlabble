package token

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type Delete struct{}

func NewDelete() Delete {
	return Delete{}
}

func (d Delete) Generator() generator.Generator {
	cs := grammar.Clauses(d)
	fs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		fs[i] = c.Container()
	}
	return generator.NewGenerators(fs...)
}

func (d Delete) Container() generator.Container {
	return generator.NewContainer(
		generator.NewExpression(keyword.Delete),
	)
}

func (d Delete) Prev() grammar.Clause {
	return nil
}

func (d Delete) From(t Table) From {
	f := NewFrom(t)
	f.prev = d
	return f
}
