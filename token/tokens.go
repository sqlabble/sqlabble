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
