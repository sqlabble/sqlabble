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
	joiners := tableNodes(o)
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
	line1, values1 := o.column1.line()
	line2, values2 := o.column2.line()

	line := token.NewLine(
		token.Word(keyword.On),
	).Join(
		line1,
		token.NewLine(token.Word(operator.Eq)),
		line2,
	)

	lines := token.NewLines(line)
	values := append(values1, values2...)
	if o.joiner == nil {
		return lines, values
	}

	lines0, values0 := o.joiner.self()
	return token.NewTokenizers(lines0, lines), append(values0, values...)
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
