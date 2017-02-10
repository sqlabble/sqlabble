package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type Select struct {
	columns []grammar.Column
}

func NewSelect(columns ...grammar.Column) Select {
	return Select{columns: columns}
}

func (s Select) Generator() generator.Generator {
	cs := grammar.Clauses(s)
	fs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		fs[i] = c.Container()
	}
	return generator.NewGenerators(fs...)
}

func (s Select) Container() generator.Container {
	fs := make([]generator.Generator, len(s.columns))
	for i, c := range s.columns {
		fs[i] = c.Generator()
	}
	return generator.NewContainer(
		generator.NewExpression(string(keyword.Select)),
		generator.NewComma(fs...),
	)
}

func (c Select) Prev() grammar.Clause {
	return nil
}

func (c Select) From(t grammar.Table) From {
	f := NewFrom(t)
	f.prev = c
	return f
}
