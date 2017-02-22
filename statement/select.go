package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
)

type Select struct {
	distinct bool
	columns  []ColumnOrAliasOrSubquery
}

func NewSelect(columns ...ColumnOrAliasOrSubquery) Select {
	return Select{
		distinct: false,
		columns:  columns,
	}
}

func NewSelectDistinct(columns ...ColumnOrAliasOrSubquery) Select {
	return Select{
		distinct: true,
		columns:  columns,
	}
}

func (s Select) From(t TableOrAliasOrJoiner) From {
	f := NewFrom(t)
	f.prev = s
	return f
}

func (s Select) nodeize() (token.Tokenizer, []interface{}) {
	return nodeizeClauses(s)
}

func (s Select) self() (token.Tokenizer, []interface{}) {
	tokenizers := make(token.Tokenizers, len(s.columns))
	values := []interface{}{}
	for i, c := range s.columns {
		var vals []interface{}
		tokenizers[i], vals = c.nodeize()
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
		tokenizers.Prefix(token.Comma, token.Space),
	), values
}

func (s Select) previous() Clause {
	return nil
}
