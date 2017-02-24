package statement

import (
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

func Link(n1, n2 Nodeizer) {
	switch t1 := n1.(type) {
	case Childer:
		switch t2 := n2.(type) {
		case Parenter:
			t1.setChild(t2)
			t2.setParent(t1)
		}
	case Nexter:
		switch t2 := n2.(type) {
		case Prever:
			t1.setNext(t2)
			t2.setPrev(t1)
		}
	}
}

type Nodeizer interface {
	nodeize() (tokenizer.Tokenizer, []interface{})
}

type Childer interface {
	Nodeizer
	child() Parenter
	setChild(Parenter)
	separator() token.Tokens
}

type Parent struct {
	c Parenter
}

func (p *Parent) child() Parenter {
	return p.c
}

func (p *Parent) setChild(c Parenter) {
	p.c = c
}

type Parenter interface {
	Nodeizer
	parent() Childer
	setParent(Childer)
}

type Child struct {
	p Childer
}

func (c *Child) parent() Childer {
	return c.p
}

func (c *Child) setParent(p Childer) {
	c.p = p
}

type Prever interface {
	Nodeizer
	prev() Nexter
	setPrev(Nexter)
}

type Next struct {
	p Nexter
}

func (n *Next) prev() Nexter {
	return n.p
}

func (n *Next) setPrev(p Nexter) {
	n.p = p
}

type Nexter interface {
	Nodeizer
	next() Prever
	setNext(Prever)
}

type Prev struct {
	n Prever
}

func (p *Prev) next() Prever {
	return p.n
}

func (p *Prev) setNext(n Prever) {
	p.n = n
}

func Traverse(t Nodeizer) Nodeizer {
	{
		for p := t; p != nil; {
			t = p
			prever, ok := p.(Prever)
			if !ok {
				break
			}
			p = prever.prev()
		}
	}
	{
		for p := t; p != nil; {
			t = p
			parenter, ok := p.(Parenter)
			if !ok {
				break
			}
			p = parenter.parent()
		}
	}
	return t
}

func Nodize(n Nodeizer) (tokenizer.Tokenizers, []interface{}) {
	if n == nil {
		return nil, nil
	}

	var tokenizers tokenizer.Tokenizers
	values := []interface{}{}

	if childer, ok := n.(Childer); ok {
		if child := childer.child(); child != nil {
			t1, vals1 := n.nodeize()
			t2, vals2 := Nodize(child)
			first, _ := t1.FirstLine()
			_, last := t1.LastLine()
			t12 := tokenizer.
				NewContainer(first).
				SetMiddle(t2.Prefix(childer.separator()...)).
				SetLast(last)

			tokenizers = tokenizer.NewTokenizers(t12)
			values = append(append(values, vals1...), vals2...)
		}
	}

	if tokenizers == nil {
		t, vals := n.nodeize()
		tokenizers = tokenizer.NewTokenizers(t)
		values = vals
	}

	if nexter, ok := n.(Nexter); ok {
		if next := nexter.next(); next != nil {
			ts, vals := Nodize(next)
			tokenizers = append(tokenizers, ts...)
			values = append(values, vals...)
		}
	}

	return tokenizers, values
}
