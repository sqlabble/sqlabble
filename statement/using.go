package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Using struct {
	Prev
	Next
	column Column
}

func NewUsing(column Column) *Using {
	return &Using{
		column: column,
	}
}

func (u *Using) Join(table TableOrAlias) *Join {
	j := NewJoin(table)
	Link(u, j)
	return j
}

func (u *Using) InnerJoin(table TableOrAlias) *Join {
	j := NewInnerJoin(table)
	Link(u, j)
	return j
}

func (u *Using) LeftJoin(table TableOrAlias) *Join {
	j := NewLeftJoin(table)
	Link(u, j)
	return j
}

func (u *Using) RightJoin(table TableOrAlias) *Join {
	j := NewRightJoin(table)
	Link(u, j)
	return j
}

func (u *Using) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t1, v1 := u.column.nodeize()
	return t1.Prepend(token.Word(keyword.Using)), v1
}

// isTableOrAliasOrJoiner always returns true.
// This method exists only to implement the interface TableOrAliasOrJoiner.
// This is a shit of duck typing, but anyway it works.
func (u *Using) isTableOrAliasOrJoiner() bool {
	return true
}
