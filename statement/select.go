package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Select struct {
	Prev
	distinct bool
	columns  []ColOrAliasOrFuncOrSub
}

func NewSelect(columns ...ColOrAliasOrFuncOrSub) *Select {
	return &Select{
		distinct: false,
		columns:  columns,
	}
}

func NewSelectDistinct(columns ...ColOrAliasOrFuncOrSub) *Select {
	return &Select{
		distinct: true,
		columns:  columns,
	}
}

func (s *Select) From(t TableOrAliasOrJoiner) *From {
	f := NewFrom(t)
	Link(s, f)
	return f
}

func (s *Select) nodeize() (tokenizer.Tokenizer, []interface{}) {
	tokenizers := make(tokenizer.Tokenizers, len(s.columns))
	values := []interface{}{}
	for i, c := range s.columns {
		var vals []interface{}
		tokenizers[i], vals = c.nodeize()
		values = append(values, vals...)
	}
	tokens := token.NewTokens(token.Word(keyword.Select))
	if s.distinct {
		tokens = tokens.Append(
			token.Word(keyword.Distinct),
		)
	}
	return tokenizer.NewContainer(
		tokenizer.NewLine(tokens...),
	).SetMiddle(
		tokenizers.Prefix(token.Comma),
	), values
}
