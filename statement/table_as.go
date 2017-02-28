package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type TableAs struct {
	Prev
	Next
	table *Table
	alias string
}

func NewTableAs(alias string) *TableAs {
	return &TableAs{
		alias: alias,
	}
}

func (t *TableAs) Join(table TableOrAlias) *Join {
	j := NewJoin(table)
	Link(t, j)
	return j
}

func (t *TableAs) InnerJoin(table TableOrAlias) *Join {
	j := NewInnerJoin(table)
	Link(t, j)
	return j
}

func (t *TableAs) LeftJoin(table TableOrAlias) *Join {
	j := NewLeftJoin(table)
	Link(t, j)
	return j
}

func (t *TableAs) RightJoin(table TableOrAlias) *Join {
	j := NewRightJoin(table)
	Link(t, j)
	return j
}

func (t *TableAs) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return t.self()
}

func (t *TableAs) self() (tokenizer.Tokenizer, []interface{}) {
	t1, v1 := t.table.nodeize()
	t2 := tokenizer.NewLine(
		token.QuoteStart,
		token.Word(t.alias),
		token.QuoteEnd,
	)

	return tokenizer.ConcatTokenizers(
		t1,
		t2,
		tokenizer.NewLine(
			token.Word(keyword.As),
		),
	), v1
}

// isTableOrAlias always returns true.
// This method exists only to implement the interface TableOrAlias.
// This is a shit of duck typing, but anyway it works.
func (t *TableAs) isTableOrAlias() bool {
	return true
}

// isTableOrAliasOrJoiner always returns true.
// This method exists only to implement the interface TableOrAliasOrJoiner.
// This is a shit of duck typing, but anyway it works.
func (t *TableAs) isTableOrAliasOrJoiner() bool {
	return true
}
