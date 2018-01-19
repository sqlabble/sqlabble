package tokenizer

import "github.com/sqlabble/sqlabble/token"

type Tokenizer interface {
	Tokenize(int) token.Tokens
	Prepend(tokens ...token.Token) Tokenizer
	Append(tokens ...token.Token) Tokenizer
	FirstLine() (Line, Tokenizer)
	LastLine() (Tokenizer, Line)
}
