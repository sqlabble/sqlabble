package statement

import "github.com/minodisk/sqlabble/tokenizer"

type Subquery struct {
	statement Statement
}

func NewSubquery(statement Statement) Subquery {
	return Subquery{
		statement: statement,
	}
}

func (s Subquery) As(alias string) SubqueryAs {
	a := NewSubqueryAs(alias)
	a.subquery = s
	return a
}

func (s Subquery) Eq(value ValOrSub) ComparisonOperation {
	e := NewEq(value)
	e.column = s
	return e
}

func (s Subquery) NotEq(value ValOrSub) ComparisonOperation {
	n := NewNotEq(value)
	n.column = s
	return n
}

func (s Subquery) Gt(value ValOrSub) ComparisonOperation {
	g := NewGt(value)
	g.column = s
	return g
}

func (s Subquery) Gte(value ValOrSub) ComparisonOperation {
	g := NewGte(value)
	g.column = s
	return g
}

func (s Subquery) Lt(value ValOrSub) ComparisonOperation {
	l := NewLt(value)
	l.column = s
	return l
}

func (s Subquery) Lte(value ValOrSub) ComparisonOperation {
	l := NewLte(value)
	l.column = s
	return l
}

func (s Subquery) Like(value ValOrSub) ComparisonOperation {
	l := NewLike(value)
	l.column = s
	return l
}

func (s Subquery) RegExp(value ValOrSub) ComparisonOperation {
	r := NewRegExp(value)
	r.column = s
	return r
}

func (s Subquery) Between(from, to ValOrSub) Between {
	b := NewBetween(from, to)
	b.column = s
	return b
}

func (s Subquery) In(params ValsOrSub) ContainingOperation {
	i := NewIn(params)
	i.column = s
	return i
}

func (s Subquery) NotIn(params ValsOrSub) ContainingOperation {
	n := NewNotIn(params)
	n.column = s
	return n
}

func (s Subquery) IsNull() NullOperation {
	i := NewIsNull()
	i.column = s
	return i
}

func (s Subquery) IsNotNull() NullOperation {
	i := NewIsNotNull()
	i.column = s
	return i
}

func (s Subquery) nodeize() (tokenizer.Tokenizer, []interface{}) {
	t, values := s.statement.nodeize()
	return tokenizer.NewParentheses(t), values
}

// isValOrSub always returns true.
// This method exists only to implement the interface ValOrSub.
// This is a shit of duck typing, but anyway it works.
func (s Subquery) isValOrSub() bool {
	return true
}

// isValsOrSub always returns true.
// This method exists only to implement the interface isValsOrSub.
// This is a shit of duck typing, but anyway it works.
func (s Subquery) isValsOrSub() bool {
	return true
}

// isValOrFuncOrSub always returns true.
// This method exists only to implement the interface ValOrFuncOrSub.
// This is a shit of duck typing, but anyway it works.
func (s Subquery) isValOrFuncOrSub() bool {
	return true
}

// isValOrColOrFuncOrSub always returns true.
// This method exists only to implement the interface ValOrColOrFuncOrSub.
// This is a shit of duck typing, but anyway it works.
func (s Subquery) isValOrColOrFuncOrSub() bool {
	return true
}

// isColOrSub always returns true.
// This method exists only to implement the interface ColOrSub.
// This is a shit of duck typing, but anyway it works.
func (s Subquery) isColOrSub() bool {
	return true
}

// isColOrAliasOrSub always returns true.
// This method exists only to implement the interface ColOrAliasOrSub.
// This is a shit of duck typing, but anyway it works.
func (s Subquery) isColOrAliasOrSub() bool {
	return true
}
