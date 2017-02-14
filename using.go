package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/keyword"
)

type using struct {
	join tableNode
	col  column
}

func newUsing(col column) using {
	return using{
		col: col,
	}
}

func (t using) generator() generator.Generator {
	ts := tableNodes(t)
	fs := make([]generator.Generator, len(ts))
	for i, t := range ts {
		fs[i] = t.expression()
	}
	return generator.NewGenerators(fs...)
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

func (o using) previous() tableNode {
	if o.join == nil {
		return nil
	}
	return o.join.previous()
}

func (o using) Join(table tableNode) tableNode {
	j := newJoin(table)
	j.prev = o
	return j
}

func (o using) InnerJoin(table tableNode) tableNode {
	ij := newInnerJoin(table)
	ij.prev = o
	return ij
}

func (o using) LeftJoin(table tableNode) tableNode {
	lj := newLeftJoin(table)
	lj.prev = o
	return lj
}

func (o using) RightJoin(table tableNode) tableNode {
	rj := newRightJoin(table)
	rj.prev = o
	return rj
}
