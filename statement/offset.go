package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
)

// Offset skips specified rows before beginning to return rows.
type Offset struct {
	prev  Clause
	count int
}

// NewOffset return a new Offset.
func NewOffset(count int) Offset {
	return Offset{
		count: count,
	}
}

func (o Offset) nodeize() (token.Tokenizer, []interface{}) {
	clauses := clauseNodes(o)
	cs := make(token.Containers, len(clauses))
	values := []interface{}{}
	for i, c := range clauses {
		var vals []interface{}
		cs[i], vals = c.container()
		values = append(values, vals...)
	}
	return cs, values
}

func (o Offset) container() (token.Container, []interface{}) {
	line, values := token.ParamsToLine(o.count)
	return token.NewContainer(
		token.NewLine(token.Word(keyword.Offset)),
	).SetMiddle(
		line,
	), values
}

func (o Offset) previous() Clause {
	return o.prev
}
