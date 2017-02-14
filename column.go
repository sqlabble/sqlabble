package sqlabble

import (
	"github.com/minodisk/sqlabble/direction"
	"github.com/minodisk/sqlabble/generator"
)

type column struct {
	name string
}

func newColumn(name string) column {
	return column{
		name: name,
	}
}

func (c column) node() generator.Node {
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

func (c column) Eq(value interface{}) comparisonOperation {
	return newEq(c, value)
}

func (c column) NotEq(value interface{}) comparisonOperation {
	return newNotEq(c, value)
}

func (c column) Gt(value interface{}) comparisonOperation {
	return newGt(c, value)
}

func (c column) Gte(value interface{}) comparisonOperation {
	return newGte(c, value)
}

func (c column) Lt(value interface{}) comparisonOperation {
	return newLt(c, value)
}

func (c column) Lte(value interface{}) comparisonOperation {
	return newLte(c, value)
}

func (c column) Like(value string) comparisonOperation {
	return newLike(c, value)
}

func (c column) RegExp(value string) comparisonOperation {
	return newRegExp(c, value)
}

func (c column) Between(from, to interface{}) between {
	return newBetween(c, from, to)
}

func (c column) In(values ...interface{}) containingOperation {
	return newIn(c, values...)
}

func (c column) NotIn(values ...interface{}) containingOperation {
	return newNotIn(c, values...)
}

func (c column) IsNull() nullyOperation {
	return newIsNull(c)
}

func (c column) IsNotNull() nullyOperation {
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
