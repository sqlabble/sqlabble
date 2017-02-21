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
	return nodeizeClauses(i)
}

func (i InsertInto) self() (token.Tokenizer, []interface{}) {
	tableTokenizer, values := i.table.nodeize()
	ts := make(token.Tokenizers, len(i.columns))
	for j, c := range i.columns {
		var vals []interface{}
		ts[j], vals = c.nodeize()
		values = append(values, vals...)
	}
	return token.NewContainer(
		token.NewLine(token.Word(keyword.InsertInto)),
	).SetMiddle(
		token.ConcatTokenizers(
			tableTokenizer,
			token.NewParentheses(
				ts.Prefix(
					token.Comma,
					token.Space,
				),
			),
			token.NewLine(token.Space),
		),
	), values
}

func (i InsertInto) previous() Clause {
	return nil
}

func (i InsertInto) Values(paramses ...Params) Values {
	v := NewValues(paramses...)
	v.prev = i
	return v
}

func (i InsertInto) DefaultValues() DefaultValues {
	v := NewDefaultValues()
	v.prev = i
	return v
}
