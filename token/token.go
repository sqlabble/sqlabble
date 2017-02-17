package token

import (
	"fmt"

	"github.com/minodisk/sqlabble/node"
)

type TokensNode struct {
	tokens Tokens
	values []interface{}
}

func NewTokensNode(tokens Tokens, values ...interface{}) TokensNode {
	if len(values) == 0 {
		values = []interface{}{}
	}
	return TokensNode{
		tokens: tokens,
		values: values,
	}
}

func (n TokensNode) ToSQL(ctx node.Context) (string, []interface{}) {
	sql := n.tokens.String(ctx)

	h := ctx.CurrentHead()
	ctx = ctx.ClearHead()
	if ctx.IsBreaking() {
		p := ctx.Prefix()
		return fmt.Sprintf("%s%s%s\n", p, h, sql), n.values
	}
	return fmt.Sprintf("%s%s", h, sql), n.values
}

type Token interface {
	String(node.Context) string
}

type Tokens []Token

func (ts Tokens) String(ctx node.Context) string {
	sql := ""
	for _, t := range ts {
		sql += t.String(ctx)
	}
	return sql
}

func (ts Tokens) Append(tokens ...Token) Tokens {
	if len(ts) == 0 {
		return tokens
	}
	return append(append(ts, Space), tokens...)
}

func (ts Tokens) Add(tokens Tokens) Tokens {
	return append(append(ts, Space), tokens...)
}

type quote string

func (q quote) String(ctx node.Context) string {
	return ctx.Quote
}

type Word string

func (w Word) String(ctx node.Context) string {
	return string(w)
}

func (w Word) Append(tokens ...Token) Tokens {
	return append(append(Tokens{w}, Space), tokens...)
}

func Wrap(inner, outer Token) Tokens {
	return Tokens{outer, inner, outer}
}

var (
	Quote            = new(quote)
	Space            = Word(" ")
	Comma            = Word(",")
	ParenthesesStart = Word("(")
	ParenthesesEnd   = Word(")")
	LineStart        = new(lineStart)
	LineEnd          = Word("\n")
)

type lineStart string

func (l lineStart) String(ctx node.Context) string {
	return ctx.Prefix()
}
