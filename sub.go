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
	return newEq(s, value)
}

func (s sub) NotEq(value interface{}) comparisonOperation {
	return newNotEq(s, value)
}

func (s sub) Gt(value interface{}) comparisonOperation {
	return newGt(s, value)
}

func (s sub) Gte(value interface{}) comparisonOperation {
	return newGte(s, value)
}

func (s sub) Lt(value interface{}) comparisonOperation {
	return newLt(s, value)
}

func (s sub) Lte(value interface{}) comparisonOperation {
	return newLte(s, value)
}

func (s sub) Like(value string) comparisonOperation {
	return newLike(s, value)
}

func (s sub) RegExp(value string) comparisonOperation {
	return newRegExp(s, value)
}

// func (c sub) Between(from, to interface{}) between {
// 	return newBetween(c, from, to)
// }
//
// func (c sub) In(values ...interface{}) containingOperation {
// 	return newIn(c, values...)
// }
//
// func (c sub) NotIn(values ...interface{}) containingOperation {
// 	return newNotIn(c, values...)
// }
//
// func (c sub) IsNull() nullyOperation {
// 	return newIsNull(c)
// }
//
// func (c sub) IsNotNull() nullyOperation {
// 	return newIsNotNull(c)
// }
