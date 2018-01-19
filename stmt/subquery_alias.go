package stmt

import (
	"github.com/sqlabble/sqlabble/keyword"
	"github.com/sqlabble/sqlabble/token"
	"github.com/sqlabble/sqlabble/tokenizer"
)

type SubqueryAlias struct {
	subquery Subquery
	Alias    string
}

func NewSubqueryAlias(alias string) SubqueryAlias {
	return SubqueryAlias{
		Alias: alias,
	}
}

func (a SubqueryAlias) Join(table TableOrAlias) Join {
	nj := NewJoin(table)
	nj.prev = a
	return nj
}

func (a SubqueryAlias) InnerJoin(table TableOrAlias) Join {
	ij := NewInnerJoin(table)
	ij.prev = a
	return ij
}

func (a SubqueryAlias) LeftJoin(table TableOrAlias) Join {
	lj := NewLeftJoin(table)
	lj.prev = a
	return lj
}

func (a SubqueryAlias) RightJoin(table TableOrAlias) Join {
	rj := NewRightJoin(table)
	rj.prev = a
	return rj
}

func (a SubqueryAlias) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(a)
}

func (a SubqueryAlias) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
	t1, v1 := a.subquery.nodeize()
	t2 := tokenizer.NewLine(
		token.LQuote,
		token.Word(a.Alias),
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

func (a SubqueryAlias) previous() Prever {
	return nil
}

// isTableOrAlias always returns true.
// This method exists only to implement the interface TableOrAlias.
// This is a shit of duck typing, but anyway it works.
func (a SubqueryAlias) isTableOrAlias() bool {
	return true
}

// isJoinerOrAlias always returns true.
// This method exists only to implement the interface TableOrAliasOrJoiner.
// This is a shit of duck typing, but anyway it works.
func (a SubqueryAlias) isJoinerOrAlias() bool {
	return true
}

// isColOrAliasOrFuncOrSub always returns true.
// This method exists only to implement the interface ColOrAliasOrFuncOrSub.
// This is a shit of duck typing, but anyway it works.
func (a SubqueryAlias) isValOrColOrAliasOrFuncOrSubOrFormula() bool {
	return true
}
