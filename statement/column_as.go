package statement

import (
	"github.com/minodisk/sqlabble/node"
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

func (c ColumnAs) node() node.Node {
	return token.NewTokensNode(
		c.tokenize(),
	)
}

func (c ColumnAs) tokenize() token.Tokens {
	return c.column.tokenize().
		Append(token.Word(operator.As)).
		Add(token.Wrap(
			token.Word(c.alias),
			token.Quote,
		))
}

func (c ColumnAs) columnName() string {
	return c.column.columnName()
}
