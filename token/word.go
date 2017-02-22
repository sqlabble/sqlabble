package token

// Tokens that don't change behavior according to Format setting.
const (
	Placeholder      = Word("?")
	Space            = Word(" ")
	Comma            = Word(",")
	ParenthesesStart = Word("(")
	ParenthesesEnd   = Word(")")
)

type Word string

func NewWord(char string) *Word {
	w := Word(char)
	return &w
}

func (w Word) String(format Format) string {
	return string(w)
}

func (w Word) Append(tokens ...Token) []Token {
	return append([]Token{w}, tokens...)
}
