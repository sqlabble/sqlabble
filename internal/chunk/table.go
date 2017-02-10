package chunk

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
