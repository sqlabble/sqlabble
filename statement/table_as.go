package statement

import (
	"github.com/minodisk/sqlabble/node"
	"github.com/minodisk/sqlabble/operator"
)

type TableAs struct {
	table Table
	alias string
}

func (t TableAs) node() node.Node {
	ts := tableNodes(t)
	ns := make([]node.Node, len(ts))
	for i, t := range ts {
		ns[i] = t.expression()
	}
	return node.NewNodes(ns...)
}

func (t TableAs) expression() node.Expression {
	return node.JoinExpressions(
		t.table.expression(),
		node.NewExpression(string(operator.As)),
		node.NewExpression(t.alias),
	)
}

func (t TableAs) previous() Joiner {
	return nil
}

func (t TableAs) Join(table Joiner) Joiner {
	nj := NewJoin(table)
	nj.prev = t
	return nj
}

func (t TableAs) InnerJoin(table Joiner) Joiner {
	ij := NewInnerJoin(table)
	ij.prev = t
	return ij
}

func (t TableAs) LeftJoin(table Joiner) Joiner {
	lj := NewLeftJoin(table)
	lj.prev = t
	return lj
}

func (t TableAs) RightJoin(table Joiner) Joiner {
	rj := NewRightJoin(table)
	rj.prev = t
	return rj
}
