package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type Join struct {
	joinType string
	table    Joiner
	prev     Joiner
}

func NewJoin(table Joiner) Join {
	return Join{
		joinType: keyword.Join,
		table:    table,
	}
}

func NewInnerJoin(table Joiner) Join {
	return Join{
		joinType: keyword.InnerJoin,
		table:    table,
	}
}

func NewLeftJoin(table Joiner) Join {
	return Join{
		joinType: keyword.LeftJoin,
		table:    table,
	}
}

func NewRightJoin(table Joiner) Join {
	return Join{
		joinType: keyword.RightJoin,
		table:    table,
	}
}

func (j Join) node() generator.Node {
	ts := tableNodes(j)
	es := make([]generator.Node, len(ts))
	for i, t := range ts {
		es[i] = t.expression()
	}
	return generator.NewNodes(es...)
}

func (j Join) expression() generator.Expression {
	return generator.NewExpression(j.joinType).
		Append(j.table.expression())
}

func (j Join) previous() Joiner {
	if j.prev == nil {
		return nil
	}
	return j.prev
}

func (j Join) Join(table Joiner) Joiner {
	nj := NewJoin(table)
	nj.prev = j
	return nj
}

func (j Join) InnerJoin(table Joiner) Joiner {
	ij := NewInnerJoin(table)
	ij.prev = j
	return ij
}

func (j Join) LeftJoin(table Joiner) Joiner {
	lj := NewLeftJoin(table)
	lj.prev = j
	return lj
}

func (j Join) RightJoin(table Joiner) Joiner {
	rj := NewRightJoin(table)
	rj.prev = j
	return rj
}

func (j Join) On(column1, column2 Column) On {
	o := NewOn(column1, column2)
	o.join = j
	return o
}

func (j Join) Using(col Column) Using {
	o := NewUsing(col)
	o.joiner = j
	return o
}
