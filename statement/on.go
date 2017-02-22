package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/operator"
	"github.com/minodisk/sqlabble/token"
)

type On struct {
	join             Join
	column1, column2 Column
}

func NewOn(column1, column2 Column) On {
	return On{
		column1: column1,
		column2: column2,
	}
}

func (o On) Join(table TableOrAlias) Join {
	j := NewJoin(table)
	j.prev = o
	return j
}

func (o On) InnerJoin(table TableOrAlias) Join {
	j := NewInnerJoin(table)
	j.prev = o
	return j
}

func (o On) LeftJoin(table TableOrAlias) Join {
	j := NewLeftJoin(table)
	j.prev = o
	return j
}

func (o On) RightJoin(table TableOrAlias) Join {
	j := NewRightJoin(table)
	j.prev = o
	return j
}

func (o On) nodeize() (token.Tokenizer, []interface{}) {
	return nodeizeJoiners(o)
}

func (o On) self() (token.Tokenizer, []interface{}) {
	t0, v0 := o.join.self()
	t1, v1 := o.column1.nodeize()
	t2, v2 := o.column2.nodeize()
	return token.ConcatTokenizers(
		t0,
		token.ConcatTokenizers(
			t1,
			t2,
			token.NewLine(
				token.Space,
				token.Word(operator.Eq),
				token.Space,
			),
		),
		token.NewLine(
			token.Space,
			token.Word(keyword.On),
			token.Space,
		),
	), append(append(v0, v1...), v2...)
}

func (o On) previous() Joiner {
	return o.join.previous()
}

// isTableOrAliasOrJoiner always returns true.
// This method exists only to implement the interface TableOrAliasOrJoiner.
// This is a shit of duck typing, but anyway it works.
func (o On) isTableOrAliasOrJoiner() bool {
	return true
}
