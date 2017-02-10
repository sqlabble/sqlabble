package token

import (
	"fmt"

	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
)

func NewTable(name string) Table {
	return Table{name: name}
}

type Table struct {
	name string
}

func (t Table) Generator() generator.Generator {
	ts := grammar.Tables(t)
	fs := make([]generator.Generator, len(ts))
	for i, t := range ts {
		fs[i] = t.Expression()
	}
	return generator.NewGenerators(fs...)
}

func (t Table) Expression() generator.Expression {
	return generator.NewExpression(
		t.TableName(),
	)
}

func (t Table) TableName() string {
	return t.name
}

func (t Table) As(alias string) TableAs {
	return TableAs{
		table: t,
		alias: alias,
	}
}

func (t Table) Prev() grammar.Table {
	return nil
}

func (j Table) Join(table grammar.Table) grammar.Table {
	nj := NewJoin(table)
	nj.prev = j
	return nj
}

func (j Table) InnerJoin(table grammar.Table) grammar.Table {
	ij := NewInnerJoin(table)
	ij.prev = j
	return ij
}

func (j Table) LeftJoin(table grammar.Table) grammar.Table {
	lj := NewLeftJoin(table)
	lj.prev = j
	return lj
}

func (j Table) RightJoin(table grammar.Table) grammar.Table {
	rj := NewRightJoin(table)
	rj.prev = j
	return rj
}

type TableAs struct {
	table Table
	alias string
}

func (t TableAs) Generator() generator.Generator {
	ts := grammar.Tables(t)
	fs := make([]generator.Generator, len(ts))
	for i, t := range ts {
		fs[i] = t.Expression()
	}
	return generator.NewGenerators(fs...)
}

func (t TableAs) Expression() generator.Expression {
	return generator.NewExpression(
		fmt.Sprintf("%s AS %s", t.TableName(), t.alias),
	)
}

func (t TableAs) TableName() string {
	return t.table.name
}

func (j TableAs) Prev() grammar.Table {
	return nil
}

func (j TableAs) Join(table grammar.Table) grammar.Table {
	nj := NewJoin(table)
	nj.prev = j
	return nj
}

func (j TableAs) InnerJoin(table grammar.Table) grammar.Table {
	ij := NewInnerJoin(table)
	ij.prev = j
	return ij
}

func (j TableAs) LeftJoin(table grammar.Table) grammar.Table {
	lj := NewLeftJoin(table)
	lj.prev = j
	return lj
}

func (j TableAs) RightJoin(table grammar.Table) grammar.Table {
	rj := NewRightJoin(table)
	rj.prev = j
	return rj
}

type Join struct {
	table grammar.Table
	prev  grammar.Table
}

func NewJoin(table grammar.Table) Join {
	return Join{
		table: table,
	}
}

func (t Join) Generator() generator.Generator {
	ts := grammar.Tables(t)
	fs := make([]generator.Generator, len(ts))
	for i, t := range ts {
		fs[i] = t.Expression()
	}
	return generator.NewGenerators(fs...)
}

func (t Join) Expression() generator.Expression {
	return t.table.Expression().Prepend("JOIN ")
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
	return NewOn(j, column1, column2)
}

type InnerJoin struct {
	table grammar.Table
	prev  grammar.Table
}

func NewInnerJoin(table grammar.Table) InnerJoin {
	return InnerJoin{
		table: table,
	}
}

func (t InnerJoin) Generator() generator.Generator {
	ts := grammar.Tables(t)
	fs := make([]generator.Generator, len(ts))
	for i, t := range ts {
		fs[i] = t.Expression()
	}
	return generator.NewGenerators(fs...)
}

func (t InnerJoin) Expression() generator.Expression {
	return t.table.Expression().Prepend("INNER JOIN ")
}

func (j InnerJoin) Prev() grammar.Table {
	return j.prev
}

func (j InnerJoin) Join(table grammar.Table) grammar.Table {
	nj := NewJoin(table)
	nj.prev = j
	return nj
}

func (j InnerJoin) InnerJoin(table grammar.Table) grammar.Table {
	ij := NewInnerJoin(table)
	ij.prev = j
	return ij
}

func (j InnerJoin) LeftJoin(table grammar.Table) grammar.Table {
	lj := NewLeftJoin(table)
	lj.prev = j
	return lj
}

func (j InnerJoin) RightJoin(table grammar.Table) grammar.Table {
	rj := NewRightJoin(table)
	rj.prev = j
	return rj
}

func (j InnerJoin) On(column1, column2 Column) On {
	return NewOn(j, column1, column2)
}

type LeftJoin struct {
	table grammar.Table
	prev  grammar.Table
}

func NewLeftJoin(table grammar.Table) LeftJoin {
	return LeftJoin{
		table: table,
	}
}

func (t LeftJoin) Generator() generator.Generator {
	ts := grammar.Tables(t)
	fs := make([]generator.Generator, len(ts))
	for i, t := range ts {
		fs[i] = t.Expression()
	}
	return generator.NewGenerators(fs...)
}

func (t LeftJoin) Expression() generator.Expression {
	return t.table.Expression().Prepend("LEFT JOIN ")
}

func (j LeftJoin) Prev() grammar.Table {
	return j.prev
}

func (j LeftJoin) Join(table grammar.Table) grammar.Table {
	nj := NewJoin(table)
	nj.prev = j
	return nj
}

func (j LeftJoin) InnerJoin(table grammar.Table) grammar.Table {
	ij := NewInnerJoin(table)
	ij.prev = j
	return ij
}

func (j LeftJoin) LeftJoin(table grammar.Table) grammar.Table {
	lj := NewLeftJoin(table)
	lj.prev = j
	return lj
}

func (j LeftJoin) RightJoin(table grammar.Table) grammar.Table {
	rj := NewRightJoin(table)
	rj.prev = j
	return rj
}

func (j LeftJoin) On(column1, column2 Column) On {
	return NewOn(j, column1, column2)
}

type RightJoin struct {
	table grammar.Table
	prev  grammar.Table
}

func NewRightJoin(table grammar.Table) RightJoin {
	return RightJoin{
		table: table,
	}
}

func (t RightJoin) Generator() generator.Generator {
	ts := grammar.Tables(t)
	fs := make([]generator.Generator, len(ts))
	for i, t := range ts {
		fs[i] = t.Expression()
	}
	return generator.NewGenerators(fs...)
}

func (t RightJoin) Expression() generator.Expression {
	return t.table.Expression().Prepend("RIGHT JOIN ")
}

func (j RightJoin) Prev() grammar.Table {
	return j.prev
}

func (j RightJoin) Join(table grammar.Table) grammar.Table {
	nj := NewJoin(table)
	nj.prev = j
	return nj
}

func (j RightJoin) InnerJoin(table grammar.Table) grammar.Table {
	ij := NewInnerJoin(table)
	ij.prev = j
	return ij
}

func (j RightJoin) LeftJoin(table grammar.Table) grammar.Table {
	lj := NewLeftJoin(table)
	lj.prev = j
	return lj
}

func (j RightJoin) RightJoin(table grammar.Table) grammar.Table {
	rj := NewRightJoin(table)
	rj.prev = j
	return rj
}

func (j RightJoin) On(column1, column2 Column) On {
	return NewOn(j, column1, column2)
}

type On struct {
	join             grammar.Table
	column1, column2 Column
}

func NewOn(join grammar.Table, column1, column2 Column) On {
	return On{
		join:    join,
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
	return t.join.Expression().
		Append(generator.NewExpression("ON")).
		Append(t.column1.Expression()).
		Append(generator.NewExpression("=")).
		Append(t.column2.Expression())
}

func (o On) Prev() grammar.Table {
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
