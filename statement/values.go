package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
)

type Values struct {
	prev     Clause
	paramses []Params
}

func NewValues(paramses ...Params) Values {
	return Values{
		paramses: paramses,
	}
}

func (v Values) nodeize() (token.Tokenizer, []interface{}) {
	return nodeizeClauses(v)
}

func (v Values) self() (token.Tokenizer, []interface{}) {
	tokenizers := make(token.Tokenizers, len(v.paramses))
	values := []interface{}{}
	for i, p := range v.paramses {
		var vals []interface{}
		tokenizers[i], vals = p.nodeize()
		values = append(values, vals...)
	}

	return token.NewContainer(
		token.NewLine(token.Word(keyword.Values)),
	).SetMiddle(
		token.NewTokenizers(tokenizers...).Prefix(
			token.Comma,
			token.Space,
		),
	), values
}

func (v Values) previous() Clause {
	return v.prev
}

type DefaultValues struct {
	prev Clause
}

func NewDefaultValues() DefaultValues {
	return DefaultValues{}
}

func (v DefaultValues) nodeize() (token.Tokenizer, []interface{}) {
	return nodeizeClauses(v)
}

func (v DefaultValues) self() (token.Tokenizer, []interface{}) {
	return token.NewContainer(
		token.NewLine(token.Word(keyword.DefaultValues)),
	), nil
}

func (v DefaultValues) previous() Clause {
	return v.prev
}
