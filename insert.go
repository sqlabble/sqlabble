package sqlabble

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type insertInto struct {
	table   tableOrTableAs
	columns []column
}

func newInsertInto(table tableOrTableAs, columns ...column) insertInto {
	return insertInto{
		table:   table,
		columns: columns,
	}
}

func (s insertInto) node() generator.Node {
	cs := clauseNodes(s)
	ns := make([]generator.Node, len(cs))
	for i, c := range cs {
		ns[i] = c.myNode()
	}
	return generator.NewNodes(ns...)
}

func (s insertInto) myNode() generator.Node {
	es := make([]generator.Expression, len(s.columns))
	for i, c := range s.columns {
		es[i] = c.expression()
	}
	return generator.NewContainer(
		generator.NewExpression(string(keyword.InsertInto)),
		generator.JoinExpressions(
			s.table.expression(),
			generator.ArrayToExpression(es...),
		),
	)
}

func (c insertInto) previous() clause {
	return nil
}

func (c insertInto) Values(vals ...interface{}) values {
	f := newValues(vals...)
	f.prevClause = c
	return f
}
