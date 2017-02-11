package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
)

type Table struct {
	name string
}

func NewTable(name string) Table {
	return Table{
		name: name,
	}
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

func (t Table) Join(table grammar.Table) grammar.Table {
	nj := NewJoin(table)
	nj.prev = t
	return nj
}

func (t Table) InnerJoin(table grammar.Table) grammar.Table {
	ij := NewInnerJoin(table)
	ij.prev = t
	return ij
}

func (t Table) LeftJoin(table grammar.Table) grammar.Table {
	lj := NewLeftJoin(table)
	lj.prev = t
	return lj
}

func (t Table) RightJoin(table grammar.Table) grammar.Table {
	rj := NewRightJoin(table)
	rj.prev = t
	return rj
}
