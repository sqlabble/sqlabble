package stmt

import (
	"github.com/sqlabble/sqlabble/keyword"
	"github.com/sqlabble/sqlabble/token"
	"github.com/sqlabble/sqlabble/tokenizer"
)

type Delete struct{}

func NewDelete() Delete {
	return Delete{}
}

func (d Delete) From(t Table) From {
	f := NewFrom(t)
	f.prev = d
	return f
}

func (d Delete) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(d)
}

func (d Delete) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.Delete)),
	), nil
}

func (d Delete) previous() Prever {
	return nil
}
