package token

import (
	"fmt"
	"strings"

	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/direction"
	"github.com/minodisk/sqlabble/internal/grammar/operator"
)

func NewColumn(name string) Column {
	return Column{name: name}
}

type Columns []Column

func (cs Columns) Generator() generator.Generator {
	names := make([]string, len(cs))
	for i, c := range cs {
		names[i] = c.name
	}
	return generator.NewExpression(
		fmt.Sprintf(
			"(%s)",
			strings.Join(names, ", "),
		),
	)
}

type Column struct {
	name string
}

func (c Column) Generator() generator.Generator {
	return c.Expression()
}

func (c Column) Expression() generator.Expression {
	return generator.NewExpression(
		c.name,
	)
}

func (c Column) ColumnName() string {
	return c.name
}

func (c Column) Define(definition string) Definition {
	d := NewDefinition(definition)
	d.column = c
	return d
}

func (c Column) As(alias string) ColumnAs {
	return ColumnAs{
		column: c,
		alias:  alias,
	}
}

func (c Column) Assign(value interface{}) Assign {
	return NewAssign(c, value)
}

func (c Column) Eq(value interface{}) Eq {
	return NewEq(c, value)
}

func (c Column) NotEq(value interface{}) NotEq {
	return NewNotEq(c, value)
}

func (c Column) Gt(value interface{}) Gt {
	return NewGt(c, value)
}

func (c Column) Gte(value interface{}) Gte {
	return NewGte(c, value)
}

func (c Column) Lt(value interface{}) Lt {
	return NewLt(c, value)
}

func (c Column) Lte(value interface{}) Lte {
	return NewLte(c, value)
}

func (c Column) Between(from, to interface{}) Between {
	return NewBetween(c, from, to)
}

func (c Column) In(values ...interface{}) In {
	return NewIn(c, values...)
}

func (c Column) NotIn(values ...interface{}) NotIn {
	return NewNotIn(c, values...)
}

func (c Column) Like(value string) Like {
	return NewLike(c, value)
}

func (c Column) RegExp(value string) RegExp {
	return NewRegExp(c, value)
}

func (c Column) IsNull() IsNull {
	return NewIsNull(c)
}

func (c Column) IsNotNull() IsNotNull {
	return NewIsNotNull(c)
}

func (c Column) Asc() Order {
	return Order{
		column:    c,
		direction: direction.ASC,
	}
}

func (c Column) Desc() Order {
	return Order{
		column:    c,
		direction: direction.DESC,
	}
}

type ColumnAs struct {
	column grammar.Column
	alias  string
}

func (c ColumnAs) Generator() generator.Generator {
	return c.Expression()
}

func (c ColumnAs) Expression() generator.Expression {
	return c.column.Expression().
		Append(generator.NewExpression(operator.As)).
		Append(generator.NewExpression(c.Alias()))
}

func (c ColumnAs) ColumnName() string {
	return c.column.ColumnName()
}

func (c ColumnAs) Alias() string {
	return c.alias
}
