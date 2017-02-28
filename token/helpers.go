package token

func Repeat(t Token, n int) Tokens {
	ts := make(Tokens, n)
	for i := 0; i < n; i++ {
		ts[i] = t
	}
	return ts
}

func Flatten(tss ...Tokens) Tokens {
	tokens := Tokens{}
	for _, ts := range tss {
		tokens = append(tokens, ts...)
	}
	return tokens
}

func Placeholders(i int) Tokens {
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
