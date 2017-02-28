package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Join struct {
	prev     Prever
	joinType string
	table    TableOrAlias
}

func NewJoin(table TableOrAlias) Join {
	return Join{
		joinType: keyword.Join,
		table:    table,
	}
}

func NewInnerJoin(table TableOrAlias) Join {
	return Join{
		joinType: keyword.InnerJoin,
		table:    table,
	}
}

func NewLeftJoin(table TableOrAlias) Join {
	return Join{
		joinType: keyword.LeftJoin,
		table:    table,
	}
}

func NewRightJoin(table TableOrAlias) Join {
	return Join{
		joinType: keyword.RightJoin,
		table:    table,
	}
}

func (j Join) Join(table TableOrAlias) Join {
	j1 := NewJoin(table)
	j1.prev = j
	return j1
}

func (j Join) InnerJoin(table TableOrAlias) Join {
	j1 := NewInnerJoin(table)
	j1.prev = j
	return j1
}

func (j Join) LeftJoin(table TableOrAlias) Join {
	j1 := NewLeftJoin(table)
	j1.prev = j
	return j1
}

func (j Join) RightJoin(table TableOrAlias) Join {
	j1 := NewRightJoin(table)
	j1.prev = j
	return j1
}

func (j Join) On(column1, column2 Column) On {
	o := NewOn(column1, column2)
	o.join = j
	return o
}

func (j Join) Using(col Column) Using {
	o := NewUsing(col)
	o.join = j
	return o
}

func (j Join) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(j)
}

func (j Join) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
	if j.table == nil {
		return nil, nil
	}
	t, v := j.table.nodeize()
	return t.Prepend(
		token.Word(j.joinType),
	), v
}

func (j Join) previous() Prever {
	if j.prev == nil {
		return nil
	}
	return j.prev
}
