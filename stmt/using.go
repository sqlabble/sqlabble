package stmt

import (
	"github.com/sqlabble/sqlabble/keyword"
	"github.com/sqlabble/sqlabble/token"
	"github.com/sqlabble/sqlabble/tokenizer"
)

type Using struct {
	join   Join
	column Column
}

func NewUsing(column Column) Using {
	return Using{
		column: column,
	}
}

func (u Using) Join(table TableOrAlias) Join {
	j := NewJoin(table)
	j.prev = u
	return j
}

func (u Using) InnerJoin(table TableOrAlias) Join {
	j := NewInnerJoin(table)
	j.prev = u
	return j
}

func (u Using) LeftJoin(table TableOrAlias) Join {
	j := NewLeftJoin(table)
	j.prev = u
	return j
}

func (u Using) RightJoin(table TableOrAlias) Join {
	j := NewRightJoin(table)
	j.prev = u
	return j
}

func (u Using) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(u)
}

func (u Using) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
	t1, v1 := u.join.nodeizeSelf()
	t2, v2 := u.column.nodeize()
	return tokenizer.ConcatTokenizers(
		t1,
		t2,
		tokenizer.NewLine(
			token.Word(keyword.Using),
		),
	), append(v1, v2...)
}

func (u Using) previous() Prever {
	return u.join.previous()
}

// isJoinerOrAlias always returns true.
// This method exists only to implement the interface TableOrAliasOrJoiner.
// This is a shit of duck typing, but anyway it works.
func (u Using) isJoinerOrAlias() bool {
	return true
}
