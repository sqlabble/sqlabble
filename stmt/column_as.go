package stmt

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type ColumnAs struct {
	column ColOrSub
	Alias  string
}

func NewColumnAs(alias string) ColumnAs {
	return ColumnAs{
		Alias: alias,
	}
}

func (c ColumnAs) nodeize() (tokenizer.Tokenizer, []interface{}) {
	if c.column == nil {
		return tokenizer.NewLine(
			token.Word(keyword.As),
		).Append(
			token.QuoteStart,
			token.Word(c.Alias),
			token.QuoteEnd,
		), nil
	}

	t1, v1 := c.column.nodeize()
	return tokenizer.ConcatTokenizers(
		t1,
		tokenizer.NewLine(
			token.QuoteStart,
			token.Word(c.Alias),
			token.QuoteEnd,
		),
		tokenizer.NewLine(
			token.Word(keyword.As),
		),
	), v1
}

// isColumnOrColumnAsOrSubquery always returns true.
// This method exists only to implement the interface ColumnOrColumnAsOrSubquery.
// This is a shit of duck typing, but anyway it works.
func (c ColumnAs) isColOrAliasOrSub() bool {
	return true
}

// isColOrAliasOrFuncOrSub always returns true.
// This method exists only to implement the interface ColOrAliasOrFuncOrSub.
// This is a shit of duck typing, but anyway it works.
func (c ColumnAs) isColOrAliasOrFuncOrSub() bool {
	return true
}
