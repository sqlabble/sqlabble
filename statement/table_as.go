package statement

import (
	"github.com/minodisk/sqlabble/operator"
	"github.com/minodisk/sqlabble/token"
)

type TableAs struct {
	table Table
	alias string
}

func (t TableAs) nodeize() (token.Tokenizer, []interface{}) {
	return t.self()
}

func (t TableAs) self() (token.Tokenizer, []interface{}) {
	tokenizer, values := t.table.nodeize()
	return tokenizer.
			Append(
				token.Space,
				token.Word(operator.As),
				token.Space,
			).
			Append(
				token.Wrap(
					token.Word(t.alias),
					token.Quote,
				)...,
			),
		values
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
