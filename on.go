package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
	"github.com/minodisk/sqlabble/internal/operator"
)

type on struct {
	join             tableNode
	column1, column2 column
}

func newOn(column1, column2 column) on {
	return on{
		column1: column1,
		column2: column2,
	}
}

func (o on) generator() generator.Generator {
	ts := tableNodes(o)
	fs := make([]generator.Generator, len(ts))
	for i, t := range ts {
		fs[i] = t.expression()
	}
	return generator.NewGenerators(fs...)
}

func (o on) expression() generator.Expression {
	e := generator.NewExpression(keyword.On).
		Append(o.column1.expression()).
		Append(generator.NewExpression(string(operator.Equal))).
		Append(o.column2.expression())
	if o.join == nil {
		return e
	}
	return o.join.expression().
		Append(e)
}

func (o on) previous() tableNode {
	if o.join == nil {
		return nil
	}
	return o.join.previous()
}

func (o on) Join(table tableNode) tableNode {
	j := newJoin(table)
	j.prev = o
	return j
}

func (o on) InnerJoin(table tableNode) tableNode {
	ij := newInnerJoin(table)
	ij.prev = o
	return ij
}

func (o on) LeftJoin(table tableNode) tableNode {
	lj := newLeftJoin(table)
	lj.prev = o
	return lj
}

func (o on) RightJoin(table tableNode) tableNode {
	rj := newRightJoin(table)
	rj.prev = o
	return rj
}
