package sqlabble

import (
	"github.com/minodisk/sqlabble/internal/generator"
	"github.com/minodisk/sqlabble/internal/grammar/direction"
)

type column struct {
	name string
}

func newColumn(name string) column {
	return column{
		name: name,
	}
}

func (c column) generator() generator.Generator {
	return c.expression()
}

func (c column) expression() generator.Expression {
	return generator.NewExpression(
		c.name,
	)
}

func (c column) columnName() string {
	return c.name
}

func (c column) As(alias string) columnAs {
	a := newColumnAs(alias)
	a.col = c
	return a
}

func (c column) Define(def string) definition {
	d := newDefinition(def)
	d.col = c
	return d
}

func (c column) Assign(value interface{}) assign {
	return newAssign(c, value)
}

func (c column) Eq(value interface{}) eq {
	return newEq(c, value)
}

func (c column) NotEq(value interface{}) notEq {
	return newNotEq(c, value)
}

func (c column) Gt(value interface{}) gt {
	return newGt(c, value)
}

func (c column) Gte(value interface{}) gte {
	return newGte(c, value)
}

func (c column) Lt(value interface{}) lt {
	return newLt(c, value)
}

func (c column) Lte(value interface{}) lte {
	return newLte(c, value)
}

func (c column) Between(from, to interface{}) between {
	return newBetween(c, from, to)
}

func (c column) In(values ...interface{}) in {
	return newIn(c, values...)
}

func (c column) NotIn(values ...interface{}) notIn {
	return newNotIn(c, values...)
}

func (c column) Like(value string) like {
	return newLike(c, value)
}

func (c column) RegExp(value string) regExp {
	return newRegExp(c, value)
}

func (c column) IsNull() isNull {
	return newIsNull(c)
}

func (c column) IsNotNull() isNotNull {
	return newIsNotNull(c)
}

func (c column) Asc() order {
	return order{
		col: c,
		dir: direction.ASC,
	}
}

func (c column) Desc() order {
	return order{
		col: c,
		dir: direction.DESC,
	}
}
