package sqlabble

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type using struct {
	join joiner
	col  column
}

func newUsing(col column) using {
	return using{
		col: col,
	}
}

func (t using) node() generator.Node {
	ts := tableNodes(t)
	us := make([]generator.Node, len(ts))
	for i, t := range ts {
		us[i] = t.expression()
	}
	return generator.NewParallelNodes(us...)
}

func (t using) expression() generator.Expression {
	e := generator.NewExpression(keyword.Using).
		Append(t.col.expression())
	if t.join == nil {
		return e
	}
	return t.join.expression().
		Append(e)
}

func (o using) previous() joiner {
	if o.join == nil {
		return nil
	}
	return o.join.previous()
}

func (o using) Join(table joiner) joiner {
	j := newJoin(table)
	j.prev = o
	return j
}

func (o using) InnerJoin(table joiner) joiner {
	ij := newInnerJoin(table)
	ij.prev = o
	return ij
}

func (o using) LeftJoin(table joiner) joiner {
	lj := newLeftJoin(table)
	lj.prev = o
	return lj
}

func (o using) RightJoin(table joiner) joiner {
	rj := newRightJoin(table)
	rj.prev = o
	return rj
}
