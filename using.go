package sqlabble

import (
	"github.com/minodisk/sqlabble/generator"
	"github.com/minodisk/sqlabble/keyword"
)

type using struct {
	join tableOrTableAs
	col  column
}

func newUsing(col column) using {
	return using{
		col: col,
	}
}

func (t using) node() generator.Node {
	ts := tableNodes(t)
	fs := make([]generator.Node, len(ts))
	for i, t := range ts {
		fs[i] = t.expression()
	}
	return generator.NewNodes(fs...)
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

func (o using) previous() tableOrTableAs {
	if o.join == nil {
		return nil
	}
	return o.join.previous()
}

func (o using) Join(table tableOrTableAs) tableOrTableAs {
	j := newJoin(table)
	j.prev = o
	return j
}

func (o using) InnerJoin(table tableOrTableAs) tableOrTableAs {
	ij := newInnerJoin(table)
	ij.prev = o
	return ij
}

func (o using) LeftJoin(table tableOrTableAs) tableOrTableAs {
	lj := newLeftJoin(table)
	lj.prev = o
	return lj
}

func (o using) RightJoin(table tableOrTableAs) tableOrTableAs {
	rj := newRightJoin(table)
	rj.prev = o
	return rj
}
