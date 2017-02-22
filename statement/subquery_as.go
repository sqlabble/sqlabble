package statement

import (
	"github.com/minodisk/sqlabble/operator"
	"github.com/minodisk/sqlabble/token"
)

type SubqueryAs struct {
	subquery Subquery
	alias    string
}

func NewSubqueryAs(alias string) SubqueryAs {
	return SubqueryAs{
		alias: alias,
	}
}

func (a SubqueryAs) nodeize() (token.Tokenizer, []interface{}) {
	t1, v1 := a.subquery.nodeize()
	t2 := token.NewLine(token.Word(a.alias))
	return token.ConcatTokenizers(
		t1,
		t2,
		token.NewLine(
			token.Space,
			token.Word(operator.As),
			token.Space,
		),
	), v1
}

// isTableOrAlias always returns true.
// This method exists only to implement the interface TableOrAlias.
// This is a shit of duck typing, but anyway it works.
func (a SubqueryAs) isTableOrAlias() bool {
	return true
}
