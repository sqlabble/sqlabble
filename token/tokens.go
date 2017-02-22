package token

import "github.com/minodisk/sqlabble/node"

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
