package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/keyword"
)

type Join struct {
	joinType string
	table    grammar.Table
	prev     grammar.Table
}

func NewJoin(table grammar.Table) Join {
	return Join{
		joinType: keyword.Join,
		table:    table,
	}
}

func NewInnerJoin(table grammar.Table) Join {
	return Join{
		joinType: keyword.InnerJoin,
		table:    table,
	}
}

func NewLeftJoin(table grammar.Table) Join {
	return Join{
		joinType: keyword.LeftJoin,
		table:    table,
	}
}

func NewRightJoin(table grammar.Table) Join {
	return Join{
		joinType: keyword.RightJoin,
		table:    table,
	}
}

func (t Join) Generator() generator.Generator {
	ts := grammar.Tables(t)
	es := make([]generator.Generator, len(ts))
	for i, t := range ts {
		es[i] = t.Expression()
	}
	return generator.NewGenerators(es...)
}

func (t Join) Expression() generator.Expression {
	return generator.NewExpression(t.joinType).
		Append(t.table.Expression())
}

func (j Join) Prev() grammar.Table {
	if j.prev == nil {
		return nil
	}
	return j.prev
}

func (j Join) Join(table grammar.Table) grammar.Table {
	nj := NewJoin(table)
	nj.prev = j
	return nj
}

func (j Join) InnerJoin(table grammar.Table) grammar.Table {
	ij := NewInnerJoin(table)
	ij.prev = j
	return ij
}

func (j Join) LeftJoin(table grammar.Table) grammar.Table {
	lj := NewLeftJoin(table)
	lj.prev = j
	return lj
}

func (j Join) RightJoin(table grammar.Table) grammar.Table {
	rj := NewRightJoin(table)
	rj.prev = j
	return rj
}

func (j Join) On(column1, column2 Column) On {
	o := NewOn(column1, column2)
	o.join = j
	return o
}

func (j Join) Using(column Column) Using {
	o := NewUsing(column)
	o.join = j
	return o
}
