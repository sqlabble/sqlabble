package statement

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/operator"
)

type On struct {
	join             Joiner
	column1, column2 Column
}

func NewOn(column1, column2 Column) On {
	return On{
		column1: column1,
		column2: column2,
	}
}

func (o On) node() generator.Node {
	ts := tableNodes(o)
	ns := make([]generator.Node, len(ts))
	for i, t := range ts {
		ns[i] = t.expression()
	}
	return generator.NewNodes(ns...)
}

func (o On) expression() generator.Expression {
	e := generator.NewExpression(keyword.On).
		Append(o.column1.expression()).
		Append(generator.NewExpression(string(operator.Eq))).
		Append(o.column2.expression())
	if o.join == nil {
		return e
	}
	return o.join.expression().
		Append(e)
}

func (o On) previous() Joiner {
	if o.join == nil {
		return nil
	}
	return o.join.previous()
}

func (o On) Join(table Joiner) Joiner {
	j := NewJoin(table)
	j.prev = o
	return j
}

func (o On) InnerJoin(table Joiner) Joiner {
	ij := NewInnerJoin(table)
	ij.prev = o
	return ij
}

func (o On) LeftJoin(table Joiner) Joiner {
	lj := NewLeftJoin(table)
	lj.prev = o
	return lj
}

func (o On) RightJoin(table Joiner) Joiner {
	rj := NewRightJoin(table)
	rj.prev = o
	return rj
}
