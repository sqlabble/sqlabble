package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
)

type On struct {
	join             grammar.Table
	column1, column2 Column
}

func NewOn(column1, column2 Column) On {
	return On{
		column1: column1,
		column2: column2,
	}
}

func (t On) Generator() generator.Generator {
	ts := grammar.Tables(t)
	fs := make([]generator.Generator, len(ts))
	for i, t := range ts {
		fs[i] = t.Expression()
	}
	return generator.NewGenerators(fs...)
}

func (t On) Expression() generator.Expression {
	e := generator.NewExpression("ON").
		Append(t.column1.Expression()).
		Append(generator.NewExpression("=")).
		Append(t.column2.Expression())
	if t.join == nil {
		return e
	}
	return t.join.Expression().
		Append(e)
}

func (o On) Prev() grammar.Table {
	if o.join == nil {
		return nil
	}
	return o.join.Prev()
}

func (o On) Join(table grammar.Table) grammar.Table {
	j := NewJoin(table)
	j.prev = o
	return j
}

func (o On) InnerJoin(table grammar.Table) grammar.Table {
	ij := NewInnerJoin(table)
	ij.prev = o
	return ij
}

func (o On) LeftJoin(table grammar.Table) grammar.Table {
	lj := NewLeftJoin(table)
	lj.prev = o
	return lj
}

func (o On) RightJoin(table grammar.Table) grammar.Table {
	rj := NewRightJoin(table)
	rj.prev = o
	return rj
}
