package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
)

type Delete struct{}

func NewDelete() Delete {
	return Delete{}
}

func (d Delete) nodeize() (token.Tokenizer, []interface{}) {
	clauses := clauseNodes(d)
	cs := make(token.Containers, len(clauses))
	vals := []interface{}{}
	for i, clause := range clauses {
		var vs []interface{}
		cs[i], vs = clause.container()
		vals = append(vals, vs...)
	}
	return cs, vals
}

func (d Delete) container() (token.Container, []interface{}) {
	return token.NewContainer(
		token.NewLine(token.Word(keyword.Delete)),
	), nil
}

func (d Delete) previous() Clause {
	return nil
}

func (d Delete) From(t Table) From {
	f := NewFrom(t)
	f.prev = d
	return f
}
