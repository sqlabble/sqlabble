package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
)

type Select struct {
	distinct bool
	columns  []ColumnOrColumnAs
}

func NewSelect(columns ...ColumnOrColumnAs) Select {
	return Select{
		distinct: false,
		columns:  columns,
	}
}

func NewSelectDistinct(columns ...ColumnOrColumnAs) Select {
	return Select{
		distinct: true,
		columns:  columns,
	}
}

func (s Select) nodeize() (token.Tokenizer, []interface{}) {
	clauses := clauseNodes(s)
	cs := make(token.Containers, len(clauses))
	values := []interface{}{}
	for i, c := range clauses {
		var vals []interface{}
		cs[i], vals = c.container()
		values = append(values, vals...)
	}
	if len(values) == 0 {
		values = nil
	}
	return cs, values
}

func (s Select) container() (token.Container, []interface{}) {
	lines := make(token.Lines, len(s.columns))
	values := []interface{}{}
	for i, c := range s.columns {
		var vals []interface{}
		lines[i], vals = c.line()
		values = append(values, vals...)
	}
	tokens := token.NewTokens(token.Word(keyword.Select))
	if s.distinct {
		tokens = tokens.Append(
			token.Space,
			token.Word(keyword.Distinct),
		)
	}
	return token.NewContainer(
		token.NewLine(tokens...),
	).SetMiddle(
		lines.Prefix(token.Comma, token.Space),
	), values
}

func (s Select) previous() Clause {
	return nil
}

func (s Select) From(t Joiner) From {
	f := NewFrom(t)
	f.prev = s
	return f
}
