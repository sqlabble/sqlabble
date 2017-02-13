package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type Select struct {
	distinct bool
	columns  []grammar.Column
}

func NewSelect(columns ...grammar.Column) Select {
	return Select{
		columns: columns,
	}
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
	k := generator.NewExpression(keyword.Select)
	if s.distinct {
		k = k.Append(generator.NewExpression(keyword.Distinct))
	}
	return generator.NewContainer(
		k,
		generator.NewComma(fs...),
	)
}

func (c Select) Prev() grammar.Clause {
	return nil
}

func (c Select) Distinct(columns ...grammar.Column) Select {
	c.columns = append(c.columns, columns...)
	c.distinct = true
	return c
}

func (c Select) From(t grammar.Table) From {
	f := NewFrom(t)
	f.prev = c
	return f
}
