package tokenizer

import "github.com/sqlabble/sqlabble/token"

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

func (ts Tokenizers) Prepend(tokens ...token.Token) Tokenizer {
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

func (ts Tokenizers) Append(tokens ...token.Token) Tokenizer {
	last := len(ts) - 1
	ts[last] = ts[last].Append(tokens...)
	return ts
}

func (ts Tokenizers) FirstLine() (Line, Tokenizer) {
	ts = NewTokenizers(ts...)
	if len(ts) == 0 {
		return EmptyLine, nil
	}

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
	ts = NewTokenizers(ts...)
	if len(ts) == 0 {
		return nil, EmptyLine
	}

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

func (ts Tokenizers) Prefix(tokens ...token.Token) Tokenizers {
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

func (ts Tokenizers) Tokenize(depth int) token.Tokens {
	tokens := token.Tokens{}
	for _, t := range ts {
		if t == nil {
			continue
		}
		tokens = append(tokens, t.Tokenize(depth)...)
	}
	return tokens
}
