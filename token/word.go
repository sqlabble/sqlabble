package token

import "fmt"

// Tokens that don't change behavior according to Format setting.
var (
	Space          = NewWord(" ")
	Placeholder    = NewWord("?")
	Comma          = NewWord(",")
	ParenStart     = NewWord("(")
	ParenEnd       = NewWord(")")
	FuncParenStart = NewWord("(")
	FuncParenEnd   = NewWord(")")
)

type Word string

func NewWord(char string) *Word {
	w := Word(char)
	return &w
}

func (w Word) String(format Format) string {
	return string(w)
}

func (w Word) Debug() string {
	return fmt.Sprintf("Word(%s)", w)
}

func (w Word) Append(tokens ...Token) []Token {
	return append([]Token{w}, tokens...)
}
