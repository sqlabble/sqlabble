package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/node"
	"github.com/minodisk/sqlabble/token"
)

type Join struct {
	joinType string
	joiner   Joiner
	prev     Joiner
}

func NewJoin(table Joiner) Join {
	return Join{
		joinType: keyword.Join,
		joiner:   table,
	}
}

func NewInnerJoin(table Joiner) Join {
	return Join{
		joinType: keyword.InnerJoin,
		joiner:   table,
	}
}

func NewLeftJoin(table Joiner) Join {
	return Join{
		joinType: keyword.LeftJoin,
		joiner:   table,
	}
}

func NewRightJoin(table Joiner) Join {
	return Join{
		joinType: keyword.RightJoin,
		joiner:   table,
	}
}

func (j Join) node() node.Node {
	ts := tableNodes(j)
	ns := make([]node.Node, len(ts))
	for i, t := range ts {
		ns[i] = token.NewTokensNode(t.tokenize())
	}
	return node.NewNodes(ns...)
}

func (j Join) tokenize() token.Tokens {
	return token.Word(j.joinType).
		Append(j.joiner.tokenize()...)
}

func (j Join) expression() node.Expression {
	return node.NewExpression(j.joinType).
		Append(j.joiner.expression())
}

func (j Join) previous() Joiner {
	if j.prev == nil {
		return nil
	}
	return j.prev
}

func (j Join) Join(table Joiner) Joiner {
	nj := NewJoin(table)
	nj.prev = j
	return nj
}

func (j Join) InnerJoin(table Joiner) Joiner {
	ij := NewInnerJoin(table)
	ij.prev = j
	return ij
}

func (j Join) LeftJoin(table Joiner) Joiner {
	lj := NewLeftJoin(table)
	lj.prev = j
	return lj
}

func (j Join) RightJoin(table Joiner) Joiner {
	rj := NewRightJoin(table)
	rj.prev = j
	return rj
}

func (j Join) On(column1, column2 Column) On {
	o := NewOn(column1, column2)
	o.joiner = j
	return o
}

func (j Join) Using(col Column) Using {
	o := NewUsing(col)
	o.joiner = j
	return o
}
