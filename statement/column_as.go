package statement

import (
	"github.com/minodisk/sqlabble/operator"
	"github.com/minodisk/sqlabble/token"
)

type ColumnAs struct {
	column ColumnOrSubquery
	alias  string
}

func NewColumnAs(alias string) ColumnAs {
	return ColumnAs{
		alias: alias,
	}
}

func (c ColumnAs) nodeize() (token.Tokenizer, []interface{}) {
	if c.column == nil {
		return token.NewLine(
			token.Word(operator.As),
			token.Space,
		).Append(
			token.Wrap(
				token.Word(c.alias),
				token.Quote,
			)...,
		), nil
	}

	t1, v1 := c.column.nodeize()
	return token.ConcatTokenizers(
		t1,
		token.NewLine(
			token.Wrap(
				token.Word(c.alias),
				token.Quote,
			)...,
		),
		token.NewLine(
			token.Space,
			token.Word(operator.As),
			token.Space,
		),
	), v1
}

// isColumnOrColumnAsOrSubquery always returns true.
// This method exists only to implement the interface ColumnOrColumnAsOrSubquery.
// This is a shit of duck typing, but anyway it works.
func (c ColumnAs) isColumnOrColumnAsOrSubquery() bool {
	return true
}
