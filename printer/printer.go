// Printer is a contrasting package of scanners in parsers.
// This package converts the tokens to text
// according to the setting of the formatter.
package printer

import (
	"github.com/minodisk/sqlabble/node"
	"github.com/minodisk/sqlabble/token"
)

// Print converts the tokens to text according to ctx.
func Print(tokens token.Tokens, ctx node.Context) string {
	if ctx.IsBreaking() {
		return tokens.String(ctx)
	}

	ts := token.Tokens{}
	var prev token.Token
	for i, t := range tokens {
		switch t {
		case token.Indent:
			continue
		case token.LineEnd:
			continue
		case token.LineStart:
			if i == 0 || prev == token.ParenthesesStart {
				continue
			}
			prev = token.Space
			ts = append(ts, prev)
			continue
		case token.ParenthesesEnd:
			if prev == token.Space {
				ts = ts[:len(ts)-1]
			}
		case token.Comma:
			if prev == token.Space {
				ts = ts[:len(ts)-1]
			}
		}
		ts = append(ts, t)
		prev = t
	}

	return ts.String(ctx)
}
