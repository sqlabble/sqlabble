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

func (s sub) Eq(value interface{}) comparisonOperation {
	e := newEq(value)
	e.col = s
	return e
}

func (s sub) NotEq(value interface{}) comparisonOperation {
	n := newNotEq(value)
	n.col = s
	return n
}

func (s sub) Gt(value interface{}) comparisonOperation {
	g := newGt(value)
	g.col = s
	return g
}

func (s sub) Gte(value interface{}) comparisonOperation {
	g := newGte(value)
	g.col = s
	return g
}

func (s sub) Lt(value interface{}) comparisonOperation {
	l := newLt(value)
	l.col = s
	return l
}

func (s sub) Lte(value interface{}) comparisonOperation {
	l := newLte(value)
	l.col = s
	return l
}

func (s sub) Like(value interface{}) comparisonOperation {
	l := newLike(value)
	l.col = s
	return l
}

func (s sub) RegExp(value interface{}) comparisonOperation {
	r := newRegExp(value)
	r.col = s
	return r
}

func (s sub) Between(from, to interface{}) between {
	b := newBetween(from, to)
	b.col = s
	return b
}

func (s sub) In(values ...interface{}) containingOperation {
	i := newIn(values...)
	i.col = s
	return i
}

func (s sub) NotIn(values ...interface{}) containingOperation {
	n := newNotIn(values...)
	n.col = s
	return n
}

func (s sub) IsNull() nullyOperation {
	i := newIsNull()
	i.col = s
	return i
}

func (s sub) IsNotNull() nullyOperation {
	i := newIsNotNull()
	i.col = s
	return i
}
