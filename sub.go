package sqlabble

import "github.com/minodisk/sqlabble/generator"

type sub struct {
	stmt Statement
}

func newSub(stmt Statement) sub {
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

func (c sub) Eq(value interface{}) comparisonOperation {
	e := newEq(value)
	e.col = c
	return e
}

func (c sub) NotEq(value interface{}) comparisonOperation {
	n := newNotEq(value)
	n.col = c
	return n
}

func (c sub) Gt(value interface{}) comparisonOperation {
	g := newGt(value)
	g.col = c
	return g
}

func (c sub) Gte(value interface{}) comparisonOperation {
	g := newGte(value)
	g.col = c
	return g
}

func (c sub) Lt(value interface{}) comparisonOperation {
	l := newLt(value)
	l.col = c
	return l
}

func (c sub) Lte(value interface{}) comparisonOperation {
	l := newLte(value)
	l.col = c
	return l
}

func (c sub) Like(value interface{}) comparisonOperation {
	l := newLike(value)
	l.col = c
	return l
}

func (c sub) RegExp(value interface{}) comparisonOperation {
	r := newRegExp(value)
	r.col = c
	return r
}

func (c sub) Between(from, to interface{}) between {
	b := newBetween(from, to)
	b.col = c
	return b
}

func (c sub) In(values ...interface{}) containingOperation {
	i := newIn(values...)
	i.col = c
	return i
}

func (c sub) NotIn(values ...interface{}) containingOperation {
	n := newNotIn(values...)
	n.col = c
	return n
}

func (c sub) IsNull() nullyOperation {
	i := newIsNull()
	i.col = c
	return i
}

func (c sub) IsNotNull() nullyOperation {
	i := newIsNotNull()
	i.col = c
	return i
}
