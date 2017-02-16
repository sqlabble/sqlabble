package statement

import (
	"github.com/minodisk/sqlabble/direction"
	"github.com/minodisk/sqlabble/node"
)

// Column is a statement to indicate a column in a table.
type Column struct {
	name string
}

// NewColumn returns a new Column.
func NewColumn(name string) Column {
	return Column{
		name: name,
	}
}

// As is used to give an alias name to the column.
// Returns a new ColumnAs.
func (c Column) As(alias string) ColumnAs {
	a := NewColumnAs(alias)
	a.column = c
	return a
}

// Define is used to specify a definition for the column.
// This constitutes a part of the table creation.
// Returns a new Definition.
func (c Column) Define(definition string) Definition {
	d := NewDefinition(definition)
	d.column = c
	return d
}

// Assign is used to assign a value to the column.
// This constitutes a part of the record update statement.
// Returns a new Assign.
func (c Column) Assign(value interface{}) Assign {
	return NewAssign(c, value)
}

func (c Column) Eq(value interface{}) ComparisonOperation {
	e := NewEq(value)
	e.col = c
	return e
}

func (c Column) NotEq(value interface{}) ComparisonOperation {
	n := NewNotEq(value)
	n.col = c
	return n
}

func (c Column) Gt(value interface{}) ComparisonOperation {
	g := NewGt(value)
	g.col = c
	return g
}

func (c Column) Gte(value interface{}) ComparisonOperation {
	g := NewGte(value)
	g.col = c
	return g
}

func (c Column) Lt(value interface{}) ComparisonOperation {
	l := NewLt(value)
	l.col = c
	return l
}

func (c Column) Lte(value interface{}) ComparisonOperation {
	l := NewLte(value)
	l.col = c
	return l
}

func (c Column) Like(value interface{}) ComparisonOperation {
	l := NewLike(value)
	l.col = c
	return l
}

func (c Column) RegExp(value interface{}) ComparisonOperation {
	r := NewRegExp(value)
	r.col = c
	return r
}

func (c Column) Between(from, to interface{}) Between {
	b := NewBetween(from, to)
	b.col = c
	return b
}

func (c Column) In(values ...interface{}) ContainingOperation {
	i := NewIn(values...)
	i.col = c
	return i
}

func (c Column) NotIn(values ...interface{}) ContainingOperation {
	n := NewNotIn(values...)
	n.col = c
	return n
}

func (c Column) IsNull() NullOperation {
	i := NewIsNull()
	i.col = c
	return i
}

func (c Column) IsNotNull() NullOperation {
	i := NewIsNotNull()
	i.col = c
	return i
}

func (c Column) Asc() Order {
	return Order{
		column: c,
		dir:    direction.ASC,
	}
}

func (c Column) Desc() Order {
	return Order{
		column: c,
		dir:    direction.DESC,
	}
}

func (c Column) node() node.Node {
	return c.expression()
}

func (c Column) expression() node.Expression {
	return node.NewExpression(
		c.name,
	)
}

func (c Column) columnName() string {
	return c.name
}

// isColumnOrSubquery always returns true.
// This method exists only to implement the interface ColumnOrSubquery.
// This is a shit of duck typing, but anyway it works.
func (c Column) isColumnOrSubquery() bool {
	return true
}
