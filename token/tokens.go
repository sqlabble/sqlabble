package token

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
			tokens = append(tokens, Placeholder, Comma)
			continue
		}
		tokens = append(tokens, Placeholder)
	}
	return tokens
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

func (ts Tokens) Append(tokens ...Token) Tokens {
	return append(ts, tokens...)
}

func (ts Tokens) Sprint(format Format) string {
	sql := ""
	for _, t := range ts {
		sql += t.Sprint(format)
	}
	return sql
}
