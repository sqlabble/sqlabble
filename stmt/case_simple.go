package stmt

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type SimpleCase struct {
	param      ValOrColOrFuncOrSub
	thenOrElse SimpleThenOrElse
}

func NewSimpleCase(param ValOrColOrFuncOrSub, thenOrElse SimpleThenOrElse) SimpleCase {
	return SimpleCase{
		param:      param,
		thenOrElse: thenOrElse,
	}
}

func (c SimpleCase) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t1, v1 := c.param.nodeize()
	t2, v2 := c.thenOrElse.nodeize()
	return tokenizer.NewContainer(
		t1.Prepend(token.Word(keyword.Case)),
	).SetMiddle(
		t2,
	).SetLast(
		tokenizer.NewLine(token.Word(keyword.End)),
	), append(v1, v2...)
}

type SimpleWhen struct {
	param Param
	prev  Prever
}

func NewSimpleWhen(param Param) SimpleWhen {
	return SimpleWhen{
		param: param,
	}
}

func (w SimpleWhen) Then(param ValOrColOrFuncOrSub) SimpleThen {
	t := NewSimpleThen(param)
	t.prev = w
	return t
}

func (w SimpleWhen) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(w)
}

func (w SimpleWhen) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
	t, vals := w.param.nodeize()
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.When)),
	).SetMiddle(
		t,
	), vals
}

func (w SimpleWhen) previous() Prever {
	return w.prev
}

type SimpleThen struct {
	param ValOrColOrFuncOrSub
	prev  Prever
}

func NewSimpleThen(param ValOrColOrFuncOrSub) SimpleThen {
	return SimpleThen{
		param: param,
	}
}

func (t SimpleThen) When(param Param) SimpleWhen {
	w := NewSimpleWhen(param)
	w.prev = t
	return w
}

func (t SimpleThen) Else(param ValOrColOrFuncOrSub) SimpleElse {
	e := NewSimpleElse(param)
	e.prev = t
	return e
}

func (t SimpleThen) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(t)
}

func (t SimpleThen) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
	tk, vals := t.param.nodeize()
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.Then)),
	).SetMiddle(
		tk,
	), vals
}

func (t SimpleThen) previous() Prever {
	return t.prev
}

func (t SimpleThen) isSimpleThenOrElse() bool {
	return true
}

type SimpleElse struct {
	param ValOrColOrFuncOrSub
	prev  Prever
}

func NewSimpleElse(param ValOrColOrFuncOrSub) SimpleElse {
	return SimpleElse{
		param: param,
	}
}

func (e SimpleElse) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return nodeizePrevs(e)
}

func (e SimpleElse) nodeizeSelf() (tokenizer.Tokenizer, []interface{}) {
	t, vals := e.param.nodeize()
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.Else)),
	).SetMiddle(
		t,
	), vals
}

func (e SimpleElse) previous() Prever {
	return e.prev
}

func (e SimpleElse) isSimpleThenOrElse() bool {
	return true
}
