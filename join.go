package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type join struct {
	joinType string
	table    tableNode
	prev     tableNode
}

func newJoin(table tableNode) join {
	return join{
		joinType: keyword.Join,
		table:    table,
	}
}

func newInnerJoin(table tableNode) join {
	return join{
		joinType: keyword.InnerJoin,
		table:    table,
	}
}

func newLeftJoin(table tableNode) join {
	return join{
		joinType: keyword.LeftJoin,
		table:    table,
	}
}

func newRightJoin(table tableNode) join {
	return join{
		joinType: keyword.RightJoin,
		table:    table,
	}
}

func (t join) generator() generator.Generator {
	ts := tableNodes(t)
	es := make([]generator.Generator, len(ts))
	for i, t := range ts {
		es[i] = t.expression()
	}
	return generator.NewGenerators(es...)
}

func (t join) expression() generator.Expression {
	return generator.NewExpression(t.joinType).
		Append(t.table.expression())
}

func (j join) previous() tableNode {
	if j.prev == nil {
		return nil
	}
	return j.prev
}

func (j join) Join(table tableNode) tableNode {
	nj := newJoin(table)
	nj.prev = j
	return nj
}

func (j join) InnerJoin(table tableNode) tableNode {
	ij := newInnerJoin(table)
	ij.prev = j
	return ij
}

func (j join) LeftJoin(table tableNode) tableNode {
	lj := newLeftJoin(table)
	lj.prev = j
	return lj
}

func (j join) RightJoin(table tableNode) tableNode {
	rj := newRightJoin(table)
	rj.prev = j
	return rj
}

func (j join) On(column1, column2 column) on {
	o := newOn(column1, column2)
	o.join = j
	return o
}

func (j join) Using(col column) using {
	o := newUsing(col)
	o.join = j
	return o
}
