package token

import (
	"fmt"

	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar"
	"github.com/minodisk/sqlabble/internal/grammar/operator"
)

type And struct {
	operations []grammar.Operation
}

func NewAnd(cs ...grammar.Operation) And {
	return And{operations: cs}
}

func (a And) Generator() generator.Generator {
	fs := make([]generator.Generator, len(a.operations))
	for i, o := range a.operations {
		fs[i] = o.Generator()
	}
	return generator.NewOperator(a.Operator(), fs...)
}

func (a And) Operator() operator.Operator {
	return operator.And
}

func (a And) Operations() []grammar.Operation {
	return a.operations
}

type Or struct {
	operations []grammar.Operation
}

func NewOr(conditions ...grammar.Operation) Or {
	return Or{operations: conditions}
}

func (a Or) Generator() generator.Generator {
	fs := make([]generator.Generator, len(a.operations))
	for i, o := range a.operations {
		fs[i] = o.Generator()
	}
	return generator.NewOperator(a.Operator(), fs...)
}

func (o Or) Operator() operator.Operator {
	return operator.Or
}

func (o Or) Operations() []grammar.Operation {
	return o.operations
}

type Not struct {
	operation grammar.Operation
}

func NewNot(operation grammar.Operation) Not {
	return Not{operation: operation}
}

func (n Not) Generator() generator.Generator {
	return generator.NewNot(n.operation.Generator())
}

func (n Not) Operator() operator.Operator {
	return operator.Not
}

func (n Not) Operations() []grammar.Operation {
	return []grammar.Operation{n.operation}
}

type Eq struct {
	column Column
	value  interface{}
}

func NewEq(column Column, value interface{}) Eq {
	return Eq{
		column: column,
		value:  value,
	}
}

func (c Eq) Generator() generator.Generator {
	return generator.NewExpression(
		fmt.Sprintf("%s %s ?", c.column.name, c.Operator()),
		c.value,
	)
}

func (c Eq) Operator() operator.Operator {
	return operator.Equal
}

type NotEq struct {
	column Column
	value  interface{}
}

func NewNotEq(column Column, value interface{}) NotEq {
	return NotEq{
		column: column,
		value:  value,
	}
}

func (c NotEq) Generator() generator.Generator {
	return generator.NewExpression(
		fmt.Sprintf("%s %s ?", c.column.name, c.Operator()),
		c.value,
	)
}

func (c NotEq) Operator() operator.Operator {
	return operator.NotEqual
}

type Gt struct {
	column Column
	value  interface{}
}

func NewGt(column Column, value interface{}) Gt {
	return Gt{
		column: column,
		value:  value,
	}
}

func (c Gt) Generator() generator.Generator {
	return generator.NewExpression(
		fmt.Sprintf("%s %s ?", c.column.name, c.Operator()),
		c.value,
	)
}

func (c Gt) Operator() operator.Operator {
	return operator.Gt
}

type Gte struct {
	column Column
	value  interface{}
}

func NewGte(column Column, value interface{}) Gte {
	return Gte{
		column: column,
		value:  value,
	}
}

func (c Gte) Generator() generator.Generator {
	return generator.NewExpression(
		fmt.Sprintf("%s %s ?", c.column.name, c.Operator()),
		c.value,
	)
}

func (c Gte) Operator() operator.Operator {
	return operator.Gte
}

type Lt struct {
	column Column
	value  interface{}
}

func NewLt(column Column, value interface{}) Lt {
	return Lt{
		column: column,
		value:  value,
	}
}

func (c Lt) Generator() generator.Generator {
	return generator.NewExpression(
		fmt.Sprintf("%s %s ?", c.column.name, c.Operator()),
		c.value,
	)
}

func (c Lt) Operator() operator.Operator {
	return operator.Lt
}

type Lte struct {
	column Column
	value  interface{}
}

func NewLte(column Column, value interface{}) Lte {
	return Lte{
		column: column,
		value:  value,
	}
}

func (c Lte) Generator() generator.Generator {
	return generator.NewExpression(
		fmt.Sprintf("%s %s ?", c.column.name, c.Operator()),
		c.value,
	)
}

func (c Lte) Operator() operator.Operator {
	return operator.Lte
}

type Between struct {
	column   Column
	from, to interface{}
}

func NewBetween(column Column, from, to interface{}) Between {
	return Between{
		column: column,
		from:   from,
		to:     to,
	}
}

func (c Between) Generator() generator.Generator {
	return generator.NewExpression(
		fmt.Sprintf("%s %s ? %s ?", c.column.name, c.Operator(), operator.And),
		c.from, c.to,
	)
}

func (c Between) Operator() operator.Operator {
	return operator.Between
}

type In struct {
	column Column
	values []interface{}
}

func NewIn(column Column, values ...interface{}) In {
	return In{
		column: column,
		values: values,
	}
}

func (c In) Generator() generator.Generator {
	return generator.NewExpression(
		fmt.Sprintf(
			"%s %s (%s)",
			c.column.name,
			c.Operator(),
			Placeholders(len(c.values)),
		),
		c.values...,
	)
}

func (c In) Operator() operator.Operator {
	return operator.In
}

type NotIn struct {
	column Column
	values []interface{}
}

func NewNotIn(column Column, values ...interface{}) NotIn {
	return NotIn{
		column: column,
		values: values,
	}
}

func (c NotIn) Generator() generator.Generator {
	return generator.NewExpression(
		fmt.Sprintf(
			"%s %s (%s)",
			c.column.name,
			c.Operator(),
			Placeholders(len(c.values)),
		),
		c.values...,
	)
}

func (c NotIn) Operator() operator.Operator {
	return operator.NotIn
}

type Like struct {
	column Column
	like   string
}

func NewLike(column Column, like string) Like {
	return Like{
		column: column,
		like:   like,
	}
}

func (l Like) Generator() generator.Generator {
	return generator.NewExpression(
		fmt.Sprintf("%s %s ?", l.column.name, l.Operator()),
		l.like,
	)
}

func (l Like) Operator() operator.Operator {
	return operator.Like
}

type RegExp struct {
	column Column
	regexp string
}

func NewRegExp(column Column, regexp string) RegExp {
	return RegExp{
		column: column,
		regexp: regexp,
	}
}

func (l RegExp) Generator() generator.Generator {
	return generator.NewExpression(
		fmt.Sprintf("%s %s ?", l.column.name, l.Operator()),
		l.regexp,
	)
}

func (l RegExp) Operator() operator.Operator {
	return operator.RegExp
}

type IsNull struct {
	column Column
}

func NewIsNull(column Column) IsNull {
	return IsNull{
		column: column,
	}
}

func (l IsNull) Generator() generator.Generator {
	return generator.NewExpression(
		fmt.Sprintf("%s %s", l.column.name, l.Operator()),
	)
}

func (l IsNull) Operator() operator.Operator {
	return operator.IsNull
}

type IsNotNull struct {
	column Column
}

func NewIsNotNull(column Column) IsNotNull {
	return IsNotNull{
		column: column,
	}
}

func (l IsNotNull) Generator() generator.Generator {
	return generator.NewExpression(
		fmt.Sprintf("%s %s", l.column.name, l.Operator()),
	)
}

func (l IsNotNull) Operator() operator.Operator {
	return operator.IsNotNull
}
