package stmt

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type ColumnAlias struct {
	column ColOrSub
	Alias  string
}

func NewColumnAlias(alias string) ColumnAlias {
	return ColumnAlias{
		Alias: alias,
	}
}

func (c ColumnAlias) nodeize() (tokenizer.Tokenizer, []interface{}) {
	if c.column == nil {
		return tokenizer.NewLine(
			token.Word(keyword.As),
		).Append(
			token.LQuote,
			token.Word(c.Alias),
			token.RQuote,
		), nil
	}
	if c.Alias == "" {
		return c.column.nodeize()
	}

	t1, v1 := c.column.nodeize()
	return tokenizer.ConcatTokenizers(
		t1,
		tokenizer.NewLine(
			token.LQuote,
			token.Word(c.Alias),
			token.RQuote,
		),
		tokenizer.NewLine(
			token.Word(keyword.As),
		),
	), v1
}

func (c ColumnAlias) nodeizeAlias() (tokenizer.Tokenizer, []interface{}) {
	return tokenizer.NewLine(
		token.LQuote,
		token.Word(c.Alias),
		token.RQuote,
	), nil
}

// isColOrAlias always returns true.
// This method exists only to implement the interface ColOrAlias.
// This is a shit of duck typing, but anyway it works.
func (c ColumnAlias) isColOrAlias() bool {
	return true
}

// isColumnOrColumnAsOrSubquery always returns true.
// This method exists only to implement the interface ColumnOrColumnAsOrSubquery.
// This is a shit of duck typing, but anyway it works.
func (c ColumnAlias) isColOrAliasOrSub() bool {
	return true
}

// isColOrAliasOrFuncOrSub always returns true.
// This method exists only to implement the interface ColOrAliasOrFuncOrSub.
// This is a shit of duck typing, but anyway it works.
func (c ColumnAlias) isColOrAliasOrFuncOrSub() bool {
	return true
}
