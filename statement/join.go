package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type join struct {
	joinType string
	table    joiner
	prev     joiner
}

func NewJoin(table joiner) join {
	return join{
		joinType: keyword.Join,
		table:    table,
	}
}

func NewInnerJoin(table joiner) join {
	return join{
		joinType: keyword.InnerJoin,
		table:    table,
	}
}

func NewLeftJoin(table joiner) join {
	return join{
		joinType: keyword.LeftJoin,
		table:    table,
	}
}

func NewRightJoin(table joiner) join {
	return join{
		joinType: keyword.RightJoin,
		table:    table,
	}
}

func (j join) node() generator.Node {
	ts := tableNodes(j)
	es := make([]generator.Node, len(ts))
	for i, t := range ts {
		es[i] = t.expression()
	}
	return generator.NewNodes(es...)
}

func (j join) expression() generator.Expression {
	return generator.NewExpression(j.joinType).
		Append(j.table.expression())
}

func (j join) previous() joiner {
	if j.prev == nil {
		return nil
	}
	return j.prev
}

func (j join) Join(table joiner) joiner {
	nj := NewJoin(table)
	nj.prev = j
	return nj
}

func (j join) InnerJoin(table joiner) joiner {
	ij := NewInnerJoin(table)
	ij.prev = j
	return ij
}

func (j join) LeftJoin(table joiner) joiner {
	lj := NewLeftJoin(table)
	lj.prev = j
	return lj
}

func (j join) RightJoin(table joiner) joiner {
	rj := NewRightJoin(table)
	rj.prev = j
	return rj
}

func (j join) On(column1, column2 column) on {
	o := NewOn(column1, column2)
	o.join = j
	return o
}

func (j join) Using(col column) using {
	o := NewUsing(col)
	o.join = j
	return o
}
