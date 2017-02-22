package tokenizer

import "github.com/minodisk/sqlabble/token"

type Container struct {
	first  Line
	middle Tokenizer
	last   Line
}

func WrapParenthesesLines(lines ...Line) Container {
	return NewContainer(
		NewLine(token.ParenthesesStart),
	).SetMiddle(
		NewLines(lines...),
	).SetLast(
		NewLine(token.ParenthesesEnd),
	)
}

func NewContainer(first Line) Container {
	return Container{
		first: first,
	}
}

func NewParentheses(middle Tokenizer) Container {
	return Container{
		first:  NewLine(token.ParenthesesStart),
		middle: middle,
		last:   NewLine(token.ParenthesesEnd),
	}
}

func (c Container) Middle() Tokenizer {
	return c.middle
}

func (c Container) SetMiddle(middle Tokenizer) Container {
	c.middle = middle
	return c
}

func (c Container) SetLast(line Line) Container {
	c.last = line
	return c
}

func (c Container) Tokenize(depth int) token.Tokens {
	var tokens token.Tokens
	if c.middle != nil {
		tokens = c.middle.Tokenize(depth + 1)
	}

	return token.Flatten(
		c.first.Tokenize(depth),
		tokens,
		c.last.Tokenize(depth),
	)
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
	return c.first, NewContainer(
		EmptyLine,
	).SetMiddle(c.middle).SetLast(c.last)
}

func (c Container) LastLine() (Tokenizer, Line) {
	return NewContainer(
		c.first,
	).SetMiddle(c.middle), c.last
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
