package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type Using struct {
	join   grammar.Table
	column Column
}

func NewUsing(column Column) Using {
	return Using{
		column: column,
	}
}

func (t Using) Generator() generator.Generator {
	ts := grammar.Tables(t)
	fs := make([]generator.Generator, len(ts))
	for i, t := range ts {
		fs[i] = t.Expression()
	}
	return generator.NewGenerators(fs...)
}

func (t Using) Expression() generator.Expression {
	e := generator.NewExpression(keyword.Using).
		Append(t.column.Expression())
	if t.join == nil {
		return e
	}
	return t.join.Expression().
		Append(e)
}

func (o Using) Prev() grammar.Table {
	if o.join == nil {
		return nil
	}
	return o.join.Prev()
}

func (o Using) Join(table grammar.Table) grammar.Table {
	j := NewJoin(table)
	j.prev = o
	return j
}

func (o Using) InnerJoin(table grammar.Table) grammar.Table {
	ij := NewInnerJoin(table)
	ij.prev = o
	return ij
}

func (o Using) LeftJoin(table grammar.Table) grammar.Table {
	lj := NewLeftJoin(table)
	lj.prev = o
	return lj
}

func (o Using) RightJoin(table grammar.Table) grammar.Table {
	rj := NewRightJoin(table)
	rj.prev = o
	return rj
}
