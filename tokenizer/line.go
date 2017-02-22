package tokenizer

import "github.com/minodisk/sqlabble/token"

var (
	EmptyLine = Line{}
)

func ParamsToLine(values ...interface{}) (Line, []interface{}) {
	if len(values) == 0 {
		return EmptyLine, nil
	}
	return NewLine(token.PlaceholderTokens(len(values))...), values
}

type Line struct {
	tokens []token.Token
}

func NewLine(tokens ...token.Token) Line {
	return Line{
		tokens: tokens,
	}
}

func (l Line) Tokenize(depth int) token.Tokens {
	if len(l.tokens) == 0 {
		return token.Tokens{}
	}
	return token.Flatten(
		[]token.Token{token.LineStart},
		token.Repeat(token.Indent, depth),
		l.tokens,
		[]token.Token{token.LineEnd},
	)
}

func (l Line) Prepend(tokens ...token.Token) Tokenizer {
	l.tokens = append(tokens, l.tokens...)
	return l
}

func (l Line) Append(tokens ...token.Token) Tokenizer {
	l.tokens = append(l.tokens, tokens...)
	return l
}

func (l Line) FirstLine() (Line, Tokenizer) {
	return l, nil
}

func (l Line) LastLine() (Tokenizer, Line) {
	return nil, l
}

func (l Line) P(tokens ...token.Token) Line {
	l.tokens = append(tokens, l.tokens...)
	return l
}

func (l Line) A(tokens ...token.Token) Line {
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
	if t1 == nil {
		first, t2 := t2.FirstLine()
		return NewTokenizers(
			first.Prepend(sep.tokens...),
			t2,
		)
	}
	if t2 == nil {
		t1, last := t1.LastLine()
		return NewTokenizers(
			t1,
			last.Append(sep.tokens...),
		)
	}
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

type Lines []Line

func NewLines(lines ...Line) Lines {
	return Lines(lines)
}

func (b Lines) Tokenize(depth int) token.Tokens {
	tokens := []token.Token{}
	for _, line := range b {
		tokens = append(tokens, line.Tokenize(depth)...)
	}
	return tokens
}

func (b Lines) Prepend(tokens ...token.Token) Tokenizer {
	b[0] = b[0].Prepend(tokens...).(Line)
	return b
}

func (b Lines) Append(tokens ...token.Token) Tokenizer {
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

func (b Lines) Prefix(tokens ...token.Token) Lines {
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
