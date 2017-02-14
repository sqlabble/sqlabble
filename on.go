package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
	"github.com/minodisk/sqlabble/internal/operator"
)

type on struct {
	join             tableOrTableAs
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

func (o on) previous() tableOrTableAs {
	if o.join == nil {
		return nil
	}
	return o.join.previous()
}

func (o on) Join(table tableOrTableAs) tableOrTableAs {
	j := newJoin(table)
	j.prev = o
	return j
}

func (o on) InnerJoin(table tableOrTableAs) tableOrTableAs {
	ij := newInnerJoin(table)
	ij.prev = o
	return ij
}

func (o on) LeftJoin(table tableOrTableAs) tableOrTableAs {
	lj := newLeftJoin(table)
	lj.prev = o
	return lj
}

func (o on) RightJoin(table tableOrTableAs) tableOrTableAs {
	rj := newRightJoin(table)
	rj.prev = o
	return rj
}
