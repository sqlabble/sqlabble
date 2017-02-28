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

func (w Word) Sprint(format Format) string {
	return string(w)
}

func (w Word) String() string {
	return fmt.Sprintf("Word(%s)", string(w))
}
