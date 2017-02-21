package statement

import "github.com/minodisk/sqlabble/token"

type Subquery struct {
	statement Statement
}

func NewSubquery(statement Statement) Subquery {
	return Subquery{
		statement: statement,
	}
}

func (s Subquery) Eq(value ParamOrSubquery) ComparisonOperation {
	e := NewEq(value)
	e.column = s
	return e
}

func (s Subquery) NotEq(value ParamOrSubquery) ComparisonOperation {
	n := NewNotEq(value)
	n.column = s
	return n
}

func (s Subquery) Gt(value ParamOrSubquery) ComparisonOperation {
	g := NewGt(value)
	g.column = s
	return g
}

func (s Subquery) Gte(value ParamOrSubquery) ComparisonOperation {
	g := NewGte(value)
	g.column = s
	return g
}

func (s Subquery) Lt(value ParamOrSubquery) ComparisonOperation {
	l := NewLt(value)
	l.column = s
	return l
}

func (s Subquery) Lte(value ParamOrSubquery) ComparisonOperation {
	l := NewLte(value)
	l.column = s
	return l
}

func (s Subquery) Like(value ParamOrSubquery) ComparisonOperation {
	l := NewLike(value)
	l.column = s
	return l
}

func (s Subquery) RegExp(value ParamOrSubquery) ComparisonOperation {
	r := NewRegExp(value)
	r.column = s
	return r
}

func (s Subquery) Between(from, to ParamOrSubquery) Between {
	b := NewBetween(from, to)
	b.column = s
	return b
}

func (s Subquery) In(params ParamsOrSubquery) ContainingOperation {
	i := NewIn(params)
	i.column = s
	return i
}

func (s Subquery) NotIn(params ParamsOrSubquery) ContainingOperation {
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

func (s Subquery) nodeize() (token.Tokenizer, []interface{}) {
	tokenizer, values := s.statement.nodeize()
	return token.NewParentheses(tokenizer), values
}

// isColumnOrSubquery always returns true.
// This method exists only to implement the interface ColumnOrSubquery.
// This is a shit of duck typing, but anyway it works.
func (s Subquery) isColumnOrSubquery() bool {
	return true
}

// isColumnOrColumnAsOrSubquery always returns true.
// This method exists only to implement the interface ColumnOrColumnAsOrSubquery.
// This is a shit of duck typing, but anyway it works.
func (s Subquery) isColumnOrColumnAsOrSubquery() bool {
	return true
}

// isParamOrSubquery always returns true.
// This method exists only to implement the interface ParamOrSubquery.
// This is a shit of duck typing, but anyway it works.
func (s Subquery) isParamOrSubquery() bool {
	return true
}

// isParamsOrSubquery always returns true.
// This method exists only to implement the interface ParamsOrSubquery.
// This is a shit of duck typing, but anyway it works.
func (s Subquery) isParamsOrSubquery() bool {
	return true
}
