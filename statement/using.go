package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
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

func (u Using) nodeize() (token.Tokenizer, []interface{}) {
	return nodeizeJoiners(u)
}

func (u Using) self() (token.Tokenizer, []interface{}) {
	tokens := token.NewTokens(
		token.Word(keyword.Using),
		token.Space,
	)
	t2, v2 := u.column.nodeize()
	t2 = t2.Prepend(tokens...)
	if u.joiner == nil {
		return t2, v2
	}

	t1, v1 := u.joiner.nodeize()
	return token.ConcatTokenizers(
		t1,
		t2,
		token.NewLine(token.Space),
	), append(v1, v2...)
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
