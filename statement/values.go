package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
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

func (v Values) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizeClauses(v)
}

func (v Values) self() (tokenizer.Tokenizer, []interface{}) {
	tokenizers := make(tokenizer.Tokenizers, len(v.paramses))
	values := []interface{}{}
	for i, p := range v.paramses {
		var vals []interface{}
		tokenizers[i], vals = p.nodeize()
		values = append(values, vals...)
	}

	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.Values)),
	).SetMiddle(
		tokenizer.NewTokenizers(tokenizers...).Prefix(
			token.Comma,
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

func (v DefaultValues) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizeClauses(v)
}

func (v DefaultValues) self() (tokenizer.Tokenizer, []interface{}) {
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.DefaultValues)),
	), nil
}

func (v DefaultValues) previous() Clause {
	return v.prev
}
