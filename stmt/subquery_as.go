package stmt

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type SubqueryAs struct {
	subquery Subquery
	Alias    string
}

func NewSubqueryAs(alias string) SubqueryAs {
	return SubqueryAs{
		Alias: alias,
	}
}

func (a SubqueryAs) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t1, v1 := a.subquery.nodeize()
	t2 := tokenizer.NewLine(token.Word(a.Alias))
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
func (a SubqueryAs) isTableOrAlias() bool {
	return true
}

// isTableOrAliasOrJoiner always returns true.
// This method exists only to implement the interface TableOrAliasOrJoiner.
// This is a shit of duck typing, but anyway it works.
func (a SubqueryAs) isTableOrAliasOrJoiner() bool {
	return true
}
