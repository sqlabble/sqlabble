package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
)

type InsertInto struct {
	table   Table
	columns []Column
}

func NewInsertInto(table Table, columns ...Column) InsertInto {
	return InsertInto{
		table:   table,
		columns: columns,
	}
}

func (i InsertInto) nodeize() (token.Tokenizer, []interface{}) {
	clauses := clauseNodes(i)
	cs := make(token.Containers, len(clauses))
	values := []interface{}{}
	for j, c := range clauses {
		var vals []interface{}
		cs[j], vals = c.container()
		values = append(values, vals...)
	}
	return cs, values
}

func (i InsertInto) container() (token.Container, []interface{}) {
	tableTokenizer, values := i.table.nodeize()
	lines := make(token.Lines, len(i.columns))
	for j, c := range i.columns {
		var vals []interface{}
		lines[j], vals = c.line()
		values = append(values, vals...)
	}
	return token.NewContainer(
		token.NewLine(token.Word(keyword.InsertInto)),
	).SetMiddle(
		token.NewTokenizers(
			tableTokenizer,
			token.WrapParenthesesLines(
				lines.Prefix(token.Comma)...,
			),
		),
	), values
}

func (i InsertInto) previous() Clause {
	return nil
}

func (i InsertInto) Values(values ...interface{}) Values {
	v := NewValues(values...)
	v.prevClause = i
	return v
}

func (i InsertInto) DefaultValues() DefaultValues {
	v := NewDefaultValues()
	v.prev = i
	return v
}
