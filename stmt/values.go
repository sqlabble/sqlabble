package stmt

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type Values struct {
	prev   Prever
	valses []Vals
}

func NewValues(valses ...Vals) Values {
	return Values{
		valses: valses,
	}
}

func (v Values) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(v)
}

func (v Values) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
	tokenizers := make(tokenizer.Tokenizers, len(v.valses))
	values := []interface{}{}
	for i, p := range v.valses {
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

func (v Values) previous() Prever {
	return v.prev
}

type DefaultValues struct {
	prev Prever
}

func NewDefaultValues() DefaultValues {
	return DefaultValues{}
}

func (v DefaultValues) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(v)
}

func (v DefaultValues) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.DefaultValues)),
	), nil
}

func (v DefaultValues) previous() Prever {
	return v.prev
}
