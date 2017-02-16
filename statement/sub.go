package statement

import "github.com/minodisk/sqlabble/node"

type Subquery struct {
	statement Statement
}

func NewSubquery(statement Statement) Subquery {
	return Subquery{
		statement: statement,
	}
}

func (s Subquery) node() node.Node {
	return node.NewParentheses(s.statement.node())
}

// isColumnOrSubquery always returns true.
// This method exists only to implement the interface columnOrSubquery.
// This is a shit of duck typing, but anyway it works.
func (s Subquery) isColumnOrSubquery() bool {
	return true
}

func (s Subquery) Eq(value interface{}) ComparisonOperation {
	e := NewEq(value)
	e.col = s
	return e
}

func (s Subquery) NotEq(value interface{}) ComparisonOperation {
	n := NewNotEq(value)
	n.col = s
	return n
}

func (s Subquery) Gt(value interface{}) ComparisonOperation {
	g := NewGt(value)
	g.col = s
	return g
}

func (s Subquery) Gte(value interface{}) ComparisonOperation {
	g := NewGte(value)
	g.col = s
	return g
}

func (s Subquery) Lt(value interface{}) ComparisonOperation {
	l := NewLt(value)
	l.col = s
	return l
}

func (s Subquery) Lte(value interface{}) ComparisonOperation {
	l := NewLte(value)
	l.col = s
	return l
}

func (s Subquery) Like(value interface{}) ComparisonOperation {
	l := NewLike(value)
	l.col = s
	return l
}

func (s Subquery) RegExp(value interface{}) ComparisonOperation {
	r := NewRegExp(value)
	r.col = s
	return r
}

func (s Subquery) Between(from, to interface{}) Between {
	b := NewBetween(from, to)
	b.col = s
	return b
}

func (s Subquery) In(values ...interface{}) ContainingOperation {
	i := NewIn(values...)
	i.col = s
	return i
}

func (s Subquery) NotIn(values ...interface{}) ContainingOperation {
	n := NewNotIn(values...)
	n.col = s
	return n
}

func (s Subquery) IsNull() NullOperation {
	i := NewIsNull()
	i.col = s
	return i
}

func (s Subquery) IsNotNull() NullOperation {
	i := NewIsNotNull()
	i.col = s
	return i
}
