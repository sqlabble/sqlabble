package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
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

func (c ColumnAs) nodeize() (tokenizer.Tokenizer, []interface{}) {
	if c.column == nil {
		return tokenizer.NewLine(
			token.Word(keyword.As),
			token.Space,
		).Append(
			token.Wrap(
				token.Word(c.alias),
				token.Quote,
			)...,
		), nil
	}

	t1, v1 := c.column.nodeize()
	return tokenizer.ConcatTokenizers(
		t1,
		tokenizer.NewLine(
			token.Wrap(
				token.Word(c.alias),
				token.Quote,
			)...,
		),
		tokenizer.NewLine(
			token.Space,
			token.Word(keyword.As),
			token.Space,
		),
	), v1
}

// isColumnOrColumnAsOrSubquery always returns true.
// This method exists only to implement the interface ColumnOrColumnAsOrSubquery.
// This is a shit of duck typing, but anyway it works.
func (c ColumnAs) isColumnOrAliasOrSubquery() bool {
	return true
}
