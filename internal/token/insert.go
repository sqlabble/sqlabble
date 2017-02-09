package token

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type InsertInto struct {
	table   grammar.Table
	columns Columns
}

func NewInsertInto(table grammar.Table, columns ...Column) InsertInto {
	return InsertInto{
		table:   table,
		columns: columns,
	}
}

func (s InsertInto) Generator() generator.Generator {
	cs := grammar.Clauses(s)
	gs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		gs[i] = c.Container()
	}
	return generator.NewGenerators(gs...)
}

func (s InsertInto) Container() generator.Container {
	return generator.NewContainer(
		generator.NewExpression(string(keyword.InsertInto)),
		generator.NewJoin(
			" ",
			s.table.Generator(),
			s.columns.Generator(),
		),
	)
}

func (c InsertInto) Prev() grammar.Clause {
	return nil
}

func (c InsertInto) Values(values ...interface{}) Values {
	f := NewValues(values...)
	f.clause = c
	return f
}
