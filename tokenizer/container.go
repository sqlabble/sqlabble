package tokenizer

import "github.com/minodisk/sqlabble/token"

type Container struct {
	first  Tokenizer
	middle Tokenizer
	last   Tokenizer
}

func WrapParenthesesLines(lines ...Line) Container {
	return NewContainer(
		NewLine(token.LParen),
	).SetMiddle(
		NewLines(lines...),
	).SetLast(
		NewLine(token.RParen),
	)
}

func NewContainer(first Tokenizer) Container {
	return Container{
		first: first,
	}
}

func NewParentheses(middle Tokenizer) Container {
	return Container{
		first:  NewLine(token.LParen),
		middle: middle,
		last:   NewLine(token.RParen),
	}
}

func (c Container) Middle() Tokenizer {
	return c.middle
}

func (c Container) SetMiddle(middle Tokenizer) Container {
	c.middle = middle
	return c
}

func (c Container) SetLast(line Tokenizer) Container {
	c.last = line
	return c
}

func (c Container) Tokenize(depth int) token.Tokens {
	tokens := token.Tokens{}
	if c.first != nil {
		ts := c.first.Tokenize(depth)
		tokens = append(tokens, ts...)
	}
	if c.middle != nil {
		ts := c.middle.Tokenize(depth + 1)
		tokens = append(tokens, ts...)
	}
	if c.last != nil {
		ts := c.last.Tokenize(depth)
		tokens = append(tokens, ts...)
	}
	return tokens
}

func (c Container) Prepend(tokens ...token.Token) Tokenizer {
	c.first = c.first.Prepend(tokens...).(Line)
	return c
}

func (c Container) Append(tokens ...token.Token) Tokenizer {
	c.last = c.last.Append(tokens...).(Line)
	return c
}

func (c Container) FirstLine() (Line, Tokenizer) {
	f, t := c.first.FirstLine()
	return f, NewContainer(
		t,
	).SetMiddle(c.middle).SetLast(c.last)
}

func (c Container) LastLine() (Tokenizer, Line) {
	t, l := c.last.LastLine()
	return NewContainer(
		c.first,
	).SetMiddle(c.middle).SetLast(t), l
}

type Containers []Container

func (cs Containers) Tokenize(depth int) token.Tokens {
	tokens := token.Tokens{}
	for _, c := range cs {
		tokens = append(tokens, c.Tokenize(depth)...)
	}
	return tokens
}

func (cs Containers) Prepend(tokens ...token.Token) Tokenizer {
	cs[0] = cs[0].Prepend(tokens...).(Container)
	return cs
}

func (cs Containers) Append(tokens ...token.Token) Tokenizer {
	n := len(cs) - 1
	cs[n] = cs[n].Append(tokens...).(Container)
	return cs
}

func (cs Containers) FirstLine() (Line, Tokenizer) {
	line, t := cs[0].FirstLine()
	ts := Tokenizers{t}
	for i := 1; i < len(cs); i++ {
		ts = append(ts, cs[i])
	}
	return line, ts
}

func (cs Containers) LastLine() (Tokenizer, Line) {
	t, line := cs[len(cs)-1].LastLine()
	ts := Tokenizers{}
	for i := 0; i < len(cs)-1; i++ {
		ts = append(ts, cs[i])
	}
	ts = append(ts, t)
	return ts, line
}
