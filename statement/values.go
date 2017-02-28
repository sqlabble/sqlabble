package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Values struct {
	Next
	paramses []Params
}

func NewValues(paramses ...Params) *Values {
	return &Values{
		paramses: paramses,
	}
}

func (v *Values) nodeize() (tokenizer.Tokenizer, []interface{}) {
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

type DefaultValues struct {
	Next
}

func NewDefaultValues() *DefaultValues {
	return &DefaultValues{}
}

func (v *DefaultValues) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.DefaultValues)),
	), nil
}
