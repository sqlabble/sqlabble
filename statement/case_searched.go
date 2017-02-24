package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type SearchedCase struct {
	Parent
}

func NewSearchedCase() *SearchedCase {
	return &SearchedCase{}
}

func (c *SearchedCase) When(condition ComparisonOrLogicalOperation) *SearchedWhen {
	w := NewSearchedWhen(condition)
	Link(c, w)
	return w
}

func (c *SearchedCase) separator() token.Tokens {
	return nil
}

func (c *SearchedCase) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.Case)),
	).SetLast(
		tokenizer.NewLine(token.Word(keyword.End)),
	), nil
}

type SearchedWhen struct {
	Child
	Prev
	Next
	condition ComparisonOrLogicalOperation
}

func NewSearchedWhen(condition ComparisonOrLogicalOperation) *SearchedWhen {
	return &SearchedWhen{
		condition: condition,
	}
}

func (w *SearchedWhen) Then(param ValOrColOrFuncOrSub) *SearchedThen {
	t := NewSearchedThen(param)
	Link(w, t)
	return t
}

func (w *SearchedWhen) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t, vals := w.condition.nodeize()
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.When)),
	).SetMiddle(
		t,
	), vals
}

type SearchedThen struct {
	Prev
	Next
	param ValOrColOrFuncOrSub
}

func NewSearchedThen(param ValOrColOrFuncOrSub) *SearchedThen {
	return &SearchedThen{
		param: param,
	}
}

func (t *SearchedThen) When(condition ComparisonOrLogicalOperation) *SearchedWhen {
	w := NewSearchedWhen(condition)
	Link(t, w)
	return w
}

func (t *SearchedThen) Else(param ValOrColOrFuncOrSub) *Else {
	e := NewElse(param)
	Link(t, e)
	return e
}

func (t *SearchedThen) nodeize() (tokenizer.Tokenizer, []interface{}) {
	tk, vals := t.param.nodeize()
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.Then)),
	).SetMiddle(
		tk,
	), vals
}
