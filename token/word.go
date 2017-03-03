package token

import "fmt"

// Tokens that don't change behavior according to Format setting.
var (
	Space       = NewWord(" ")
	Placeholder = NewWord("?")
	Period      = NewWord(".")
	Comma       = NewWord(",")
	LParen      = NewWord("(")
	RParen      = NewWord(")")
	FuncLParen  = NewWord("(")
	FuncRParen  = NewWord(")")
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
