package node

import "fmt"

var (
	// EmptyExpression is nully Expression.
	EmptyExpression = NewExpression("")
)

// Expression is a Node that have a sql string and a slice of values.
// It will be converted to a single line query.
type Expression struct {
	sql    string
	values []interface{}
}

// NewExpression returns a new Expression.
func NewExpression(sql string, values ...interface{}) Expression {
	if len(values) == 0 {
		values = []interface{}{}
	}
	return Expression{
		sql:    sql,
		values: values,
	}
}

// ToSQL returns a query and a slice of values.
func (e Expression) ToSQL(ctx Context) (string, []interface{}) {
	h := ctx.currentHead()
	ctx = ctx.clearHead()
	if ctx.isBreaking() {
		p := ctx.pre()
		return fmt.Sprintf("%s%s%s\n", p, h, e.sql), e.values
	}
	return fmt.Sprintf("%s%s", h, e.sql), e.values
}

// Prepend adds exp to the beginning of e.
func (e Expression) Prepend(exp Expression) Expression {
	if exp.IsEmpty() {
		return e
	}
	if e.IsEmpty() {
		return exp
	}
	e.sql = exp.sql + " " + e.sql
	e.values = append(exp.values, e.values...)
	return e
}

// Append adds exp to the end of e.
func (e Expression) Append(exp Expression) Expression {
	if exp.IsEmpty() {
		return e
	}
	if e.IsEmpty() {
		return exp
	}
	e.sql = e.sql + " " + exp.sql
	e.values = append(e.values, exp.values...)
	return e
}

// WrapSQL adds the specified characters before and after sql.
func (e Expression) WrapSQL(pre, post string) Expression {
	e.sql = pre + e.sql + post
	return e
}

// IsEmpty returns e is empty or not.
func (e Expression) IsEmpty() bool {
	return e.sql == "" && len(e.values) == 0
}
