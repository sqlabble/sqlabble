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

// isColumnOrSubquery always returns true.
// This method exists only to implement the interface columnOrSubquery.
// This is a shit of duck typing, but anyway it works.
func (c column) isColumnOrSubquery() bool {
	return true
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
	e := newEq(value)
	e.col = c
	return e
}

func (c column) NotEq(value interface{}) comparisonOperation {
	n := newNotEq(value)
	n.col = c
	return n
}

func (c column) Gt(value interface{}) comparisonOperation {
	g := newGt(value)
	g.col = c
	return g
}

func (c column) Gte(value interface{}) comparisonOperation {
	g := newGte(value)
	g.col = c
	return g
}

func (c column) Lt(value interface{}) comparisonOperation {
	l := newLt(value)
	l.col = c
	return l
}

func (c column) Lte(value interface{}) comparisonOperation {
	l := newLte(value)
	l.col = c
	return l
}

func (c column) Like(value string) comparisonOperation {
	l := newLike(value)
	l.col = c
	return l
}

func (c column) RegExp(value string) comparisonOperation {
	r := newRegExp(value)
	r.col = c
	return r
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
