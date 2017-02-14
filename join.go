package sqlabble

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type join struct {
	joinType string
	table    joiner
	prev     joiner
}

func newJoin(table joiner) join {
	return join{
		joinType: keyword.Join,
		table:    table,
	}
}

func newInnerJoin(table joiner) join {
	return join{
		joinType: keyword.InnerJoin,
		table:    table,
	}
}

func newLeftJoin(table joiner) join {
	return join{
		joinType: keyword.LeftJoin,
		table:    table,
	}
}

func newRightJoin(table joiner) join {
	return join{
		joinType: keyword.RightJoin,
		table:    table,
	}
}

func (t join) node() generator.Node {
	ts := tableNodes(t)
	es := make([]generator.Node, len(ts))
	for i, t := range ts {
		es[i] = t.expression()
	}
	return generator.NewNodes(es...)
}

func (t join) expression() generator.Expression {
	return generator.NewExpression(t.joinType).
		Append(t.table.expression())
}

func (j join) previous() joiner {
	if j.prev == nil {
		return nil
	}
	return j.prev
}

func (j join) Join(table joiner) joiner {
	nj := newJoin(table)
	nj.prev = j
	return nj
}

func (j join) InnerJoin(table joiner) joiner {
	ij := newInnerJoin(table)
	ij.prev = j
	return ij
}

func (j join) LeftJoin(table joiner) joiner {
	lj := newLeftJoin(table)
	lj.prev = j
	return lj
}

func (j join) RightJoin(table joiner) joiner {
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
