package statement

import "github.com/minodisk/sqlabble/token"

func Contract(t1 Childer, t2 Parenter) {
	t1.setChild(t2)
	t2.setParent(t1)
}

func Link(t1 Nexter, t2 Prever) {
	t1.setNext(t2)
	t2.setPrev(t1)
}

type Childer interface {
	Nodeizer
	children() []Parenter
	setChild(Parenter)
}

type Separatorer interface {
	Childer
	separator() token.Tokens
}

type Parent struct {
	c []Parenter
}

func (p *Parent) children() []Parenter {
	return p.c
}

func (p *Parent) setChild(c Parenter) {
	p.c = append(p.c, c)
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
