package stmt

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type TableAlias struct {
	table Table
	Alias string
}

func NewTableAlias(alias string) TableAlias {
	return TableAlias{
		Alias: alias,
	}
}

func (t TableAlias) Join(table TableOrAlias) Join {
	nj := NewJoin(table)
	nj.prev = t
	return nj
}

func (t TableAlias) InnerJoin(table TableOrAlias) Join {
	ij := NewInnerJoin(table)
	ij.prev = t
	return ij
}

func (t TableAlias) LeftJoin(table TableOrAlias) Join {
	lj := NewLeftJoin(table)
	lj.prev = t
	return lj
}

func (t TableAlias) RightJoin(table TableOrAlias) Join {
	rj := NewRightJoin(table)
	rj.prev = t
	return rj
}

func (t TableAlias) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(t)
}

func (t TableAlias) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
	t1, v1 := t.table.nodeize()
	t2 := tokenizer.NewLine(
		token.LQuote,
		token.Word(t.Alias),
		token.RQuote,
	)

	return tokenizer.ConcatTokenizers(
		t1,
		t2,
		tokenizer.NewLine(
			token.Word(keyword.As),
		),
	), v1
}

func (t TableAlias) previous() Prever {
	return nil
}

// isTableOrAlias always returns true.
// This method exists only to implement the interface TableOrAlias.
// This is a shit of duck typing, but anyway it works.
func (t TableAlias) isTableOrAlias() bool {
	return true
}

// isTableOrAliasOrJoiner always returns true.
// This method exists only to implement the interface TableOrAliasOrJoiner.
// This is a shit of duck typing, but anyway it works.
func (t TableAlias) isTableOrAliasOrJoiner() bool {
	return true
}
