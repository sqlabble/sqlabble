package statement

import "github.com/minodisk/sqlabble/generator"

type sub struct {
	stmt Statement
}

func NewSub(stmt Statement) sub {
	return sub{
		stmt: stmt,
	}
}

func (s sub) node() generator.Node {
	return generator.NewParentheses(s.stmt.node())
}

// isColumnOrSubquery always returns true.
// This method exists only to implement the interface columnOrSubquery.
// This is a shit of duck typing, but anyway it works.
func (s sub) isColumnOrSubquery() bool {
	return true
}

func (s sub) Eq(value interface{}) comparisonOperation {
	e := NewEq(value)
	e.col = s
	return e
}

func (s sub) NotEq(value interface{}) comparisonOperation {
	n := NewNotEq(value)
	n.col = s
	return n
}

func (s sub) Gt(value interface{}) comparisonOperation {
	g := NewGt(value)
	g.col = s
	return g
}

func (s sub) Gte(value interface{}) comparisonOperation {
	g := NewGte(value)
	g.col = s
	return g
}

func (s sub) Lt(value interface{}) comparisonOperation {
	l := NewLt(value)
	l.col = s
	return l
}

func (s sub) Lte(value interface{}) comparisonOperation {
	l := NewLte(value)
	l.col = s
	return l
}

func (s sub) Like(value interface{}) comparisonOperation {
	l := NewLike(value)
	l.col = s
	return l
}

func (s sub) RegExp(value interface{}) comparisonOperation {
	r := NewRegExp(value)
	r.col = s
	return r
}

func (s sub) Between(from, to interface{}) between {
	b := NewBetween(from, to)
	b.col = s
	return b
}

func (s sub) In(values ...interface{}) containingOperation {
	i := NewIn(values...)
	i.col = s
	return i
}

func (s sub) NotIn(values ...interface{}) containingOperation {
	n := NewNotIn(values...)
	n.col = s
	return n
}

func (s sub) IsNull() nullyOperation {
	i := NewIsNull()
	i.col = s
	return i
}

func (s sub) IsNotNull() nullyOperation {
	i := NewIsNotNull()
	i.col = s
	return i
}
