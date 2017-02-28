package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Join struct {
	Prev
	Next
	joinType string
	table    TableOrAlias
}

func NewJoin(table TableOrAlias) *Join {
	return &Join{
		joinType: keyword.Join,
		table:    table,
	}
}

func NewInnerJoin(table TableOrAlias) *Join {
	return &Join{
		joinType: keyword.InnerJoin,
		table:    table,
	}
}

func NewLeftJoin(table TableOrAlias) *Join {
	return &Join{
		joinType: keyword.LeftJoin,
		table:    table,
	}
}

func NewRightJoin(table TableOrAlias) *Join {
	return &Join{
		joinType: keyword.RightJoin,
		table:    table,
	}
}

func (j *Join) Join(table TableOrAlias) *Join {
	nj := NewJoin(table)
	Link(j, nj)
	return nj
}

func (j *Join) InnerJoin(table TableOrAlias) *Join {
	nj := NewInnerJoin(table)
	Link(j, nj)
	return nj
}

func (j *Join) LeftJoin(table TableOrAlias) *Join {
	nj := NewLeftJoin(table)
	Link(j, nj)
	return nj
}

func (j *Join) RightJoin(table TableOrAlias) *Join {
	nj := NewRightJoin(table)
	Link(j, nj)
	return nj
}

func (j *Join) On(column1, column2 Column) *On {
	o := NewOn(column1, column2)
	Link(j, o)
	return o
}

func (j *Join) Using(col Column) *Using {
	u := NewUsing(col)
	Link(j, u)
	return u
}

func (j *Join) nodeize() (tokenizer.Tokenizer, []interface{}) {
	if j.table == nil {
		return nil, nil
	}
	t, v := j.table.nodeize()
	return t.Prepend(
		token.Word(j.joinType),
	), v
}
