package token

import "github.com/minodisk/sqlabble/node"

var (
	EmptyLine = Line{}
)

type Tokenizers []Tokenizer

func NewTokenizers(tokenizers ...Tokenizer) Tokenizers {
	ts := Tokenizers{}
	for _, t := range tokenizers {
		if t == nil {
			continue
		}
		ts = append(ts, t)
	}
	return ts
}

func (ts Tokenizers) Prepend(tokens ...Token) Tokenizer {
	if len(ts) == 0 {
		return NewLine(tokens...)
	}
	if ts[0] == nil {
		ts[0] = NewLine(tokens...)
		return ts
	}
	ts[0] = ts[0].Prepend(tokens...)
	return ts
}

func (ts Tokenizers) Append(tokens ...Token) Tokenizer {
	last := len(ts) - 1
	ts[last] = ts[last].Append(tokens...)
	return ts
}

func (ts Tokenizers) FirstLine() (Line, Tokenizer) {
	line, t := ts[0].FirstLine()
	if t == nil {
		ts = ts[1:]
		if len(ts) == 0 {
			return line, nil
		}
		return line, ts
	}
	ts[0] = t
	return line, ts
}

func (ts Tokenizers) LastLine() (Tokenizer, Line) {
	n := len(ts) - 1
	t, line := ts[n].LastLine()
	if t == nil {
		ts = ts[:n]
		if len(ts) == 0 {
			return nil, line
		}
		return ts, line
	}
	ts[n] = t
	return ts, line
}

func (ts Tokenizers) Prefix(tokens ...Token) Tokenizers {
	for i, t := range ts {
		if i == 0 {
			continue
		}
		ts[i] = t.Prepend(tokens...)
	}
	return ts
}

func (ts Tokenizers) Join(sep Line) Tokenizer {
	var tokenizer Tokenizer
	for i := 0; i < len(ts); i++ {
		t := ts[i]
		if i == 0 {
			tokenizer = t
			continue
		}
		t1, last := tokenizer.LastLine()
		first, t2 := t.FirstLine()
		line := last.Join(sep, first)
		tokenizer = NewTokenizers(t1, line, t2)
	}
	return tokenizer
}

func (ts Tokenizers) Tokenize(depth int) Tokens {
	tokens := Tokens{}
	for _, t := range ts {
		if t == nil {
			continue
		}
		tokens = append(tokens, t.Tokenize(depth)...)
	}
	return tokens
}

type Tokenizer interface {
	Tokenize(int) Tokens
	Prepend(tokens ...Token) Tokenizer
	Append(tokens ...Token) Tokenizer
	FirstLine() (Line, Tokenizer)
	LastLine() (Tokenizer, Line)
}

func PlaceholderTokens(i int) Tokens {
	tokens := Tokens{}
	for ; i > 0; i-- {
		if i > 1 {
			tokens = append(tokens, Placeholder, Comma, Space)
			continue
		}
		tokens = append(tokens, Placeholder)
	}
	return tokens
}

func ParamsToLine(values ...interface{}) (Line, []interface{}) {
	if len(values) == 0 {
		return EmptyLine, nil
	}
	return NewLine(PlaceholderTokens(len(values))...), values
}

type Line struct {
	tokens []Token
}

func NewLine(tokens ...Token) Line {
	return Line{
		tokens: tokens,
	}
}

func (l Line) Tokenize(depth int) Tokens {
	if len(l.tokens) == 0 {
		return Tokens{}
	}
	return Flatten(
		[]Token{LineStart},
		Repeat(Indent, depth),
		l.tokens,
		[]Token{LineEnd},
	)
}

func (l Line) Prepend(tokens ...Token) Tokenizer {
	l.tokens = append(tokens, l.tokens...)
	return l
}

func (l Line) Append(tokens ...Token) Tokenizer {
	l.tokens = append(l.tokens, tokens...)
	return l
}

func (l Line) FirstLine() (Line, Tokenizer) {
	return l, nil
}

func (l Line) LastLine() (Tokenizer, Line) {
	return nil, l
}

func (l Line) P(tokens ...Token) Line {
	l.tokens = append(tokens, l.tokens...)
	return l
}

func (l Line) A(tokens ...Token) Line {
	l.tokens = append(l.tokens, tokens...)
	return l
}

func (l Line) Join(lines ...Line) Line {
	tokens := l.tokens
	for _, l := range lines {
		tokens = append(tokens, l.tokens...)
	}
	return NewLine(tokens...)
}

func ConcatTokenizers(t1, t2 Tokenizer, sep Line) Tokenizers {
	t1, last := t1.LastLine()
	first, t2 := t2.FirstLine()
	ts := NewTokenizers(
		t1,
		last.A(sep.tokens...).A(first.tokens...),
		t2,
	)
	return ts
}

func ConcatLines(lines1, lines2 Lines, seps ...Line) Lines {
	ls1 := lines1[:len(lines1)-1]
	last := lines1[len(lines1)-1]
	first := lines2[0]
	ls2 := lines2[1:]
	return ls1.AppendLines(
		last.Join(seps...).Join(first),
	).AppendLines(
		ls2...,
	)
}

func WrapParenthesesLines(lines ...Line) Container {
	return NewContainer(
		NewLine(ParenthesesStart),
	).SetMiddle(
		NewLines(lines...),
	).SetLast(
		NewLine(ParenthesesEnd),
	)
}

type Lines []Line

func NewLines(lines ...Line) Lines {
	return Lines(lines)
}

func (b Lines) Tokenize(depth int) Tokens {
	tokens := []Token{}
	for _, line := range b {
		tokens = append(tokens, line.Tokenize(depth)...)
	}
	return tokens
}

func (b Lines) Prepend(tokens ...Token) Tokenizer {
	b[0] = b[0].Prepend(tokens...).(Line)
	return b
}

func (b Lines) Append(tokens ...Token) Tokenizer {
	n := len(b)
	if n == 0 {
		return NewLines(NewLine(tokens...))
	}
	n = n - 1
	b[n] = b[n].Append(tokens...).(Line)
	return b
}

func (b Lines) FirstLine() (Line, Tokenizer) {
	if len(b) == 0 {
		return EmptyLine, nil
	}
	return b[0], b[1:]
}

func (b Lines) LastLine() (Tokenizer, Line) {
	if len(b) == 0 {
		return nil, EmptyLine
	}
	n := len(b) - 1
	return b[:n], b[n]
}

func (b Lines) Prefix(tokens ...Token) Lines {
	for i, line := range b {
		if i == 0 {
			continue
		}
		b[i] = line.P(tokens...)
	}
	return b
}

func (b Lines) AppendLines(lines ...Line) Lines {
	b = append(b, lines...)
	return b
}

type Container struct {
	first  Line
	middle Tokenizer
	last   Line
}

func NewContainer(first Line) Container {
	return Container{
		first: first,
	}
}

func NewParentheses(middle Tokenizer) Container {
	return Container{
		first:  NewLine(ParenthesesStart),
		middle: middle,
		last:   NewLine(ParenthesesEnd),
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

func (c Container) Tokenize(depth int) Tokens {
	var tokens Tokens
	if c.middle != nil {
		tokens = c.middle.Tokenize(depth + 1)
	}

	return Flatten(
		c.first.Tokenize(depth),
		tokens,
		c.last.Tokenize(depth),
	)
}

func (c Container) Prepend(tokens ...Token) Tokenizer {
	c.first = c.first.Prepend(tokens...).(Line)
	return c
}

func (c Container) Append(tokens ...Token) Tokenizer {
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

func (cs Containers) Tokenize(depth int) Tokens {
	tokens := Tokens{}
	for _, c := range cs {
		tokens = append(tokens, c.Tokenize(depth)...)
	}
	return tokens
}

func (cs Containers) Prepend(tokens ...Token) Tokenizer {
	cs[0] = cs[0].Prepend(tokens...).(Container)
	return cs
}

func (cs Containers) Append(tokens ...Token) Tokenizer {
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

//---------------------------------------------------

type Token interface {
	String(node.Context) string
}

type Tokens []Token

func NewTokens(tokens ...Token) Tokens {
	ts := Tokens{}
	for _, t := range tokens {
		if t != nil {
			ts = append(ts, t)
		}
	}
	return ts
}

func (ts Tokens) Append(tokens ...Token) Tokens {
	return append(ts, tokens...)
}

func (ts Tokens) String(ctx node.Context) string {
	sql := ""
	for _, t := range ts {
		sql += t.String(ctx)
	}
	return sql
}

func Wrap(inner, outer Token) Tokens {
	return Tokens{outer, inner, outer}
}

func Repeat(t Token, n int) Tokens {
	ts := make(Tokens, n)
	for i := 0; i < n; i++ {
		ts[i] = t
	}
	return ts
}

func Flatten(tss ...[]Token) Tokens {
	tokens := Tokens{}
	for _, ts := range tss {
		tokens = append(tokens, ts...)
	}
	return tokens
}

type quote string

func (q quote) String(ctx node.Context) string {
	return ctx.Quote
}

type Word string

func NewWord(char string) *Word {
	w := Word(char)
	return &w
}

func (w Word) String(ctx node.Context) string {
	return string(w)
}

func (w Word) Append(tokens ...Token) []Token {
	return append([]Token{w}, tokens...)
}

var (
	Quote            = *new(quote)
	Placeholder      = Word("?")
	Space            = Word(" ")
	Comma            = Word(",")
	ParenthesesStart = Word("(")
	ParenthesesEnd   = *new(parenthesesEnd)
	LineStart        = *new(lineStart)
	LineEnd          = *new(lineEnd)
	Indent           = *new(indent)
)

type parenthesesEnd string

func (p parenthesesEnd) String(ctx node.Context) string {
	return ")"
}

type lineStart string

func (l lineStart) String(ctx node.Context) string {
	return ctx.Prefix()
}

type lineEnd string

func (l lineEnd) String(ctx node.Context) string {
	return ctx.LineEnd()
}

type indent string

func (i indent) String(ctx node.Context) string {
	return ctx.Indent()
}
