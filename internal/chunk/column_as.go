package chunk

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/operator"
)

type ColumnAs struct {
	column grammar.Column
	alias  string
}

func NewColumnAs(alias string) ColumnAs {
	return ColumnAs{
		alias: alias,
	}
}

func (c ColumnAs) Generator() generator.Generator {
	return c.Expression()
}

func (c ColumnAs) Expression() generator.Expression {
	a := generator.NewExpression(operator.As).
		Append(generator.NewExpression(c.Alias()))
	if c.column == nil {
		return a
	}
	return c.column.Expression().
		Append(a)
}

func (c ColumnAs) ColumnName() string {
	return c.column.ColumnName()
}

func (c ColumnAs) Alias() string {
	return c.alias
}
