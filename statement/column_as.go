package statement

import (
	"github.com/minodisk/sqlabble/operator"
	"github.com/minodisk/sqlabble/token"
)

type ColumnAs struct {
	column Column
	alias  string
}

func NewColumnAs(alias string) ColumnAs {
	return ColumnAs{
		alias: alias,
	}
}

func (c ColumnAs) nodeize() (token.Tokenizer, []interface{}) {
	return c.line()
}

func (c ColumnAs) line() (token.Line, []interface{}) {
	line, values := c.column.line()
	return line.
		A(
			token.Word(operator.As),
			token.Space,
		).
		A(token.Wrap(
			token.Quote,
			token.Word(c.alias),
		)...), values
}

func (c ColumnAs) values() []interface{} {
	return nil
}

func (c ColumnAs) columnName() string {
	return c.column.columnName()
}
