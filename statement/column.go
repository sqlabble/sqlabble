package statement

import (
	"github.com/minodisk/sqlabble/direction"
	"github.com/minodisk/sqlabble/generator"
)

type column struct {
	name string
}

func NewColumn(name string) column {
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
	a := NewColumnAs(alias)
	a.col = c
	return a
}

func (c column) Define(def string) definition {
	d := NewDefinition(def)
	d.col = c
	return d
}

func (c column) Assign(value interface{}) assign {
	return NewAssign(c, value)
}

func (c column) Eq(value interface{}) comparisonOperation {
	e := NewEq(value)
	e.col = c
	return e
}

func (c column) NotEq(value interface{}) comparisonOperation {
	n := NewNotEq(value)
	n.col = c
	return n
}

func (c column) Gt(value interface{}) comparisonOperation {
	g := NewGt(value)
	g.col = c
	return g
}

func (c column) Gte(value interface{}) comparisonOperation {
	g := NewGte(value)
	g.col = c
	return g
}

func (c column) Lt(value interface{}) comparisonOperation {
	l := NewLt(value)
	l.col = c
	return l
}

func (c column) Lte(value interface{}) comparisonOperation {
	l := NewLte(value)
	l.col = c
	return l
}

func (c column) Like(value interface{}) comparisonOperation {
	l := NewLike(value)
	l.col = c
	return l
}

func (c column) RegExp(value interface{}) comparisonOperation {
	r := NewRegExp(value)
	r.col = c
	return r
}

func (c column) Between(from, to interface{}) between {
	b := NewBetween(from, to)
	b.col = c
	return b
}

func (c column) In(values ...interface{}) containingOperation {
	i := NewIn(values...)
	i.col = c
	return i
}

func (c column) NotIn(values ...interface{}) containingOperation {
	n := NewNotIn(values...)
	n.col = c
	return n
}

func (c column) IsNull() nullyOperation {
	i := NewIsNull()
	i.col = c
	return i
}

func (c column) IsNotNull() nullyOperation {
	i := NewIsNotNull()
	i.col = c
	return i
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
