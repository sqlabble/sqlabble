package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/node"
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

func (i InsertInto) node() node.Node {
	cs := clauseNodes(i)
	ns := make([]node.Node, len(cs))
	for j, c := range cs {
		ns[j] = c.myNode()
	}
	return node.NewNodes(ns...)
}

func (i InsertInto) myNode() node.Node {
	es := make([]node.Expression, len(i.columns))
	for j, c := range i.columns {
		es[j] = c.expression()
	}
	return node.NewContainer(
		node.NewExpression(string(keyword.InsertInto)),
		node.JoinExpressions(
			i.table.expression(),
			node.ArrayToExpression(es...),
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
