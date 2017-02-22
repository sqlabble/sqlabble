package statement

import (
	"github.com/minodisk/sqlabble/token"
	"github.com/minodisk/sqlabble/tokenizer"
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

// Assign is used to assign a params to the column.
// This constitutes a part of the record update statement.
// Returns a new Assign.
func (c Column) Assign(param ParamOrSubquery) Assign {
	return NewAssign(c, param)
}

func (c Column) Eq(value ParamOrSubquery) ComparisonOperation {
	e := NewEq(value)
	e.column = c
	return e
}

func (c Column) NotEq(value ParamOrSubquery) ComparisonOperation {
	n := NewNotEq(value)
	n.column = c
	return n
}

func (c Column) Gt(value ParamOrSubquery) ComparisonOperation {
	g := NewGt(value)
	g.column = c
	return g
}

func (c Column) Gte(value ParamOrSubquery) ComparisonOperation {
	g := NewGte(value)
	g.column = c
	return g
}

func (c Column) Lt(value ParamOrSubquery) ComparisonOperation {
	l := NewLt(value)
	l.column = c
	return l
}

func (c Column) Lte(value ParamOrSubquery) ComparisonOperation {
	l := NewLte(value)
	l.column = c
	return l
}

func (c Column) Like(value ParamOrSubquery) ComparisonOperation {
	l := NewLike(value)
	l.column = c
	return l
}

func (c Column) RegExp(value ParamOrSubquery) ComparisonOperation {
	r := NewRegExp(value)
	r.column = c
	return r
}

func (c Column) Between(from, to ParamOrSubquery) Between {
	b := NewBetween(from, to)
	b.column = c
	return b
}

func (c Column) In(params ParamsOrSubquery) ContainingOperation {
	i := NewIn(params)
	i.column = c
	return i
}

func (c Column) NotIn(params ParamsOrSubquery) ContainingOperation {
	n := NewNotIn(params)
	n.column = c
	return n
}

func (c Column) IsNull() NullOperation {
	i := NewIsNull()
	i.column = c
	return i
}

func (c Column) IsNotNull() NullOperation {
	i := NewIsNotNull()
	i.column = c
	return i
}

func (c Column) EqAll(params Subquery) NonScalarOperation {
	n := NewEqAll(params)
	n.column = c
	return n
}

func (c Column) NotEqAll(params Subquery) NonScalarOperation {
	n := NewNotEqAll(params)
	n.column = c
	return n
}

func (c Column) GtAll(params Subquery) NonScalarOperation {
	n := NewGtAll(params)
	n.column = c
	return n
}

func (c Column) GteAll(params Subquery) NonScalarOperation {
	n := NewGteAll(params)
	n.column = c
	return n
}

func (c Column) LtAll(params Subquery) NonScalarOperation {
	n := NewLtAll(params)
	n.column = c
	return n
}

func (c Column) LteAll(params Subquery) NonScalarOperation {
	n := NewLteAll(params)
	n.column = c
	return n
}

func (c Column) EqAny(params Subquery) NonScalarOperation {
	n := NewEqAny(params)
	n.column = c
	return n
}

func (c Column) NotEqAny(params Subquery) NonScalarOperation {
	n := NewNotEqAny(params)
	n.column = c
	return n
}

func (c Column) GtAny(params Subquery) NonScalarOperation {
	n := NewGtAny(params)
	n.column = c
	return n
}

func (c Column) GteAny(params Subquery) NonScalarOperation {
	n := NewGteAny(params)
	n.column = c
	return n
}

func (c Column) LtAny(params Subquery) NonScalarOperation {
	n := NewLtAny(params)
	n.column = c
	return n
}

func (c Column) LteAny(params Subquery) NonScalarOperation {
	n := NewLteAny(params)
	n.column = c
	return n
}

func (c Column) Asc() Order {
	o := NewAsc()
	o.column = c
	return o
}

func (c Column) Desc() Order {
	o := NewDesc()
	o.column = c
	return o
}

func (c Column) nodeize() (tokenizer.Tokenizer, []interface{}) {
	return tokenizer.NewLine(token.Word(c.name)), nil
}

// isColumnOrSubquery always returns true.
// This method exists only to implement the interface ColumnOrSubquery.
// This is a shit of duck typing, but anyway it works.
func (c Column) isColumnOrSubquery() bool {
	return true
}

// isColumnOrColumnAsOrSubquery always returns true.
// This method exists only to implement the interface ColumnOrColumnAsOrSubquery.
// This is a shit of duck typing, but anyway it works.
func (c Column) isColumnOrAliasOrSubquery() bool {
	return true
}
