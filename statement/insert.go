package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type InsertInto struct {
	table   Joiner
	columns []Column
}

func NewInsertInto(table Joiner, columns ...Column) InsertInto {
	return InsertInto{
		table:   table,
		columns: columns,
	}
}

func (i InsertInto) node() generator.Node {
	cs := clauseNodes(i)
	ns := make([]generator.Node, len(cs))
	for j, c := range cs {
		ns[j] = c.myNode()
	}
	return generator.NewNodes(ns...)
}

func (i InsertInto) myNode() generator.Node {
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

func (i InsertInto) previous() Clause {
	return nil
}

func (i InsertInto) Values(values ...interface{}) Values {
	v := NewValues(values...)
	v.prevClause = i
	return v
}

func (i InsertInto) DefaultValues() DefaultValues {
	v := NewDefaultValues()
	v.prev = i
	return v
}
