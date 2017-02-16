package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/node"
)

type Using struct {
	joiner Joiner
	column Column
}

func NewUsing(column Column) Using {
	return Using{
		column: column,
	}
}

func (u Using) node() node.Node {
	ts := tableNodes(u)
	us := make([]node.Node, len(ts))
	for i, t := range ts {
		us[i] = t.expression()
	}
	return node.NewNodes(us...)
}

func (u Using) expression() node.Expression {
	e := node.NewExpression(keyword.Using).
		Append(u.column.expression())
	if u.joiner == nil {
		return e
	}
	return u.joiner.expression().
		Append(e)
}

func (u Using) previous() Joiner {
	if u.joiner == nil {
		return nil
	}
	return u.joiner.previous()
}

func (u Using) Join(table Joiner) Joiner {
	j := NewJoin(table)
	j.prev = u
	return j
}

func (u Using) InnerJoin(table Joiner) Joiner {
	ij := NewInnerJoin(table)
	ij.prev = u
	return ij
}

func (u Using) LeftJoin(table Joiner) Joiner {
	lj := NewLeftJoin(table)
	lj.prev = u
	return lj
}

func (u Using) RightJoin(table Joiner) Joiner {
	rj := NewRightJoin(table)
	rj.prev = u
	return rj
}
