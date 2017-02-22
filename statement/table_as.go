package statement

import (
	"github.com/minodisk/sqlabble/operator"
	"github.com/minodisk/sqlabble/token"
)

type TableAs struct {
	table Table
	alias string
}

func (t TableAs) Join(table TableOrAlias) Join {
	nj := NewJoin(table)
	nj.prev = t
	return nj
}

func (t TableAs) InnerJoin(table TableOrAlias) Join {
	ij := NewInnerJoin(table)
	ij.prev = t
	return ij
}

func (t TableAs) LeftJoin(table TableOrAlias) Join {
	lj := NewLeftJoin(table)
	lj.prev = t
	return lj
}

func (t TableAs) RightJoin(table TableOrAlias) Join {
	rj := NewRightJoin(table)
	rj.prev = t
	return rj
}

func (t TableAs) nodeize() (token.Tokenizer, []interface{}) {
	return t.self()
}

func (t TableAs) self() (token.Tokenizer, []interface{}) {
	t1, v1 := t.table.nodeize()
	t2 := token.NewLine(
		token.Wrap(
			token.Word(t.alias),
			token.Quote,
		)...,
	)

	return token.ConcatTokenizers(
		t1,
		t2,
		token.NewLine(
			token.Space,
			token.Word(operator.As),
			token.Space,
		),
	), v1
}

func (t TableAs) previous() Joiner {
	return nil
}

// isTableOrAlias always returns true.
// This method exists only to implement the interface TableOrAlias.
// This is a shit of duck typing, but anyway it works.
func (t TableAs) isTableOrAlias() bool {
	return true
}

// isTableOrAliasOrJoiner always returns true.
// This method exists only to implement the interface TableOrAliasOrJoiner.
// This is a shit of duck typing, but anyway it works.
func (t TableAs) isTableOrAliasOrJoiner() bool {
	return true
}
