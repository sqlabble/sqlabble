package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
	"github.com/minodisk/sqlabble/internal/operator"
)

type on struct {
	join             joiner
	column1, column2 column
}

func newOn(column1, column2 column) on {
	return on{
		column1: column1,
		column2: column2,
	}
}

func (o on) node() generator.Node {
	ts := tableNodes(o)
	ns := make([]generator.Node, len(ts))
	for i, t := range ts {
		ns[i] = t.expression()
	}
	return generator.NewNodes(ns...)
}

func (o on) expression() generator.Expression {
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

func (o on) previous() joiner {
	if o.join == nil {
		return nil
	}
	return o.join.previous()
}

func (o on) Join(table joiner) joiner {
	j := newJoin(table)
	j.prev = o
	return j
}

func (o on) InnerJoin(table joiner) joiner {
	ij := newInnerJoin(table)
	ij.prev = o
	return ij
}

func (o on) LeftJoin(table joiner) joiner {
	lj := newLeftJoin(table)
	lj.prev = o
	return lj
}

func (o on) RightJoin(table joiner) joiner {
	rj := newRightJoin(table)
	rj.prev = o
	return rj
}
