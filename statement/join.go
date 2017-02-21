package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
)

type Join struct {
	joinType string
	joiner   Joiner
	prev     Joiner
}

func NewJoin(joiner Joiner) Join {
	return Join{
		joinType: keyword.Join,
		joiner:   joiner,
	}
}

func NewInnerJoin(joiner Joiner) Join {
	return Join{
		joinType: keyword.InnerJoin,
		joiner:   joiner,
	}
}

func NewLeftJoin(joiner Joiner) Join {
	return Join{
		joinType: keyword.LeftJoin,
		joiner:   joiner,
	}
}

func NewRightJoin(joiner Joiner) Join {
	return Join{
		joinType: keyword.RightJoin,
		joiner:   joiner,
	}
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

func (j Join) nodeize() (token.Tokenizer, []interface{}) {
	return nodeizeJoiners(j)
}

func (j Join) self() (token.Tokenizer, []interface{}) {
	t, values := j.joiner.nodeize()
	return t.Prepend(
		token.Word(j.joinType),
		token.Space,
	), values
}

func (j Join) previous() Joiner {
	if j.prev == nil {
		return nil
	}
	return j.prev
}
