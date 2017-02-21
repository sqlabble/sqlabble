package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/operator"
	"github.com/minodisk/sqlabble/token"
)

type On struct {
	joiner           Joiner
	column1, column2 Column
}

func NewOn(column1, column2 Column) On {
	return On{
		column1: column1,
		column2: column2,
	}
}

func (o On) nodeize() (token.Tokenizer, []interface{}) {
	joiners := collectJoiners(o)
	ts := make(token.Tokenizers, len(joiners))
	values := []interface{}{}
	for i, j := range joiners {
		var vals []interface{}
		ts[i], vals = j.self()
		values = append(values, vals...)
	}
	return ts, values
}

func (o On) self() (token.Tokenizer, []interface{}) {
	t1, v1 := o.column1.nodeize()
	t2, v2 := o.column2.nodeize()
	t12 := token.ConcatTokenizers(
		t1,
		t2,
		token.NewLine(
			token.Space,
			token.Word(operator.Eq),
			token.Space,
		),
	).Prepend(
		token.Word(keyword.On),
		token.Space,
	)

	v12 := append(v1, v2...)
	if o.joiner == nil {
		return t12, v12
	}

	t0, v0 := o.joiner.self()
	return token.ConcatTokenizers(
		t0,
		t12,
		token.NewLine(
			token.Space,
		),
	), append(v0, v12...)
}

func (o On) previous() Joiner {
	if o.joiner == nil {
		return nil
	}
	return o.joiner.previous()
}

func (o On) Join(table Joiner) Joiner {
	j := NewJoin(table)
	j.prev = o
	return j
}

func (o On) InnerJoin(table Joiner) Joiner {
	ij := NewInnerJoin(table)
	ij.prev = o
	return ij
}

func (o On) LeftJoin(table Joiner) Joiner {
	lj := NewLeftJoin(table)
	lj.prev = o
	return lj
}

func (o On) RightJoin(table Joiner) Joiner {
	rj := NewRightJoin(table)
	rj.prev = o
	return rj
}
