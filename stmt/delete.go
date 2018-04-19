package stmt

import (
	"github.com/sqlabble/sqlabble/keyword"
	"github.com/sqlabble/sqlabble/token"
	"github.com/sqlabble/sqlabble/tokenizer"
)

type Delete struct {
	tables []Table
}

func NewDelete(tables ...Table) Delete {
	return Delete{
		tables: tables,
	}
}

func (d Delete) From(t JoinerOrAlias) From {
	f := NewFrom(t)
	f.prev = d
	return f
}

func (d Delete) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(d)
}

func (d Delete) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
	tokenizers := make(tokenizer.Tokenizers, len(d.tables))
	values := []interface{}{}
	for i, t := range d.tables {
		var vals []interface{}
		tokenizers[i], vals = t.nodeize()
		values = append(values, vals...)
	}

	return tokenizers.Prepend(
		token.Word(keyword.Delete),
	), values
}

func (d Delete) previous() Prever {
	return nil
}
