package statement

import "github.com/minodisk/sqlabble/token"

type Table struct {
	name string
}

func NewTable(name string) Table {
	return Table{
		name: name,
	}
}

func (t Table) nodeize() (token.Tokenizer, []interface{}) {
	tables := collectJoiners(t)
	ts := make(token.Tokenizers, len(tables))
	values := []interface{}{}
	for i, t := range tables {
		var vals []interface{}
		ts[i], vals = t.self()
		values = append(values, vals...)
	}
	return ts, values
}

func (t Table) self() (token.Tokenizer, []interface{}) {
	if t.name == "" {
		return token.EmptyLine, nil
	}
	return token.NewLines(token.NewLine(
		token.Word(t.name),
	)), nil
}

func (t Table) previous() Joiner {
	return nil
}

func (t Table) As(alias string) TableAs {
	return TableAs{
		table: t,
		alias: alias,
	}
}

func (t Table) Join(table Joiner) Joiner {
	nj := NewJoin(table)
	nj.prev = t
	return nj
}

func (t Table) InnerJoin(table Joiner) Joiner {
	ij := NewInnerJoin(table)
	ij.prev = t
	return ij
}

func (t Table) LeftJoin(table Joiner) Joiner {
	lj := NewLeftJoin(table)
	lj.prev = t
	return lj
}

func (t Table) RightJoin(table Joiner) Joiner {
	rj := NewRightJoin(table)
	rj.prev = t
	return rj
}
