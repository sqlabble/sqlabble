package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type InsertInto struct {
	table   grammar.Table
	columns []Column
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
		gs[i] = c.ClauseGenerator()
	}
	return generator.NewGenerators(gs...)
}

func (s InsertInto) ClauseGenerator() generator.Generator {
	es := make([]generator.Expression, len(s.columns))
	for i, c := range s.columns {
		es[i] = c.Expression()
	}
	return generator.NewContainer(
		generator.NewExpression(string(keyword.InsertInto)),
		generator.NewExpressions(
			s.table.Expression(),
			generator.NewArray(es...),
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
