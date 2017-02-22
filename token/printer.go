package token

// Print converts the tokens to text according to format.
// Print behaves in contrast to the scan in
// the implementation of a parser.
func Print(tokens Tokens, format Format) string {
	if format.IsBreaking {
		return tokens.String(format)
	}

	ts := Tokens{}
	var prev Token
	for i, t := range tokens {
		switch t {
		case Indent:
			continue
		case LineEnd:
			continue
		case LineStart:
			if i == 0 || prev == ParenthesesStart {
				continue
			}
			prev = Space
			ts = append(ts, prev)
			continue
		case ParenthesesEnd:
			if prev == Space {
				ts = ts[:len(ts)-1]
			}
		case Comma:
			if prev == Space {
				ts = ts[:len(ts)-1]
			}
		}
		ts = append(ts, t)
		prev = t
	}

	return ts.String(format)
}
