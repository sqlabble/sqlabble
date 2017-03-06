package stmt

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type SearchedCase struct {
	thenOrElse SearchedThenOrElse
}

func NewSearchedCase(thenOrElse SearchedThenOrElse) *SearchedCase {
	return &SearchedCase{
		thenOrElse: thenOrElse,
	}
}

func (c *SearchedCase) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t, v := c.thenOrElse.nodeize()
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.Case)),
	).SetMiddle(
		t,
	).SetLast(
		tokenizer.NewLine(token.Word(keyword.End)),
	), v
}

type SearchedWhen struct {
	condition ComparisonOrLogicalOperation
	prev      Prever
}

func NewSearchedWhen(condition ComparisonOrLogicalOperation) *SearchedWhen {
	return &SearchedWhen{
		condition: condition,
	}
}

func (w *SearchedWhen) Then(param ValOrColOrFuncOrSub) *SearchedThen {
	t := NewSearchedThen(param)
	t.prev = w
	return t
}

func (w *SearchedWhen) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(w)
}

func (w *SearchedWhen) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
	t, vals := w.condition.nodeize()
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.When)),
	).SetMiddle(
		t,
	), vals
}

func (w *SearchedWhen) previous() Prever {
	return w.prev
}

type SearchedThen struct {
	param ValOrColOrFuncOrSub
	prev  Prever
}

func NewSearchedThen(param ValOrColOrFuncOrSub) *SearchedThen {
	return &SearchedThen{
		param: param,
	}
}

func (t *SearchedThen) When(condition ComparisonOrLogicalOperation) *SearchedWhen {
	w := NewSearchedWhen(condition)
	w.prev = t
	return w
}

func (t *SearchedThen) Else(param ValOrColOrFuncOrSub) *SearchedElse {
	e := NewSearchedElse(param)
	e.prev = t
	return e
}

func (t *SearchedThen) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(t)
}

func (t *SearchedThen) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
	tk, vals := t.param.nodeize()
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.Then)),
	).SetMiddle(
		tk,
	), vals
}

func (t *SearchedThen) previous() Prever {
	return t.prev
}

func (t *SearchedThen) isSearchedThenOrElse() bool {
	return true
}

type SearchedElse struct {
	param ValOrColOrFuncOrSub
	prev  Prever
}

func NewSearchedElse(param ValOrColOrFuncOrSub) *SearchedElse {
	return &SearchedElse{
		param: param,
	}
}

func (e *SearchedElse) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(e)
}

func (e *SearchedElse) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
	t, vals := e.param.nodeize()
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.Else)),
	).SetMiddle(
		t,
	), vals
}

func (e *SearchedElse) previous() Prever {
	return e.prev
}

func (e *SearchedElse) isSearchedThenOrElse() bool {
	return true
}
