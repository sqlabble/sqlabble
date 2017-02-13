package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type insertInto struct {
	table   tableNode
	columns []column
}

func newInsertInto(table tableNode, columns ...column) insertInto {
	return insertInto{
		table:   table,
		columns: columns,
	}
}

func (s insertInto) generator() generator.Generator {
	cs := clauseNodes(s)
	gs := make([]generator.Generator, len(cs))
	for i, c := range cs {
		gs[i] = c.clauseGenerator()
	}
	return generator.NewGenerators(gs...)
}

func (s insertInto) clauseGenerator() generator.Generator {
	es := make([]generator.Expression, len(s.columns))
	for i, c := range s.columns {
		es[i] = c.expression()
	}
	return generator.NewContainer(
		generator.NewExpression(string(keyword.InsertInto)),
		generator.NewExpressions(
			s.table.expression(),
			generator.NewArray(es...),
		),
	)
}

func (c insertInto) previous() clauseNode {
	return nil
}

func (c insertInto) Values(vals ...interface{}) values {
	f := newValues(vals...)
	f.prevClause = c
	return f
}
