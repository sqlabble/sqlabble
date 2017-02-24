package statement

import (
	"github.com/minodisk/sqlabble/keyword"
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
)

// Offset skips specified rows before beginning to return rows.
type Offset struct {
	Next
	count int
}

// NewOffset return a new Offset.
func NewOffset(count int) *Offset {
	return &Offset{
		count: count,
	}
}

func (o Offset) nodeize() (tokenizer.Tokenizer, []interface{}) {
	line, values := tokenizer.ParamsToLine(o.count)
	return tokenizer.NewContainer(
		tokenizer.NewLine(token.Word(keyword.Offset)),
	).SetMiddle(
		line,
	), values
}
