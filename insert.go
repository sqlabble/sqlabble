package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type insertInto struct {
	table   joiner
	columns []column
}

func newInsertInto(table joiner, columns ...column) insertInto {
	return insertInto{
		table:   table,
		columns: columns,
	}
}

func (i insertInto) node() generator.Node {
	cs := clauseNodes(i)
	ns := make([]generator.Node, len(cs))
	for j, c := range cs {
		ns[j] = c.myNode()
	}
	return generator.NewParallelNodes(ns...)
}

func (i insertInto) myNode() generator.Node {
	es := make([]generator.Expression, len(i.columns))
	for j, c := range i.columns {
		es[j] = c.expression()
	}
	return generator.NewContainer(
		generator.NewExpression(string(keyword.InsertInto)),
		generator.JoinExpressions(
			i.table.expression(),
			generator.ArrayToExpression(es...),
		),
	)
}

func (i insertInto) previous() clause {
	return nil
}

func (i insertInto) Values(vals ...interface{}) values {
	v := newValues(vals...)
	v.prevClause = i
	return v
}

func (i insertInto) DefaultValues() defaultValues {
	v := newDefaultValues()
	v.prev = i
	return v
}
