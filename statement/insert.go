package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
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

func (i InsertInto) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizeClauses(i)
}

func (i InsertInto) self() (tokenizer.Tokenizer, []interface{}) {
	tableTokenizer, values := i.table.nodeize()
	ts := make(tokenizer.Tokenizers, len(i.columns))
	for j, c := range i.columns {
		var vals []interface{}
		ts[j], vals = c.nodeize()
		values = append(values, vals...)
	}
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.InsertInto)),
	).SetMiddle(
		tokenizer.ConcatTokenizers(
			tableTokenizer,
			tokenizer.NewParentheses(
				ts.Prefix(
					token.Comma,
				),
			),
			tokenizer.EmptyLine,
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
