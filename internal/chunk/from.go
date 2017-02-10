package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type From struct {
	prev  grammar.Clause
	table grammar.Table
}

func NewFrom(table grammar.Table) From {
	return From{
		table: table,
	}
}

func (f From) Generator() generator.Generator {
	cs := grammar.Clauses(f)
	fs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		fs[i] = c.Container()
	}
	return generator.NewGenerators(fs...)
}

func (f From) Container() generator.Container {
	return generator.NewContainer(
		generator.NewExpression(string(keyword.From)),
		f.table.Generator(),
	)
}

func (f From) Prev() grammar.Clause {
	return f.prev
}

func (f From) Where(operation grammar.Operation) Where {
	w := NewWhere(operation)
	w.prev = f
	return w
}

func (f From) OrderBy(orders ...grammar.Order) OrderBy {
	return NewOrderBy(orders...)
}
