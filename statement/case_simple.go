package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

type SimpleCase struct {
	Parent
	param ValOrColOrFuncOrSub
}

func NewSimpleCase(param ValOrColOrFuncOrSub) *SimpleCase {
	return &SimpleCase{
		param: param,
	}
}

func (c *SimpleCase) When(param Param) *SimpleWhen {
	w := NewSimpleWhen(param)
	Contract(c, w)
	return w
}

func (c *SimpleCase) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t, vals := c.param.nodeize()
	return tokenizer.NewContainer(
		t.Prepend(token.Word(keyword.Case)),
	).SetLast(
		tokenizer.NewLine(token.Word(keyword.End)),
	), vals
}

type SimpleWhen struct {
	Child
	Prev
	Next
	param Param
	paren Statement
}

func NewSimpleWhen(param Param) *SimpleWhen {
	return &SimpleWhen{
		param: param,
	}
}

func (w *SimpleWhen) Then(param ValOrColOrFuncOrSub) *SimpleThen {
	t := NewSimpleThen(param)
	Link(w, t)
	return t
}

func (w *SimpleWhen) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t, vals := w.param.nodeize()
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.When)),
	).SetMiddle(
		t,
	), vals
}

type SimpleThen struct {
	Next
	Prev
	param ValOrColOrFuncOrSub
}

func NewSimpleThen(param ValOrColOrFuncOrSub) *SimpleThen {
	return &SimpleThen{
		param: param,
	}
}

func (t *SimpleThen) When(param Param) *SimpleWhen {
	w := NewSimpleWhen(param)
	Link(t, w)
	return w
}

func (t *SimpleThen) Else(param ValOrColOrFuncOrSub) *Else {
	e := NewElse(param)
	Link(t, e)
	return e
}

func (t SimpleThen) nodeize() (tokenizer.Tokenizer, []interface{}) {
	tk, vals := t.param.nodeize()
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.Then)),
	).SetMiddle(
		tk,
	), vals
}

type Else struct {
	Next
	param ValOrColOrFuncOrSub
	pre   Statement
}

func NewElse(param ValOrColOrFuncOrSub) *Else {
	return &Else{
		param: param,
	}
}

func (e Else) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t, vals := e.param.nodeize()
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.Else)),
	).SetMiddle(
		t,
	), vals
}
