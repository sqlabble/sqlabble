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

func (u using) node() generator.Node {
	ts := tableNodes(u)
	us := make([]generator.Node, len(ts))
	for i, t := range ts {
		us[i] = t.expression()
	}
	return generator.NewParallelNodes(us...)
}

func (u using) expression() generator.Expression {
	e := generator.NewExpression(keyword.Using).
		Append(u.col.expression())
	if u.join == nil {
		return e
	}
	return u.join.expression().
		Append(e)
}

func (u using) previous() joiner {
	if u.join == nil {
		return nil
	}
	return u.join.previous()
}

func (u using) Join(table joiner) joiner {
	j := newJoin(table)
	j.prev = u
	return j
}

func (u using) InnerJoin(table joiner) joiner {
	ij := newInnerJoin(table)
	ij.prev = u
	return ij
}

func (u using) LeftJoin(table joiner) joiner {
	lj := newLeftJoin(table)
	lj.prev = u
	return lj
}

func (u using) RightJoin(table joiner) joiner {
	rj := newRightJoin(table)
	rj.prev = u
	return rj
}
